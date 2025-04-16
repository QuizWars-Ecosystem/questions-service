package questions

import questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"

type Type string

const (
	Single  Type = "Single"
	Multi   Type = "Multi"
	Betting Type = "Betting"
)

func (t Type) String() string {
	return string(t)
}

func (t Type) TypeToGRPCEnum() questionsv1.Type {
	switch t {
	case Single:
		return questionsv1.Type_TYPE_SINGLE
	case Multi:
		return questionsv1.Type_TYPE_MULTI
	case Betting:
		return questionsv1.Type_TYPE_BETTING
	default:
		return questionsv1.Type_TYPE_UNSPECIFIED
	}
}

func TypeFromGRPCEnum(enum questionsv1.Type) Type {
	switch enum {
	case questionsv1.Type_TYPE_SINGLE:
		return Single
	case questionsv1.Type_TYPE_MULTI:
		return Multi
	case questionsv1.Type_TYPE_BETTING:
		return Betting
	default:
		return Single
	}
}

type Source string

const (
	Text      Source = "Text"
	Image     Source = "Image"
	Audio     Source = "Audio"
	Animation Source = "Animation"
	Video     Source = "Video"
)

func (s Source) SourceToGRPCEnum() questionsv1.Source {
	switch s {
	case Text:
		return questionsv1.Source_SOURCE_TEXT
	case Image:
		return questionsv1.Source_SOURCE_IMAGE
	case Audio:
		return questionsv1.Source_SOURCE_AUDIO
	case Animation:
		return questionsv1.Source_SOURCE_ANIMATION
	case Video:
		return questionsv1.Source_SOURCE_VIDEO
	default:
		return questionsv1.Source_SOURCE_UNSPECIFIED
	}
}

func SourceFromGRPCEnum(enum questionsv1.Source) Source {
	switch enum {
	case questionsv1.Source_SOURCE_TEXT:
		return Text
	case questionsv1.Source_SOURCE_IMAGE:
		return Image
	case questionsv1.Source_SOURCE_AUDIO:
		return Audio
	case questionsv1.Source_SOURCE_ANIMATION:
		return Animation
	case questionsv1.Source_SOURCE_VIDEO:
		return Video
	default:
		return Text
	}
}

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

func (d Difficulty) DifficultyToGRPCEnum() questionsv1.Difficulty {
	switch d {
	case Easy:
		return questionsv1.Difficulty_DIFFICULTY_EASY
	case Medium:
		return questionsv1.Difficulty_DIFFICULTY_MEDIUM
	case Hard:
		return questionsv1.Difficulty_DIFFICULTY_HARD
	case VeryHard:
		return questionsv1.Difficulty_DIFFICULTY_VERY_HARD
	default:
		return questionsv1.Difficulty_DIFFICULTY_UNSPECIFIED
	}
}

func DifficultyFromGRPCEnum(d questionsv1.Difficulty) Difficulty {
	switch d {
	case questionsv1.Difficulty_DIFFICULTY_EASY:
		return Easy
	case questionsv1.Difficulty_DIFFICULTY_MEDIUM:
		return Medium
	case questionsv1.Difficulty_DIFFICULTY_HARD:
		return Hard
	case questionsv1.Difficulty_DIFFICULTY_VERY_HARD:
		return VeryHard
	default:
		return Easy
	}
}
