// in main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/b0sc/sortRace/internal/algorithms"
	"github.com/b0sc/sortRace/internal/selection" // --- ADDED ---
	"github.com/b0sc/sortRace/internal/visual"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// --- (Optional) Add your flags here ---
	size := flag.Int("size", 70, "Number of elements in the array")
	max := flag.Int("max", 100, "Maximum value for random elements")
	flag.Parse()

	// 1. Run the SELECTION screen first
	selModel := selection.NewSelectionModel(algorithms.AllSorters)
	selProg := tea.NewProgram(selModel)
	
	// p.Run() returns the *final* model state
	finalModel, err := selProg.Run()
	if err != nil {
		fmt.Println("Error running selection:", err)
		os.Exit(1)
	}

	// Cast the final model to get our data
	selModel, ok := finalModel.(selection.Model)
	if !ok {
		fmt.Println("Error casting model")
		os.Exit(1)
	}

	// 2. Check the results from the selection
	// If user quit, or selected no algorithms, exit.
	if selModel.Quitting || len(selModel.FinalSelection) == 0 {
		fmt.Println("Race cancelled.")
		os.Exit(0)
	}

	// 3. We have a selection! Now, run the RACE screen
	data := visual.GenerateData(*size, *max)
	
	// Pass the selection (selModel.FinalSelection) to the RaceModel
	raceModel := visual.NewRaceModel(data, selModel.FinalSelection)
	raceProg := tea.NewProgram(raceModel, tea.WithAltScreen())

	if _, err := raceProg.Run(); err != nil {
		fmt.Println("Error running race:", err)
		os.Exit(1)
	}
}