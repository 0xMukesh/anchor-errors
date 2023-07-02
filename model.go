package main

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list list.Model
}

type errors struct {
	Code    string `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type program struct {
	Name string `json:"name"`
}

func (e errors) Title() string { return e.Error }
func (e errors) Description() string {
	code, err := strconv.Atoi(e.Code)
	if err != nil {
		fmt.Print(err)
	}
	hex := fmt.Sprintf("%x", code)
	return "0x" + hex + " - " + e.Message
}
func (e errors) FilterValue() string {
	code, err := strconv.Atoi(e.Code)
	if err != nil {
		fmt.Print(err)
	}
	hex := fmt.Sprintf("%x", code)
	return "0x" + hex
}

func (p program) Title() string { return p.Name }
func (p program) Description() string {
	switch p.Name {
	case "anchor":
		return "native anchor framework errors"
	default:
		return fmt.Sprintf("errors due to %s program", p.Name)
	}
}
func (p program) FilterValue() string { return p.Name }

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	return docStyle.Render(m.list.View())
}
