// in internal/algorithms/quick_sort.go
package algorithms

import "time"



// Name returns the algorithm's display name.
func (s *QuickSorter) Name() string {
	return "QuickSort"
}

// Sort implements the QuickSort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *QuickSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	localArr := CopySlice(arr)

	// Call the recursive helper
	s.quickSortRecursive(localArr, 0, len(localArr)-1, steps)

	duration := time.Since(start)

	// FINAL STEP UPDATED:
	// 1. Added the 'Duration' field.
	// 2. Closed the channel.
	// 3. Removed the 'return' statement.
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}

// partition is the helper function that performs the pivot and swap logic.
func (s *QuickSorter) partition(arr []int, low, high int, steps chan<- SortStep) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		// Report comparison against the pivot
		steps <- SortStep{Array: CopySlice(arr), Comparison: []int{j, high}}
		if arr[j] < pivot {
			i++
			// Report swap
			arr[i], arr[j] = arr[j], arr[i]
			steps <- SortStep{Array: CopySlice(arr), Swap: []int{i, j}}
		}
	}

	// Final swap for the pivot
	arr[i+1], arr[high] = arr[high], arr[i+1]
	steps <- SortStep{Array: CopySlice(arr), Swap: []int{i + 1, high}}

	return i + 1
}

// quickSortRecursive is the helper that handles the recursion.
func (s *QuickSorter) quickSortRecursive(arr []int, low, high int, steps chan<- SortStep) {
	if low < high {
		// Find pivot index
		pi := s.partition(arr, low, high, steps)

		// Recursively sort elements before and after partition
		s.quickSortRecursive(arr, low, pi-1, steps)
		s.quickSortRecursive(arr, pi+1, high, steps)
	}
}