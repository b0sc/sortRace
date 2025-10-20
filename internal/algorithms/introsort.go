// in internal/algorithms/introsort.go
package algorithms

import (
	"sort"
	"time"
)

// Name returns the algorithm's display name.
func (s *IntroSorter) Name() string {
	return "IntroSort (Go stdlib Unstable)"
}

// Sort implements Go's standard unstable sort (IntroSort).
// This is a "black box" implementation for racing, not visualization.
func (s *IntroSorter) Sort(arr []int, steps chan<- SortStep) {
	localArr := CopySlice(arr)
	
	// Send initial state
	steps <- SortStep{Array: CopySlice(localArr)}

	start := time.Now()
	// This is Go's implementation of IntroSort
	sort.Slice(localArr, func(i, j int) bool {
		return localArr[i] < localArr[j]
	})
	duration := time.Since(start)

	// Send final state
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}