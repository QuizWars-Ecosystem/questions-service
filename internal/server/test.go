package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/DavidMovas/gopherbox/pkg/closer"
	"github.com/QuizWars-Ecosystem/go-common/pkg/clients"
	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	"github.com/QuizWars-Ecosystem/go-common/pkg/log"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/handler"
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/service"
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/store"
	"github.com/QuizWars-Ecosystem/questions-service/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type TestServer struct {
	grpcServer *grpc.Server
	listener   net.Listener
	logger     *log.Logger
	cfg        *config.Config
	closer     *closer.Closer
}

func NewTestServer(ctx context.Context, cfg *config.Config) (*TestServer, error) {
	cl := closer.NewCloser()

	logger := log.NewLogger(cfg.Local, cfg.Logger.Level)

	postgresOptions := clients.NewPostgresOptions(cfg.Postgres.URL)
	postgresOptions.WithConnectTimeout(time.Second * 20)

	postgresClient, err := clients.NewPostgresClient(ctx, cfg.Postgres.URL, postgresOptions)
	if err != nil {
		logger.Zap().Error("error initializing postgres client", zap.Error(err))
		return nil, fmt.Errorf("error initializing postgres client: %w", err)
	}

	redisClient, err := clients.NewRedisClient(cfg.Redis.URL, nil)
	if err != nil {
		logger.Zap().Error("error initializing redis client", zap.Error(err))
		return nil, fmt.Errorf("error initializing redis client: %w", err)
	}

	jwtService := jwt.NewService(cfg.JWT)

	storage := store.NewStore(postgresClient, redisClient, logger.Zap(), cfg.StoreConfig)
	srv := service.NewService(storage, logger.Zap())
	hand := handler.NewHandler(srv, jwtService, logger.Zap())

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor())

	questionsv1.RegisterQuestionsServiceServer(grpcServer, hand)
	questionsv1.RegisterQuestionsClientServiceServer(grpcServer, hand)
	questionsv1.RegisterQuestionsAdminServiceServer(grpcServer, hand)

	return &TestServer{
		grpcServer: grpcServer,
		logger:     logger,
		cfg:        cfg,
		closer:     cl,
	}, nil
}

func (s *TestServer) Start() error {
	z := s.logger.Zap()

	z.Info("Starting server", zap.String("name", s.cfg.Name), zap.Int("port", s.cfg.GRPCPort))

	var err error
	s.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GRPCPort))
	if err != nil {
		z.Error("Failed to start listener", zap.String("name", s.cfg.Name), zap.Int("port", s.cfg.GRPCPort), zap.Error(err))
		return err
	}

	return s.grpcServer.Serve(s.listener)
}

func (s *TestServer) Shutdown(ctx context.Context) error {
	z := s.logger.Zap()
	z.Info("Shutting down server gracefully", zap.String("name", s.cfg.Name))

	stopChan := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(stopChan)
	}()

	select {
	case <-stopChan:
	case <-ctx.Done():
		z.Warn("Graceful shutdown timed out, forcing stop")
		s.grpcServer.Stop()
	}

	if err := s.listener.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
		return fmt.Errorf("shutting down listener: %w", err)
	}

	if err := s.logger.Close(); err != nil && !isStdoutSyncErr(err) {
		return fmt.Errorf("error closing logger: %w", err)
	}

	return s.closer.Close(ctx)
}

func isStdoutSyncErr(err error) bool {
	return strings.Contains(err.Error(), "sync")
}
