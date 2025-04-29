package admin

import (
	"time"

	"github.com/google/uuid"

	questionspb "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
)

const (
	ID         Order = "id"
	Text       Order = "text"
	Type       Order = "type"
	Difficulty Order = "difficulty"
	Category   Order = "category_id"
	Language   Order = "language"
	CreatedAt  Order = "created_at"
)

const (
	ASC  Sort = "ASC"
	DESC Sort = "DESC"
)

type CreateQuestionOptionRequest struct {
	ID        uuid.UUID
	Text      string
	IsCorrect bool
}

type UpdateQuestionRequest struct {
	Type       *questions.Type
	Difficulty *questions.Difficulty
	Source     *questions.Source
	CategoryID *int32
	Text       *string
	Hash       string
	Language   *string
}

type UpdateQuestionOptionRequest struct {
	Text      *string
	IsCorrect *bool
}

type Filter[T any] struct {
	From T
	To   T
}

type ArrayFilter[T any] struct {
	Array []T
}

type QuestionsFilter struct {
	Offset             uint64
	Limit              uint64
	Order              Order
	Sort               Sort
	TypesFilter        *ArrayFilter[questions.Type]
	DifficultiesFilter *ArrayFilter[questions.Difficulty]
	CategoriesFilter   *ArrayFilter[int32]
	LanguagesFilter    *ArrayFilter[string]
	CreatedAtFilter    *Filter[time.Time]
}

type Order string

func (o Order) String() string {
	return string(o)
}

type Sort string

func (s Sort) String() string {
	return string(s)
}

func (o Order) ToGRPCEnum() questionspb.Order {
	switch o {
	case ID:
		return questionspb.Order_ORDER_ID
	case Text:
		return questionspb.Order_ORDER_TEXT
	case Type:
		return questionspb.Order_ORDER_TYPE
	case Difficulty:
		return questionspb.Order_ORDER_DIFFICULTY
	case Category:
		return questionspb.Order_ORDER_CATEGORY
	case Language:
		return questionspb.Order_ORDER_LANGUAGE
	case CreatedAt:
		return questionspb.Order_ORDER_CREATED_AT
	default:
		return questionspb.Order_ORDER_DIFFICULTY
	}
}

func (s Sort) ToGRPCEnum() questionspb.Sort {
	switch s {
	case ASC:
		return questionspb.Sort_SORT_ASC
	case DESC:
		return questionspb.Sort_SORT_DESC
	default:
		return questionspb.Sort_SORT_DESC
	}
}

func orderFromGRPCEnum(status questionspb.Order) Order {
	switch status {
	case questionspb.Order_ORDER_ID:
		return ID
	case questionspb.Order_ORDER_TEXT:
		return Text
	case questionspb.Order_ORDER_TYPE:
		return Type
	case questionspb.Order_ORDER_DIFFICULTY:
		return Difficulty
	case questionspb.Order_ORDER_CATEGORY:
		return Category
	case questionspb.Order_ORDER_LANGUAGE:
		return Language
	case questionspb.Order_ORDER_CREATED_AT:
		return CreatedAt
	default:
		return Difficulty
	}
}

func sortFromGRPCEnum(status questionspb.Sort) Sort {
	switch status {
	case questionspb.Sort_SORT_ASC:
		return ASC
	case questionspb.Sort_SORT_DESC:
		return DESC
	default:
		return DESC
	}
}
