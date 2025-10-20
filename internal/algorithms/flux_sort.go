// in internal/algorithms/flux_sort.go
package algorithms



// Name returns the algorithm's display name.
func (s *FluxSorter) Name() string {
	return "FluxSort (Placeholder)"
}

// Sort implements a placeholder.
func (s *FluxSorter) Sort(arr []int, steps chan<- SortStep) {
	blackBoxPlaceholder(arr, steps)
}

