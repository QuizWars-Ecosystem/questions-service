package handler

import (
	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/service"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

var Empty = &emptypb.Empty{}

var (
	_ questionsv1.QuestionsServiceServer       = (*Handler)(nil)
	_ questionsv1.QuestionsAdminServiceServer  = (*Handler)(nil)
	_ questionsv1.QuestionsClientServiceServer = (*Handler)(nil)
)

type Handler struct {
	service service.IService
	auth    *jwt.Service
	logger  *zap.Logger
}

func NewHandler(service service.IService, auth *jwt.Service, logger *zap.Logger) *Handler {
	return &Handler{
		service: service,
		auth:    auth,
		logger:  logger,
	}
}
