package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n " + m.table.HelpView() + "\n"
}

func main() {
	CallClear()

	bsLogo()

	columns := []table.Column{
		{Title: "Welcome To GoSky", Width: 50},
	}

	rows := []table.Row{
		{"Login"},
		{"Post"},
		{"Repost"},
		{"Timeline"},
		{"Search"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func bsLogo() string {
	b, err := os.ReadFile("asciiGsky.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(string(b))

	return string(b)
}

func help() {

	Bold := color.New(color.Bold)
	blue := color.New(color.FgBlue, color.Bold)

	blue.Println("General FAQ: ")
	Bold.Print("Usage: ")
	fmt.Print("Hello")
	blue.Println("Options:")

}

var clear map[string]func()

func init() {

	clear = make(map[string]func())

	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux and Unix Clear Func
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls") // Windows clear func (use 'cls' instead of 'clear')
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Platform is unsupported!! Terminal screen will not be cleared!!")
	}
}
