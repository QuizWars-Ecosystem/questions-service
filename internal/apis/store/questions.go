package store

import (
	"context"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/filter"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	"go.uber.org/zap"
)

func (s *Store) GetQuestions(ctx context.Context, language string, difficulty questions.Difficulty, categoryID, amount int32) ([]*questions.Question, error) {
	ids, count, err := s.cache.GetCachedIDs(ctx, language, difficulty.String(), categoryID, amount)
	if err != nil {
		s.logger.Warn("failed to get cached ids", zap.String("language", language), zap.Error(err))
	}

	var qs []*questions.Question
	if count >= int(amount) {
		qs, err = s.db.GetQuestionsByIDs(ctx, ids)
		if err != nil {
			s.logger.Debug("failed to get questions by cached ids", zap.String("language", language), zap.Error(err))
			return nil, err
		}

		return qs, nil
	} else {

		if count > 0 {
			qs, err = s.db.GetQuestionsByIDs(ctx, ids)
			if err != nil {
				s.logger.Debug("failed to get questions by cached ids", zap.String("language", language), zap.Error(err))
			}
		}

		f := &filter.QuestionsFilter{
			Difficulties: []questions.Difficulty{difficulty},
			Categories:   []int32{categoryID},
			Language:     language,
			Amount:       amount - int32(count),
		}

		var qsa []*questions.Question
		qsa, err = s.db.GetFilteredRandomQuestions(ctx, f)
		if err != nil {
			s.logger.Debug("failed to get filtered random questions", zap.String("language", language), zap.Error(err))
			return nil, err
		}

		qs = append(qs, qsa...)
	}

	return qs, nil
}

func (s *Store) GetQuestionsBatch(ctx context.Context, filter *filter.QuestionsFilter) ([]*questions.Question, error) {
	var difficulties = make([]string, len(filter.Difficulties))
	for i, d := range filter.Difficulties {
		difficulties[i] = d.String()
	}

	ids, count, err := s.cache.GetBatchCachedIDs(ctx, filter.Language, difficulties, filter.Categories, filter.Amount)
	if err != nil {
		s.logger.Warn("failed to get cached ids", zap.String("language", filter.Language), zap.Error(err))
	}

	var qs []*questions.Question
	if count >= int(filter.Amount) {
		qs, err = s.db.GetQuestionsByIDs(ctx, ids)
		if err != nil {
			s.logger.Debug("failed to get questions by cached ids", zap.String("language", filter.Language), zap.Error(err))
			return nil, err
		}

		return qs, nil
	} else {
		filter.Amount -= int32(count)

		if count > 0 {
			qs, err = s.db.GetQuestionsByIDs(ctx, ids)
			if err != nil {
				s.logger.Debug("failed to get questions by cached ids", zap.String("language", filter.Language), zap.Error(err))
			}
		}

		var qsa []*questions.Question
		qsa, err = s.db.GetFilteredRandomQuestions(ctx, filter)
		if err != nil {
			s.logger.Debug("failed to get filtered random questions", zap.String("language", filter.Language), zap.Error(err))
			return nil, err
		}

		qs = append(qs, qsa...)
	}

	return qs, nil
}
