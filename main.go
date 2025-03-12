package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	errGenerator "github.com/ikotun/debugme/internals/handlers"
)

type model struct {
	errorMessage     string
	errorExplanation string
	width            int
}

func (m model) Init() tea.Cmd {
	return tea.Quit
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
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFA500")).
		Bold(true)

	contentStyle := lipgloss.NewStyle().
		Padding(0, 2).
		Width(m.width - 8)

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Margin(0, 0, 1, 1).
		Width(m.width - 4)

	wrappedExplanation := wordwrap(m.errorExplanation, m.width-8)

	content := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n%s",
		headerStyle.Render("❗@debugme"),
		m.errorMessage,
		headerStyle.Render("💡 Explanation:"),
		wrappedExplanation,
	)

	return borderStyle.Render(contentStyle.Render(content))
}

func wordwrap(text string, width int) string {
	var wrapped []string
	words := strings.Fields(text)

	line := ""
	for _, word := range words {
		if len(line)+len(word)+1 > width {
			wrapped = append(wrapped, line)
			line = word
		} else {
			if line != "" {
				line += " "
			}
			line += word
		}
	}
	if line != "" {
		wrapped = append(wrapped, line)
	}

	return strings.Join(wrapped, "\n")
}

func main() {
	var errorMessage string

	width, _, err := term.GetSize(0)
	if err != nil {
		width = 80
	}
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
		width:            width,
	})
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app: %v\n", err)
		os.Exit(1)
	}
}

func lookupExplanation(err string) string {
	return errGenerator.InitOpenAI(err)
}
