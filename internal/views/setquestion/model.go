package setquestion

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codeaspiras/goquiz/internal/entity"
	"github.com/codeaspiras/goquiz/internal/views/style"
)

type (
	errMsg error
)

type model struct {
	list   list.Model
	choice int
	err    error
}

func newModel(questionIndex int, question entity.QuizItem) *model {
	options := make([]list.Item, len(question.Options))
	for i, o := range question.Options {
		options[i] = item{
			index: i,
			label: o,
		}
	}

	list := list.New(
		options,
		itemDelegate{},
		30,
		10,
	)
	list.Title = fmt.Sprintf("# %d) %s", (questionIndex + 1), question.Question)
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(false)
	list.Styles.Title = style.TitleStyle
	list.Styles.PaginationStyle = style.PaginationStyle
	list.Styles.HelpStyle = style.HelpStyle

	return &model{
		list:   list,
		choice: -1,
		err:    nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			value, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = value.index
			}

			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.list.View()
}
