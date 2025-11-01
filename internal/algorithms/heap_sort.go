// in internal/algorithms/heap_sort.go
package algorithms

import "time"



// Name returns the algorithm's display name.
func (s *HeapSorter) Name() string {
	return "HeapSort"
}

// Sort implements the HeapSort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *HeapSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	localArr := CopySlice(arr)
	n := len(localArr)

	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		s.heapify(localArr, n, i, steps)
	}

	// One by one extract an element from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		localArr[0], localArr[i] = localArr[i], localArr[0]
		steps <- SortStep{Array: CopySlice(localArr), Swap: []int{0, i}}

		// call max heapify on the reduced heap
		s.heapify(localArr, i, 0, steps)
	}

	duration := time.Since(start)

	// FINAL STEP UPDATED:
	// 1. Added the 'Duration' field.
	// 2. Closed the channel.
	// 3. Removed the 'return' statement.
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}

// To heapify a subtree rooted with node i which is an index in arr[].
// n is size of heap.
func (s *HeapSorter) heapify(arr []int, n, i int, steps chan<- SortStep) {
	largest := i // Initialize largest as root
	left := 2*i + 1
	right := 2*i + 2

	// See if left child is larger than root
	if left < n {
		steps <- SortStep{Array: CopySlice(arr), Comparison: []int{left, largest}}
		if arr[left] > arr[largest] {
			largest = left
		}
	}

	// See if right child is larger than largest so far
	if right < n {
		steps <- SortStep{Array: CopySlice(arr), Comparison: []int{right, largest}}
		if arr[right] > arr[largest] {
			largest = right
		}
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		steps <- SortStep{Array: CopySlice(arr), Swap: []int{i, largest}}

		// Recursively heapify the affected sub-tree
		s.heapify(arr, n, largest, steps)
	}
}