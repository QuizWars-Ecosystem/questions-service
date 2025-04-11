package filter

import "github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"

type QuestionsFilter struct {
	Types        []questions.Type
	Sources      []questions.Source
	Difficulties []questions.Difficulty
	Categories   []int32
	Languages    []string
	Amount       int32
}
