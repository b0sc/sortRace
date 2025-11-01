// in internal/selection/selection.go
package selection

import (
	"fmt"
	"strings"
	"github.com/b0sc/sortRace/internal/algorithms"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("201"))
	checkStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("82")) // Green
)

// Model holds the state of our selection list.
type Model struct {
	sorters   []algorithms.Sorter // All available sorters
	cursor    int                 // Which item our cursor is on
	selected  map[int]bool        // Which sorters are selected
	// quitting  bool                // True if user pressed 'q'
	Quitting bool // True if user pressed 'q'
	// This field holds the result after the model quits
	FinalSelection []algorithms.Sorter
}

// NewSelectionModel creates a new selection model.
func NewSelectionModel(sorters []algorithms.Sorter) Model {
	m := Model{
		sorters:  sorters,
		cursor:   0,
		selected: make(map[int]bool),
	}
	// By default, select all sorters
	for i := 0; i < len(sorters); i++ {
		m.selected[i] = true
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles user input.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			// m.quitting = true
			m.Quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.sorters)-1 {
				m.cursor++
			}

		case " ": // Spacebar toggles selection
			m.selected[m.cursor] = !m.selected[m.cursor]

		case "enter":
			// User is done. Build the FinalSelection list.
			for i, sorter := range m.sorters {
				if m.selected[i] {
					m.FinalSelection = append(m.FinalSelection, sorter)
				}
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the selection list.
func (m Model) View() string {
	if m.Quitting {
		return "Race cancelled.\n"
	}

	var b strings.Builder
	b.WriteString("--- Select Algorithms to Race ---\n\n")

	for i, sorter := range m.sorters {
		// Cursor
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}

		// Checkbox
		checked := checkStyle.Render("[ ]")
		if m.selected[i] {
			checked = checkStyle.Render("[x]")
		}

		// Name
		name := sorter.Name()
		if m.cursor == i {
			name = cursorStyle.Render(name)
		}

		b.WriteString(fmt.Sprintf("%s %s %s\n", cursor, checked, name))
	}

	b.WriteString("\n(Use up/down to move, space to toggle, enter to start, q to quit)\n")
	return b.String()
}