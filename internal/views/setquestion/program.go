package setquestion

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codeaspiras/goquiz/internal/entity"
)

func Run(
	questionIndex int,
	question entity.QuizItem,
) (int, error) {
	model := newModel(
		questionIndex,
		question,
	)
	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle(fmt.Sprintf("#goquiz > answer question %d", questionIndex+1))
	if _, err := program.Run(); err != nil {
		return -1, err
	}

	return model.choice, nil
}
