package setquestion

type item struct {
	index int
	label string
}

func (i item) FilterValue() string {
	return i.label
}
