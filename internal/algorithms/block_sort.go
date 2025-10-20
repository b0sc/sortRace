// in internal/algorithms/block_sort.go
package algorithms

import (
	"sort"
	"time"
)

// Name returns the algorithm's display name.
func (s *BlockSorter) Name() string {
	return "Block Sort (Placeholder)"
}

// Sort implements a placeholder.
func (s *BlockSorter) Sort(arr []int, steps chan<- SortStep) {
	blackBoxPlaceholder(arr, steps)
}

// blackBoxPlaceholder runs Go's standard sort.
// (You must also copy the blackBoxPlaceholder function here)
func blackBoxPlaceholder(arr []int, steps chan<- SortStep) {
	localArr := CopySlice(arr)
	steps <- SortStep{Array: CopySlice(localArr)}
	start := time.Now()
	sort.Slice(localArr, func(i, j int) bool {
		return localArr[i] < localArr[j]
	})
	duration := time.Since(start)
	steps <- SortStep{Array: localArr, IsDone: true, Duration: duration}
	close(steps)
}