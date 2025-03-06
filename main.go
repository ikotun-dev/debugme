package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	errorMessage     string
	errorExplanation string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Bold(true)

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n%s\n",
		style.Render("❗ @debugme:"),
		m.errorMessage,
		style.Render("💡 Explanation:"),
		m.errorExplanation,
	)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	errorMessage := strings.TrimSpace(input)

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
	if strings.Contains(err, "connection refused") {
		return "This error usually happens when the server is down, or the connection details (host, port) are wrong."
	}
	return "No known explanation found. Try checking your logs or documentation."
}
