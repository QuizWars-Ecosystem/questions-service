package questions

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

type Category struct {
	ID   int32
	Name string
}

type Question struct {
	ID         uuid.UUID
	Type       Type
	Source     Source
	Difficulty Difficulty
	Category   Category
	Text       string
	Options    []*Option
	Language   string
	CreatedAt  time.Time
}

type Hashed struct {
	Question
	Hash string
}

type Option struct {
	ID        uuid.UUID
	Text      string
	IsCorrect bool
}
