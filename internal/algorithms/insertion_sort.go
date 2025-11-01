// in internal/algorithms/insertion_sort.go
package algorithms

import "time"



// Name returns the algorithm's display name.
func (s *InsertionSorter) Name() string {
	return "Insertion Sort"
}

// Sort implements the Insertion Sort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *InsertionSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	localArr := CopySlice(arr)
	n := len(localArr)

	for i := 1; i < n; i++ {
		key := localArr[i]
		j := i - 1

		// Send initial comparison for this key
		steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{j, i}}

		// Move elements of localArr[0..i-1] that are greater than key
		// to one position ahead of their current position
		for j >= 0 && localArr[j] > key {
			// Report the comparison that resulted in a move
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{j, i}}

			localArr[j+1] = localArr[j]

			// Report the "shift" (visualized as a swap)
			steps <- SortStep{Array: CopySlice(localArr), Swap: []int{j + 1, j}}
			j = j - 1
		}
		localArr[j+1] = key
	}

	duration := time.Since(start)

	// FINAL STEP UPDATED:
	// 1. Added the 'Duration' field.
	// 2. Closed the channel.
	// 3. Removed the 'return' statement.
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}