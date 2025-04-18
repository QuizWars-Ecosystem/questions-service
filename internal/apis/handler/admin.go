package handler

import (
	"context"

	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	"github.com/QuizWars-Ecosystem/go-common/pkg/uuidx"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/admin"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"

	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetFilteredQuestions(ctx context.Context, request *questionsv1.GetFilteredQuestionsRequest) (*questionsv1.GetFilteredQuestionsResponse, error) {
	err := h.auth.ValidateRoleWithContext(ctx, string(jwt.Admin))
	if err != nil {
		return nil, err
	}

	req, err := abstractions.MakeRequest[admin.QuestionsFilter](request)
	if err != nil {
		return nil, err
	}

	qs, amount, err := h.service.GetFilteredQuestions(ctx, req)
	if err != nil {
		return nil, err
	}

	questionsList := make([]*questionsv1.Question, len(qs))
	for i, q := range qs {
		var question *questionsv1.Question

		if question, err = q.Response(); err != nil {
			return nil, err
		}

		questionsList[i] = question
	}

	return &questionsv1.GetFilteredQuestionsResponse{
		Questions: questionsList,
		Page:      req.Offset,
		Size:      req.Limit,
		Order:     req.Order.ToGRPCEnum(),
		Sort:      req.Sort.ToGRPCEnum(),
		Amount:    int64(amount),
	}, nil
}

func (h *Handler) CreateCategory(ctx context.Context, request *questionsv1.CreateCategoryRequest) (*questionsv1.CreateCategoryResponse, error) {
	err := h.auth.ValidateRoleWithContext(ctx, string(jwt.Admin))
	if err != nil {
		return nil, err
	}

	id, err := h.service.CreateCategory(ctx, request.GetName())
	if err != nil {
		return nil, err
	}

	return &questionsv1.CreateCategoryResponse{
		Id: id,
	}, nil
}

func (h *Handler) CreateQuestion(ctx context.Context, request *questionsv1.CreateQuestionRequest) (*emptypb.Empty, error) {
	err := h.auth.ValidateRoleWithContext(ctx, string(jwt.Admin))
	if err != nil {
		return nil, err
	}

	req, err := abstractions.MakeRequest[questions.CreateQuestionRequest](request)
	if err != nil {
		return nil, err
	}

	err = h.service.CreateQuestion(ctx, req)
	if err != nil {
		return nil, err
	}

	return Empty, nil
}

func (h *Handler) UpdateCategory(ctx context.Context, request *questionsv1.UpdateCategoryRequest) (*emptypb.Empty, error) {
	err := h.auth.ValidateRoleWithContext(ctx, string(jwt.Admin))
	if err != nil {
		return nil, err
	}

	err = h.service.UpdateCategory(ctx, request.GetId(), request.GetName())
	if err != nil {
		return nil, err
	}

	return Empty, nil
}

func (h *Handler) UpdateQuestion(ctx context.Context, request *questionsv1.UpdateQuestionRequest) (*emptypb.Empty, error) {
	err := h.auth.ValidateRoleWithContext(ctx, string(jwt.Admin))
	if err != nil {
		return nil, err
	}

	id, err := uuidx.NewUUIDFromString(request.GetId())
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	req, err := abstractions.MakeRequest[admin.UpdateQuestionRequest](request)
	if err != nil {
		return nil, err
	}

	err = h.service.UpdateQuestion(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return Empty, nil
}

func (h *Handler) UpdateQuestionOption(ctx context.Context, request *questionsv1.UpdateQuestionOptionRequest) (*emptypb.Empty, error) {
	err := h.auth.ValidateRoleWithContext(ctx, string(jwt.Admin))
	if err != nil {
		return nil, err
	}

	id, err := uuidx.NewUUIDFromString(request.GetId())
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	req, err := abstractions.MakeRequest[admin.UpdateQuestionOptionRequest](request)
	if err != nil {
		return nil, err
	}

	err = h.service.UpdateQuestionOption(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return Empty, nil
}

func (h *Handler) DeleteQuestion(ctx context.Context, request *questionsv1.DeleteQuestionRequest) (*emptypb.Empty, error) {
	err := h.auth.ValidateRoleWithContext(ctx, string(jwt.Admin))
	if err != nil {
		return nil, err
	}

	id, err := uuidx.NewUUIDFromString(request.GetId())
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	err = h.service.DeleteQuestion(ctx, id)
	if err != nil {
		return nil, err
	}

	return Empty, nil
}

func (h *Handler) DeleteQuestionOption(ctx context.Context, request *questionsv1.DeleteQuestionOptionRequest) (*emptypb.Empty, error) {
	err := h.auth.ValidateRoleWithContext(ctx, string(jwt.Admin))
	if err != nil {
		return nil, err
	}

	id, err := uuidx.NewUUIDFromString(request.GetId())
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	err = h.service.DeleteQuestionOption(ctx, id)
	if err != nil {
		return nil, err
	}

	return Empty, nil
}
