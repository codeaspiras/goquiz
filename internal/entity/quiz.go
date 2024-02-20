package entity

type Quiz []QuizItem

type QuizItem struct {
	Question         string   `json:"question"`
	Options          []string `json:"options"`
	RightAnswerIndex int      `json:"right_answer"`
}
