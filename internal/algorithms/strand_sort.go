// in internal/algorithms/strand_sort.go
package algorithms

import "time"

// Name returns the algorithm's display name.
func (s *StrandSorter) Name() string {
	return "Strand Sort"
}

// Sort implements the Strand Sort algorithm and sends steps.
func (s *StrandSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	
	// StrandSort is not in-place. It builds a new list.
	// We'll use localArr as the "unsorted" list.
	localArr := CopySlice(arr)
	sortedList := make([]int, 0)

	for len(localArr) > 0 {
		// 1. Create a strand
		strand := make([]int, 0)
		strand = append(strand, localArr[0]) // Add first element
		localArr = localArr[1:] // Remove it from unsorted
		
		// Visualize the start of the strand
		steps <- SortStep{Array: buildVisual(sortedList, strand, localArr)}

		// 2. Build the strand
		remainingUnsorted := make([]int, 0)
		for _, val := range localArr {
			steps <- SortStep{Array: buildVisual(sortedList, strand, localArr), Comparison: []int{len(sortedList) + len(strand) - 1, len(sortedList) + len(strand)}}
			if val > strand[len(strand)-1] {
				// Add to strand
				strand = append(strand, val)
			} else {
				// Add to the *next* unsorted list
				remainingUnsorted = append(remainingUnsorted, val)
			}
		}
		localArr = remainingUnsorted

		// 3. Merge strand into sortedList
		sortedList = s.merge(sortedList, strand, steps, localArr)
		
		// Visualize the final merge
		steps <- SortStep{Array: buildVisual(sortedList, nil, localArr)}
	}

	duration := time.Since(start)
	steps <- SortStep{Array: sortedList, IsDone: true, Duration: duration}
	close(steps)
}

// merge helper for Strand Sort
func (s *StrandSorter) merge(listA, listB []int, steps chan<- SortStep, unsorted []int) []int {
	result := make([]int, 0, len(listA)+len(listB))
	i, j := 0, 0

	for i < len(listA) && j < len(listB) {
		// Visualize the merge comparison
		steps <- SortStep{Array: buildVisual(result, nil, unsorted), Comparison: []int{i, len(listA) + j}}
		
		if listA[i] < listB[j] {
			result = append(result, listA[i])
			i++
		} else {
			result = append(result, listB[j])
			j++
		}
	}

	// Add remaining elements
	for i < len(listA) {
		result = append(result, listA[i])
		i++
	}
	for j < len(listB) {
		result = append(result, listB[j])
		j++
	}
	
	return result
}

// buildVisual creates a single array slice for visualization
func buildVisual(sorted, strand, unsorted []int) []int {
	vis := make([]int, 0, len(sorted)+len(strand)+len(unsorted))
	vis = append(vis, sorted...)
	vis = append(vis, strand...)
	vis = append(vis, unsorted...)
	return vis
}