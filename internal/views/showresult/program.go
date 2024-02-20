package showresult

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codeaspiras/goquiz/internal/entity"
)

func Run(
	questions entity.Quiz,
	answer entity.Answer,
) error {
	model := newModel(questions, answer)
	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle("#goquiz > result")
	_, err := program.Run()
	return err
}
