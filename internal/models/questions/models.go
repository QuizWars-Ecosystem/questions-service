package questions

import (
	"github.com/google/uuid"
	"time"
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
	Categories []*Category
	Text       string
	Hash       string
	Answers    []*Option
	Language   string
	CreatedAt  time.Time
}

type Option struct {
	ID        uuid.UUID
	Text      string
	IsCorrect bool
}

type Meta struct {
	ID         uuid.UUID  `json:"id"`
	Categories []int32    `json:"categories"`
	Difficulty Difficulty `json:"difficulty"`
	Language   string     `json:"language"`
}

type Type string

const (
	Single  Type = "Single"
	Multi   Type = "Multi"
	Betting Type = "Betting"
)

func (t Type) String() string {
	return string(t)
}

type Source string

const (
	Text      Source = "Text"
	Image     Source = "Image"
	Audio     Source = "Audio"
	Animation Source = "Animation"
	Video     Source = "Video"
)

func (s Source) String() string {
	return string(s)
}

type Difficulty string

const (
	Easy     Difficulty = "Easy"
	Medium   Difficulty = "Medium"
	Hard     Difficulty = "Hard"
	VeryHard Difficulty = "Vary Hard"
)

func (d Difficulty) String() string {
	return string(d)
}
