// in internal/algorithms/gnome_sort.go
package algorithms

import "time"

// Name returns the algorithm's display name.
func (s *GnomeSorter) Name() string {
	return "Gnome Sort"
}

// Sort implements the Gnome Sort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *GnomeSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
    // FIX: Removed extra '='
	localArr := CopySlice(arr)
	n := len(localArr)

	index := 0
	for index < n {
		if index == 0 {
			index++
		}

		// Send comparison
        // FIX: Removed extra '='
		steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{index - 1, index}}

		if localArr[index] >= localArr[index-1] {
			// Things are in order, move forward
			index++
		} else {
			// Swap and move backward
			localArr[index], localArr[index-1] = localArr[index-1], localArr[index]
            // FIX: Removed extra '='
			steps <- SortStep{Array: CopySlice(localArr), Swap: []int{index - 1, index}}
			index--
		}
	}

	duration := time.Since(start)

	// FINAL STEP UPDATED:
	// 1. Added the 'Duration' field.
	// 2. Closed the channel.
	// 3. Removed the 'return' statement.
    // FIX: Removed extra '='
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}