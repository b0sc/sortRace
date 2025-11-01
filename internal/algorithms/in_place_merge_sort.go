// in internal/algorithms/in_place_merge_sort.go
package algorithms



// Name returns the algorithm's display name.
func (s *InPlaceMergeSorter) Name() string {
	return "In-Place Merge Sort (Placeholder)"
}

// Sort implements a placeholder.
func (s *InPlaceMergeSorter) Sort(arr []int, steps chan<- SortStep) {
	blackBoxPlaceholder(arr, steps)
}

