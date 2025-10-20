// in internal/visual/visual.go
package visual

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/b0sc/sortRace/internal/algorithms"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// --- Helper Functions ---

// GenerateData creates a slice of random integers.
// (You can expand this later with -state flags)
func GenerateData(size, maxVal int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = r.Intn(maxVal-1) + 1
	}
	return data
}

// --- Bubbletea Messages ---

// stepMsg reports a single step (compare or swap) from an algorithm.
type stepMsg struct {
	index int // Which algorithm sent this?
	step  algorithms.SortStep
}

// --- Model Definition ---

// --- RENAMED: from Model to RaceModel ---
type RaceModel struct {
	sorters     []algorithms.Sorter
	channels    []chan algorithms.SortStep
	states      [][]int
	names       []string
	comparisons map[int][]int
	swaps       map[int][]int
	done        []bool
	durations   []time.Duration
	quitting    bool
	maxVal      int
	ranks       []int
	finishCount int
}

// --- RENAMED: from NewModel to NewRaceModel ---
// --- MODIFIED: Now accepts 'selectedSorters' ---
func NewRaceModel(data []int, selectedSorters []algorithms.Sorter) RaceModel {
	// --- MODIFIED: Count is based on selection, not AllSorters ---
	sorterCount := len(selectedSorters)

	maxVal := 0
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}
	if maxVal == 0 {
		maxVal = 1
	}

	m := RaceModel{
		// --- MODIFIED: Use the 'selectedSorters' list ---
		sorters:     selectedSorters,
		channels:    make([]chan algorithms.SortStep, sorterCount),
		states:      make([][]int, sorterCount),
		names:       make([]string, sorterCount),
		comparisons: make(map[int][]int),
		swaps:       make(map[int][]int),
		done:        make([]bool, sorterCount),
		durations:   make([]time.Duration, sorterCount),
		quitting:    false,
		maxVal:      maxVal,
		ranks:       make([]int, sorterCount),
		finishCount: 0,
	}

	for i, sorter := range m.sorters {
		m.names[i] = sorter.Name()
		// --- MODIFIED: Use algorithms.CopySlice ---
		m.states[i] = algorithms.CopySlice(data)
		m.channels[i] = make(chan algorithms.SortStep)
	}
	return m
}

// --- Bubbletea Core Functions ---

// (This helper function is unchanged)
func waitForStep(ch chan algorithms.SortStep, index int) tea.Cmd {
	return func() tea.Msg {
		step, ok := <-ch
		if !ok {
			return nil
		}
		return stepMsg{index: index, step: step}
	}
}

// --- MODIFIED: Attached to 'RaceModel' ---
func (m RaceModel) Init() tea.Cmd {
	cmds := make([]tea.Cmd, len(m.sorters))
	for i := 0; i < len(m.sorters); i++ {
		go m.sorters[i].Sort(m.states[i], m.channels[i])
		cmds[i] = waitForStep(m.channels[i], i)
	}
	return tea.Batch(cmds...)
}

// --- MODIFIED: Attached to 'RaceModel' ---
func (m RaceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	case stepMsg:
		if m.done[msg.index] {
			return m, nil
		}

		m.states[msg.index] = msg.step.Array
		m.comparisons[msg.index] = msg.step.Comparison
		m.swaps[msg.index] = msg.step.Swap

		if msg.step.IsDone {
			m.done[msg.index] = true
			m.durations[msg.index] = msg.step.Duration
			m.finishCount++
			m.ranks[msg.index] = m.finishCount
			return m, nil
		}
		return m, waitForStep(m.channels[msg.index], msg.index)
	}
	return m, nil
}

// --- MODIFIED: Attached to 'RaceModel' ---
func (m RaceModel) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	var b strings.Builder
	b.WriteString("--- Race of Sorting Algorithms ---\n(Press 'q' to quit)\n\n")

	boldStyle := lipgloss.NewStyle().Bold(true)
	winnerStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("82"))
	compareStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	swapStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("197"))
	defaultStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	doneStyle := lipgloss.NewStyle().Faint(true)
	barChars := []string{" ", "▂", "▃", "▄", "▅", "▆", "▇", "█"}

	for i, name := range m.names {
		durationStr := " (Sorting...)"
		rankStr := ""
		if m.done[i] {
			durationStr = fmt.Sprintf(" (Done! %s)", m.durations[i])
			rankStr = fmt.Sprintf("#%d ", m.ranks[i])
		}

		header := fmt.Sprintf("%s%s%s", rankStr, name, durationStr)
		headerStyle := boldStyle
		if m.ranks[i] == 1 {
			headerStyle = winnerStyle
		}

		b.WriteString(headerStyle.Render(header) + "\n")
		b.WriteString(strings.Repeat("─", lipgloss.Width(header)) + "\n")

		var arrStr strings.Builder
		for j, val := range m.states[i] {
			style := defaultStyle
			if contains(m.swaps[i], j) {
				style = swapStyle
			} else if contains(m.comparisons[i], j) {
				style = compareStyle
			}
			if m.done[i] {
				style = style.Copy().Inherit(doneStyle)
			}

			barIndex := (val * (len(barChars) - 1)) / m.maxVal
			if barIndex < 0 {
				barIndex = 0
			}
			if barIndex >= len(barChars) {
				barIndex = len(barChars) - 1
			}

			bar := barChars[barIndex]
			arrStr.WriteString(style.Render(bar))
		}
		b.WriteString(arrStr.String())
		b.WriteString("\n\n")
	}

	return b.String()
}

// (This helper function is unchanged)
func contains(slice []int, val int) bool {
	if slice == nil {
		return false
	}
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}