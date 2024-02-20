package showresult

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codeaspiras/goquiz/internal/entity"
	"github.com/codeaspiras/goquiz/internal/views/style"
)

type (
	errMsg error
)

type model struct {
	questions entity.Quiz
	answer    entity.Answer
	err       error
}

func newModel(
	questions entity.Quiz,
	answer entity.Answer,
) *model {
	return &model{
		questions: questions,
		answer:    answer,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc, tea.KeyEnter:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	answersTable := ""
	for questionIndex, question := range m.questions {
		answersTable += fmt.Sprintf(" %d) %s\n", (questionIndex + 1), question.Question)
		choosedOptionIndex := m.answer.Choices[questionIndex]
		if choosedOptionIndex == -1 {
			continue
		}

		if choosedOptionIndex != question.RightAnswerIndex {
			answersTable += fmt.Sprintf("   %s\n", style.WrongAnswerStyle.Render(question.Options[choosedOptionIndex]))
		}
		answersTable += fmt.Sprintf("   %s\n", style.RightAnswerStyle.Render(question.Options[question.RightAnswerIndex]))
	}

	return fmt.Sprintf(
		`
 Sua pontuação foi de %.f%s!
 =========================================
%s`,
		m.answer.ScoreAsPercentage(), "%",
		answersTable,
	)
}
