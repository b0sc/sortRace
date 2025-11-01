// in internal/algorithms/oddeven_sort.go
package algorithms

import "time"

// Name returns the algorithm's display name.
func (s *OddEvenSorter) Name() string {
	return "Odd-Even Sort"
}

// Sort implements the Odd-Even Sort algorithm and sends steps.
func (s *OddEvenSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	localArr := CopySlice(arr)
	n := len(localArr)

	sorted := false
	for !sorted {
		sorted = true

		// Odd phase (compare 1-2, 3-4, ...)
		for i := 1; i < n-1; i += 2 {
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{i, i + 1}}
			if localArr[i] > localArr[i+1] {
				localArr[i], localArr[i+1] = localArr[i+1], localArr[i]
				steps <- SortStep{Array: CopySlice(localArr), Swap: []int{i, i + 1}}
				sorted = false
			}
		}

		// Even phase (compare 0-1, 2-3, ...)
		for i := 0; i < n-1; i += 2 {
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{i, i + 1}}
			if localArr[i] > localArr[i+1] {
				localArr[i], localArr[i+1] = localArr[i+1], localArr[i]
				steps <- SortStep{Array: CopySlice(localArr), Swap: []int{i, i + 1}}
				sorted = false
			}
		}
	}

	duration := time.Since(start)
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}