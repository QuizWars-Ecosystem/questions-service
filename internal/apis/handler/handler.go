package handler

import (
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/service"
	"go.uber.org/zap"
)

var _ questionsv1.QuestionsServiceServer = (*Handler)(nil)
var _ questionsv1.QuestionsAdminServiceServer = (*Handler)(nil)
var _ questionsv1.QuestionsClientServiceServer = (*Handler)(nil)

type Handler struct {
	service *service.Service
	logger  *zap.Logger
}

func NewHandler(service *service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
