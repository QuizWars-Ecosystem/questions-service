package questions

import (
	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ abstractions.Responseable[questionsv1.Category] = (*Category)(nil)

func (c *Category) Response() (*questionsv1.Category, error) {
	var req questionsv1.Category
	req.Id = c.ID
	req.Name = c.Name

	return &req, nil
}

var _ abstractions.Responseable[questionsv1.Question] = (*Question)(nil)

func (q *Question) Response() (*questionsv1.Question, error) {
	var req questionsv1.Question
	var err error

	req.Id = q.ID.String()
	req.Type = q.Type.TypeToGRPCEnum()
	req.Source = q.Source.SourceToGRPCEnum()
	req.Difficulty = q.Difficulty.DifficultyToGRPCEnum()
	req.Category, err = q.Category.Response()
	if err != nil {
		return nil, err
	}

	req.Text = q.Text
	req.Language = q.Language
	req.CreatedAt = timestamppb.New(q.CreatedAt)

	options := make([]*questionsv1.Option, len(q.Options))
	for i, o := range q.Options {
		var option *questionsv1.Option
		option, err = o.Response()
		if err != nil {
			return nil, err
		}

		options[i] = option
	}

	req.Options = options

	return &req, nil
}

var _ abstractions.Responseable[questionsv1.Option] = (*Option)(nil)

func (o *Option) Response() (*questionsv1.Option, error) {
	var req questionsv1.Option

	req.Id = o.ID.String()
	req.Text = o.Text
	req.IsCorrect = o.IsCorrect

	return &req, nil
}
