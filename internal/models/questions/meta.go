package questions

import (
	"fmt"

	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

var _ abstractions.Keyer = (*Meta)(nil)

type Meta struct {
	ID         uuid.UUID  `json:"id"`
	CategoryID int32      `json:"category_id"`
	Difficulty Difficulty `json:"difficulty"`
	Language   string     `json:"language"`
}

func (m *Meta) Key() string {
	return fmt.Sprintf("questions:%s:%s:%d", m.Language, m.Difficulty, m.CategoryID)
}
