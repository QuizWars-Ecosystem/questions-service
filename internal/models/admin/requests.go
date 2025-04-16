package admin

import (
	"crypto/md5"
	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	"strings"
	"time"
)

var _ abstractions.Requestable[QuestionsFilter, *questionsv1.GetFilteredQuestionsRequest] = (*QuestionsFilter)(nil)

func (q QuestionsFilter) Request(req *questionsv1.GetFilteredQuestionsRequest) (*QuestionsFilter, error) {
	q.Offset, q.Limit = offsetLimit(req.Page, req.Size)

	if req.Order != nil {
		q.Order = orderFromGRPCEnum(*req.Order)
	} else {
		q.Order = Category
	}

	if req.Sort != nil {
		q.Sort = sortFromGRPCEnum(*req.Sort)
	} else {
		q.Sort = ASC
	}

	if req.TypeFilter != nil {
		var array = make([]questions.Type, len(req.TypeFilter.Types))
		for i, element := range req.TypeFilter.Types {
			array[i] = questions.TypeFromGRPCEnum(element)
		}
		q.TypesFilter = &ArrayFilter[questions.Type]{Array: array}
	}

	if req.DifficultyFilter != nil {
		var array = make([]questions.Difficulty, len(req.DifficultyFilter.Difficulties))
		for i, element := range req.DifficultyFilter.Difficulties {
			array[i] = questions.DifficultyFromGRPCEnum(element)
		}
		q.DifficultiesFilter = &ArrayFilter[questions.Difficulty]{Array: array}
	}

	if req.CategoryFilter != nil {
		q.CategoriesFilter = &ArrayFilter[int32]{Array: req.CategoryFilter.Categories}
	}

	if req.LanguageFilter != nil {
		q.LanguagesFilter = &ArrayFilter[string]{Array: req.LanguageFilter.Languages}
	}

	if req.CreateAtFilter != nil {
		q.CreatedAtFilter = &Filter[time.Time]{
			From: req.CreateAtFilter.From.AsTime(),
			To:   req.CreateAtFilter.To.AsTime(),
		}
	}

	return &q, nil
}

func offsetLimit(page, size uint64) (uint64, uint64) {
	offset := (page - 1) * size
	limit := size

	return offset, limit
}

var _ abstractions.Requestable[UpdateQuestionRequest, *questionsv1.UpdateQuestionRequest] = (*UpdateQuestionRequest)(nil)

func (u UpdateQuestionRequest) Request(req *questionsv1.UpdateQuestionRequest) (*UpdateQuestionRequest, error) {
	if req.Type != nil {
		val := questions.TypeFromGRPCEnum(*req.Type)
		u.Type = &val
	}

	if req.Difficulty != nil {
		val := questions.DifficultyFromGRPCEnum(*req.Difficulty)
		u.Difficulty = &val
	}

	if req.CategoryId != nil {
		u.CategoryID = req.CategoryId
	}

	if req.Text != nil {
		text := strings.ToLower(strings.Trim(strings.TrimSpace(req.GetText()), " "))
		hash := md5.Sum([]byte(text))

		u.Text = req.Text
		u.Hash = string(hash[:])
	}

	if req.Language != nil {
		u.Language = req.Language
	}

	return &u, nil
}

var _ abstractions.Requestable[UpdateQuestionOptionRequest, *questionsv1.UpdateQuestionOptionRequest] = (*UpdateQuestionOptionRequest)(nil)

func (u UpdateQuestionOptionRequest) Request(req *questionsv1.UpdateQuestionOptionRequest) (*UpdateQuestionOptionRequest, error) {
	if req.Text != nil {
		u.Text = req.Text
	}

	if req.IsCorrect != nil {
		u.IsCorrect = req.IsCorrect
	}

	return &u, nil
}
