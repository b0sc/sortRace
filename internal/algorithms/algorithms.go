// in internal/algorithms/algorithms.go
package algorithms

import (
	"sort"
	"time"
)

// SortStep represents a single action taken by the algorithm.
// This is what we will send to the visualizer.
type SortStep struct {
	Array      []int // The current state of the array
	Comparison []int // Indices being compared, e.g., [0, 1]
	Swap       []int // Indices being swapped, e.g., [0, 1]
	IsDone     bool  // Flag to signal completion
	Duration  time.Duration // Time taken for the sort (only in final step)
}

// Sorter is the interface that all sorting algorithms must implement.
type Sorter interface {
	// Name returns the algorithm's display name.
	Name() string

	// Sort performs the sort, sending visualization steps to the 'steps' channel.
	// It also returns the total duration.
	Sort(arr []int, steps chan<- SortStep) 
}

// CopySlice creates a deep copy of an integer slice.
// This is crucial to prevent multiple goroutines (algorithms)
// from modifying the same data.
func CopySlice(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}

// BlackBoxPlaceholder runs Go's standard sort as a placeholder.
func BlackBoxPlaceholder(arr []int, steps chan<- SortStep) {
	localArr := CopySlice(arr)
	steps <- SortStep{Array: CopySlice(localArr)} // Initial state

	start := time.Now()
	sort.Slice(localArr, func(i, j int) bool {
		return localArr[i] < localArr[j]
	})
	duration := time.Since(start)

	steps <- SortStep{Array: localArr, IsDone: true, Duration: duration} // Final state
	close(steps)
}

// --- Struct Definitions for all Sorters ---
// (Each of these will have its implementation in a separate file,
// e.g., bubble_sort.go, quick_sort.go)

type BubbleSorter struct{}
type SelectionSorter struct{}
type InsertionSorter struct{}
type HeapSorter struct{}
type IntroSorter struct{}
type MergeSorter struct{}
type InPlaceMergeSorter struct{}
type TournamentSorter struct{}
type TreeSorter struct{}
type BlockSorter struct{}
type SmoothSorter struct{}
type TimSorter struct{}
type PatienceSorter struct{}
type CubeSorter struct{}
type QuickSorter struct{}
type FluxSorter struct{}
type CrumSorter struct{}
type LibrarySorter struct{} // Wrapper for Go's stdlib sort
type ShellSorter struct{}
type CombSorter struct{}
type CocktailShakerSorter struct{}
type GnomeSorter struct{}
type OddEvenSorter struct{}
type StrandSorter struct{}
type ExchangeSorter struct{} // Note: BubbleSort is a type of ExchangeSort
type CycleSorter struct{}

// --- Sorter Registry ---

// AllSorters is a slice containing an instance of every implemented sorter.
// This is what the main application will use to get the list of algorithms.
var AllSorters = []Sorter{
	// Simple Sorts
	&BubbleSorter{},
	&SelectionSorter{},
	&InsertionSorter{},
	&GnomeSorter{},
	&CocktailShakerSorter{},


	&CombSorter{}, // An evolution of Bubble Sort
	&ShellSorter{}, // An evolution of Insertion Sort


	// Efficient Sorts
	&HeapSorter{},
	&MergeSorter{},
	&QuickSorter{},
	&CycleSorter{},
	&ExchangeSorter{},
	&OddEvenSorter{},
	&StrandSorter{},
	&LibrarySorter{},
	
	// Novel Sorts
	&PatienceSorter{},
	&TreeSorter{},
	
	// Advanced Sorts
	&IntroSorter{},
	&TimSorter{},
	&TournamentSorter{},
	&InPlaceMergeSorter{},
	&BlockSorter{},	
	&SmoothSorter{},
	&CubeSorter{},
	&FluxSorter{},
	&CrumSorter{},






}

// --- Example Implementation (Go's Library Sort) ---
// You can put this in its own file: 'library_sort.go'

// Name returns the algorithm's name.
func (s *LibrarySorter) Name() string {
	return "Library Sort (Go stdlib)"
}

// Sort implements the Go standard library sort.
// Note: This one can't be visualized easily as it's a black box.
// We just send the initial and final state.
func (s *LibrarySorter) Sort(arr []int, steps chan<- SortStep) {
	localArr := CopySlice(arr)
	steps <- SortStep{Array: CopySlice(localArr)} // Send initial state

	start := time.Now()
	sort.Ints(localArr) // The actual sort
	duration := time.Since(start)

	// Send final state
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}

