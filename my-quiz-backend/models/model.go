package models

type Question struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
	Answer  int      `json:"answer"` // index of the correct option
}

type SubmittedAnswer struct {
	UserID   string            `json:"userId"`
	Username string            `json:"username"`
	Answers  map[string]string `json:"answers"` // map of question ID to selected option index
}
