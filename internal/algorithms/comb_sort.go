// in internal/algorithms/comb_sort.go
package algorithms

import (
	"math"
	"time"
)


// Name returns the algorithm's display name.
func (s *CombSorter) Name() string {
	return "Comb Sort"
}

// getNextGap calculates the shrinking gap.
func (s *CombSorter) getNextGap(gap int) int {
	// Shrink factor
	const shrink = 1.3
	newGap := int(math.Floor(float64(gap) / shrink))
	if newGap < 1 {
		return 1
	}
	return newGap
}

// Sort implements the Comb Sort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *CombSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	localArr := CopySlice(arr)
	n := len(localArr)

	gap := n
	swapped := true

	for gap != 1 || swapped {
		// Find next gap
		gap = s.getNextGap(gap)

		// Reset swapped flag
		swapped = false

		// A gapped Bubble Sort
		for i := 0; i < n-gap; i++ {
			// Send comparison
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{i, i + gap}}

			if localArr[i] > localArr[i+gap] {
				// Swap
				localArr[i], localArr[i+gap] = localArr[i+gap], localArr[i]
				swapped = true
				// Send swap
				steps <- SortStep{Array: CopySlice(localArr), Swap: []int{i, i + gap}}
			}
		}
	}

	duration := time.Since(start)

	// FINAL STEP UPDATED:
	// 1. Added the 'Duration' field.
	// 2. Closed the channel.
	// 3. Removed the 'return' statement.
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}