package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list                     list.Model
	isOnErrorCodeDetailsPage bool
	errorCodeDetailsPageView string
}

type ErrorCode struct {
	Code    string `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Program string `json:"program"`
}

type Program struct {
	Name   string
	Source string
}

func (e ErrorCode) Title() string { return e.Error }
func (e ErrorCode) Description() string {
	hex := strToHex(e.Code)
	return hex + " - " + e.Message
}
func (e ErrorCode) FilterValue() string { return strToHex(e.Code) }

func (p Program) Title() string { return p.Name }
func (p Program) Description() string {
	switch p.Name {
	case "anchor":
		return "native anchor framework errors"
	default:
		return fmt.Sprintf("errors of %s program", p.Name)
	}
}
func (p Program) FilterValue() string { return p.Name }

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			errorCode, isOnErrorCodesPage := m.list.SelectedItem().(ErrorCode)
			program, isOnProgramsPage := m.list.SelectedItem().(Program)

			if isOnErrorCodesPage {
				m.isOnErrorCodeDetailsPage = true
				m.errorCodeDetailsPageView = fmt.Sprintf("code - %s \nmessage - %s \nname - %s \nsource - %s", strToHex(errorCode.Code), errorCode.Message, errorCode.Error, errorCodesSource[errorCode.Program])
			} else if isOnProgramsPage {
				var items []list.Item

				codes, err := getErrorCodes(program.Name)
				if err != nil {
					fmt.Println(err)
					return m, nil
				}
				for _, i := range codes {
					items = append(items, i)
				}
				m.list.ResetFilter()
				return m, tea.Batch(m.list.SetItems(items))
			}
		case tea.KeyBackspace:
			if m.isOnErrorCodeDetailsPage {
				m.isOnErrorCodeDetailsPage = false
			} else {
				var items []list.Item

				programs, err := getAllPrograms()
				if err != nil {
					fmt.Println(err)
					return m, nil
				}
				for _, i := range programs {
					items = append(items, i)
				}
				m.list.ResetFilter()
				return m, m.list.SetItems(items)
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var view string

	if m.isOnErrorCodeDetailsPage {
		view = m.errorCodeDetailsPageView
	} else {
		view = m.list.View()
	}

	return docStyle.Render(view)
}
