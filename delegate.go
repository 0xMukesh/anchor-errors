package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type delegateKeyMap struct {
	// enter
	choose key.Binding
	// backspace
	previousPage key.Binding
}

func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.previousPage,
	}
}

func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.previousPage,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys(tea.KeyEnter.String()),
			key.WithHelp(tea.KeyEnter.String(), "choose"),
		),
		previousPage: key.NewBinding(
			key.WithKeys(tea.KeyBackspace.String(), tea.KeyCtrlLeft.String()),
			key.WithHelp(tea.KeyBackspace.String(), "goto previous page"),
		),
	}
}

func listDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	keys := newDelegateKeyMap()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string
		var items []list.Item

		_, isOnProgramPage := m.SelectedItem().(Program)

		if i, ok := m.SelectedItem().(Program); ok {
			title = i.Title()
		} else if i, ok := m.SelectedItem().(ErrorCode); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				if isOnProgramPage {
					codes, err := getErrorCodes(title)
					if err != nil {
						fmt.Println(err)
						return nil
					}
					for _, i := range codes {
						items = append(items, i)
					}
					m.ResetFilter()
					return tea.Batch(m.SetItems(items))
				}
			case key.Matches(msg, keys.previousPage):
				programs, err := getAllPrograms()
				if err != nil {
					fmt.Println(err)
					return nil
				}
				for _, i := range programs {
					items = append(items, i)
				}
				m.ResetFilter()
				return tea.Batch(m.SetItems(items))
			}
		}
		return nil
	}

	help := []key.Binding{keys.choose, keys.previousPage}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}
