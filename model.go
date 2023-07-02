package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list list.Model
}

type ErrorCode struct {
	Code    string `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type Program struct {
	Name string `json:"name"`
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
	return docStyle.Render(m.list.View())
}
