// in internal/algorithms/smooth_sort.go
package algorithms



// Name returns the algorithm's display name.
func (s *SmoothSorter) Name() string {
	return "SmoothSort (Placeholder)"
}

// Sort implements a placeholder.
func (s *SmoothSorter) Sort(arr []int, steps chan<- SortStep) {
	blackBoxPlaceholder(arr, steps)
}

