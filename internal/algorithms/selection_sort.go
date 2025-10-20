// in internal/algorithms/selection_sort.go
package algorithms

import "time"



// Name returns the algorithm's name.
func (s *SelectionSorter) Name() string {
	return "Selection Sort"
}

// Sort implements the Selection Sort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *SelectionSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	n := len(arr)

	localArr := make([]int, n)
	copy(localArr, arr)

	for i := 0; i < n-1; i++ {
		// Find the minimum element in unsorted array
		minIdx := i
		for j := i + 1; j < n; j++ {
			// Send comparison step
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{j, minIdx}}
			if localArr[j] < localArr[minIdx] {
				minIdx = j
			}
		}

		// Swap the found minimum element with the first element
		if i != minIdx { // Only swap if minIdx is different
			localArr[i], localArr[minIdx] = localArr[minIdx], localArr[i]
			// Send swap step
			steps <- SortStep{Array: CopySlice(localArr), Swap: []int{i, minIdx}}
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