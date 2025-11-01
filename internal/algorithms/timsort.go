// in internal/algorithms/timsort.go
package algorithms

import (
	"sort"
	"time"
)

// Name returns the algorithm's display name.
func (s *TimSorter) Name() string {
	return "TimSort/MergeSort (Go stdlib Stable)"
}

// Sort implements Go's standard stable sort (a TimSort/MergeSort hybrid).
// This is a "black box" implementation for racing, not visualization.
func (s *TimSorter) Sort(arr []int, steps chan<- SortStep) {
	localArr := CopySlice(arr)

	// Send initial state
	steps <- SortStep{Array: CopySlice(localArr)}

	start := time.Now()
	// This is Go's implementation of a stable sort (like TimSort)
	sort.SliceStable(localArr, func(i, j int) bool {
		return localArr[i] < localArr[j]
	})
	duration := time.Since(start)

	// Send final state
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}