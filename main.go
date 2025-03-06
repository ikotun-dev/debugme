package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	errGenerator "github.com/ikotun/llmxp/internals/handlers"
	"github.com/muesli/reflow/wordwrap"
)

type model struct {
	errorMessage     string
	errorExplanation string
	width            int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	return m, nil
}

func (m model) View() string {
	headerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Bold(true)
	_ = lipgloss.NewStyle().Padding(0, 2)

	errorRow := fmt.Sprintf("%s: %s", headerStyle.Render("Error"), m.errorMessage)
	explanationRow := fmt.Sprintf("%s: %s", headerStyle.Render("Explanation"), wordwrap.String(m.errorExplanation, m.width-4))

	table := lipgloss.JoinVertical(lipgloss.Left, errorRow, explanationRow)

	box := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1).
		Width(m.width - 10).
		Render(table)

	return box
}

func main() {
	var errorMessage string

	if len(os.Args) >= 3 && os.Args[1] == "--command" {
		cmd := exec.Command("sh", "-c", os.Args[2])
		output, _ := cmd.CombinedOutput()
		errorMessage = string(output)
	} else if len(os.Args) >= 3 && os.Args[1] == "--file" {
		content, err := os.ReadFile(os.Args[2])
		if err != nil {
			fmt.Println("Failed to read file:", err)
			os.Exit(1)
		}
		errorMessage = string(content)
	} else if len(os.Args) >= 2 {
		errorMessage = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Println("Usage: go run main.go [--command 'cmd'] OR [--file file.log] OR 'error message directly'")
		os.Exit(1)
	}

	errorMessage = strings.TrimSpace(errorMessage)
	explanation := lookupExplanation(errorMessage)

	p := tea.NewProgram(model{
		errorMessage:     errorMessage,
		errorExplanation: explanation,
	})
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app: %v\n", err)
		os.Exit(1)
	}
}

func lookupExplanation(err string) string {
	return errGenerator.InitOpenAI(err)
}
