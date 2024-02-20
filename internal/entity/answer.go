package entity

type Answer struct {
	Score   int   `json:"score"`
	Choices []int `json:"choices"`
}

func (a Answer) ScoreAsPercentage() float64 {
	return (float64(a.Score) / float64(len(a.Choices))) * 100
}
