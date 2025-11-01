// in internal/algorithms/cycle_sort.go
package algorithms

import "time" // <-- 1. MAKE SURE THIS IMPORT IS HERE

// Name returns the algorithm's display name.
func (s *CycleSorter) Name() string {
	return "Cycle Sort"
}

// Sort implements the Cycle Sort algorithm and sends steps.
func (s *CycleSorter) Sort(arr []int, steps chan<- SortStep) {
	// 2. FIX: 'Time.Now()' changed to 'time.Now()'
	start := time.Now()
	localArr := CopySlice(arr)
	n := len(localArr)

	// Traverse array elements
	for cycleStart := 0; cycleStart < n-1; cycleStart++ {
		item := localArr[cycleStart]

		// Find where to put the item
		pos := cycleStart
		for i := cycleStart + 1; i < n; i++ {
			steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{cycleStart, i}}
			if localArr[i] < item {
				pos++
			}
		}

		// If item is already in correct place, skip
		if pos == cycleStart {
			continue
		}

		// Skip duplicates
		for item == localArr[pos] {
			pos++
		}

		// Put the item in its correct place
		localArr[pos], item = item, localArr[pos]
		steps <- SortStep{Array: CopySlice(localArr), Swap: []int{cycleStart, pos}}

		// Rotate the rest of the cycle
		for pos != cycleStart {
			pos = cycleStart
			// Find where to put the displaced item
			for i := cycleStart + 1; i < n; i++ {
				steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{cycleStart, i}}
				if localArr[i] < item {
					pos++
				}
			}

			// Skip duplicates
			for item == localArr[pos] {
				pos++
			}

			// Put the item in its correct place
			localArr[pos], item = item, localArr[pos]
			steps <- SortStep{Array: CopySlice(localArr), Swap: []int{cycleStart, pos}}
		}
	}

	// 3. FIX: 'Time.Since(start)' changed to 'time.Since(start)'
	duration := time.Since(start)
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}