package main

import (
	"bufio"
	"fmt"
	"os"
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
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Bold(true)

	wrappedExplanation := wordwrap.String(m.errorExplanation, m.width)

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n%s\n",
		style.Render("❗ @debugme:"),
		m.errorMessage,
		style.Render("💡 Explanation:"),
		wrappedExplanation,
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
	return errGenerator.InitOpenAI(err)
}
