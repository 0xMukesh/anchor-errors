package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	programs, err := getAllPrograms()
	if err != nil {
		fmt.Println(err)
		return
	}
	items := []list.Item{}
	for _, program := range programs {
		items = append(items, program)
	}
	m := Model{list: list.New(items, listDelegate(), 0, 0)}
	m.list.Title = "anchor-errors"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("an error occured", err)
		os.Exit(1)
	}
}
