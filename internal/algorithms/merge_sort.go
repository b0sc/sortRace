// in internal/algorithms/merge_sort.go
package algorithms

import "time"



// Name returns the algorithm's display name.
func (s *MergeSorter) Name() string {
	return "Merge Sort"
}

// Sort implements the Merge Sort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *MergeSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	localArr := CopySlice(arr)

	// Call the recursive helper
	s.mergeSortRecursive(localArr, 0, len(localArr)-1, steps)

	duration := time.Since(start)

	// FINAL STEP UPDATED:
	// 1. Added the 'Duration' field.
	// 2. Closed the channel.
	// 3. Removed the 'return' statement.
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}

// merge is the helper function that merges two subarrays.
func (s *MergeSorter) merge(arr []int, left, mid, right int, steps chan<- SortStep) {
	n1 := mid - left + 1
	n2 := right - mid

	// Create temp arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temp arrays
	for i := 0; i < n1; i++ {
		L[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[mid+1+j]
	}

	// Merge the temp arrays back into arr[left..right]
	i := 0 // Initial index of first subarray
	j := 0 // Initial index of second subarray
	k := left // Initial index of merged subarray

	for i < n1 && j < n2 {
		// Report comparison between elements from temp arrays
		// We'll use their original indices for visualization
		steps <- SortStep{Array: CopySlice(arr), Comparison: []int{left + i, mid + 1 + j}}

		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		// Report the "move" from a temp array back to the main array
		steps <- SortStep{Array: CopySlice(arr), Swap: []int{k, k}} // Visualize the write at index k
		k++
	}

	// Copy remaining elements of L[], if any
	for i < n1 {
		arr[k] = L[i]
		steps <- SortStep{Array: CopySlice(arr), Swap: []int{k, k}}
		i++
		k++
	}

	// Copy remaining elements of R[], if any
	for j < n2 {
		arr[k] = R[j]
		steps <- SortStep{Array: CopySlice(arr), Swap: []int{k, k}}
		j++
		k++
	}
}

// mergeSortRecursive is the helper that handles the recursion.
func (s *MergeSorter) mergeSortRecursive(arr []int, left, right int, steps chan<- SortStep) {
	if left < right {
		// Same as (left+right)/2, but avoids overflow
		mid := left + (right-left)/2

		// Sort first and second halves
		s.mergeSortRecursive(arr, left, mid, steps)
		s.mergeSortRecursive(arr, mid+1, right, steps)

		// Merge the sorted halves
		s.merge(arr, left, mid, right, steps)
	}
}