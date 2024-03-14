package models

import "github.com/google/uuid"

type Question struct {
	ID      string            `json:"id"`
	Text    map[string]string `json:"text"`
	Options []string          `json:"options"`
	Answer  int               `json:"answer"` // index of the correct option
}

type SubmittedAnswer struct {
	UserID   uuid.UUID         `json:"userId"`
	Username string            `json:"username"`
	Answers  map[string]string `json:"answers"` // map of question ID to selected option index
}

type QuizResult struct {
	UserID     uuid.UUID
	Score      int
	Percentile float64
}
