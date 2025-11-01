// in internal/algorithms/bubble_sort.go
package algorithms

import "time"



// Name returns the algorithm's name.
func (s *BubbleSorter) Name() string {
	return "Bubble Sort"
}

// Sort implements the Bubble Sort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *BubbleSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	n := len(arr)

	// Make a copy to sort in place without modifying the original slice
	localArr := make([]int, n)
	copy(localArr, arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			// Send comparison step
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{j, j + 1}}

			if localArr[j] > localArr[j+1] {
				// Swap
				localArr[j], localArr[j+1] = localArr[j+1], localArr[j]
				swapped = true

				// Send swap step
				steps <- SortStep{Array: CopySlice(localArr), Swap: []int{j, j + 1}}
			}
		}
		if !swapped {
			break
		}
	}

	duration := time.Since(start)

	// FINAL STEP UPDATED:
	// 1. Added the 'Duration' field to the final step.
	// 2. Closed the channel.
	// 3. Removed the 'return duration' statement.
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}

