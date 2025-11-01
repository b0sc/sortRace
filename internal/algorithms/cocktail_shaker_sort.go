// in internal/algorithms/cocktail_shaker_sort.go
package algorithms

import "time"



// Name returns the algorithm's display name.
func (s *CocktailShakerSorter) Name() string {
	return "Cocktail Shaker Sort"
}

// Sort implements the Cocktail Shaker Sort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *CocktailShakerSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	localArr := CopySlice(arr)
	n := len(localArr)

	begin := 0
	end := n - 1
	swapped := true

	for swapped {
		// Reset swap flag on entering new loop
		swapped = false

		// --- Pass 1: Forward Bubble Sort ---
		for i := begin; i < end; i++ {
			// Send comparison
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{i, i + 1}}
			if localArr[i] > localArr[i+1] {
				localArr[i], localArr[i+1] = localArr[i+1], localArr[i]
				swapped = true
				// Send swap
				steps <- SortStep{Array: CopySlice(localArr), Swap: []int{i, i + 1}}
			}
		}

		// If nothing moved, then array is sorted
		if !swapped {
			break
		}

		// Mark the last element as sorted
		end--

		// --- Pass 2: Backward Bubble Sort ---
		// Reset swap flag for backward pass
		swapped = false

		for i := end - 1; i >= begin; i-- {
			// Send comparison
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{i, i + 1}}
			if localArr[i] > localArr[i+1] {
				localArr[i], localArr[i+1] = localArr[i+1], localArr[i]
				swapped = true
				// Send swap
				steps <- SortStep{Array: CopySlice(localArr), Swap: []int{i, i + 1}}
			}
		}

		// Mark the first element as sorted
		begin++
	}

	duration := time.Since(start)

	// FINAL STEP UPDATED:
	// 1. Added the 'Duration' field.
	// 2. Closed the channel.
	// 3. Removed the 'return' statement.
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}