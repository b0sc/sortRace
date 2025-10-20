// in internal/algorithms/exchange_sort.go
package algorithms

import "time"

// Name returns the algorithm's display name.
func (s *ExchangeSorter) Name() string {
	return "Exchange Sort (Naive)"
}

// Sort implements the Exchange Sort algorithm and sends steps.
func (s *ExchangeSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	localArr := CopySlice(arr)
	n := len(localArr)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			// Send comparison
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{i, j}}
			if localArr[i] > localArr[j] {
				// Swap
				localArr[i], localArr[j] = localArr[j], localArr[i]
				steps <- SortStep{Array: CopySlice(localArr), Swap: []int{i, j}}
			}
		}
	}

	duration := time.Since(start)
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}