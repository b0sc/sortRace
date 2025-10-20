// in internal/algorithms/shell_sort.go
package algorithms

import "time"



// Name returns the algorithm's display name.
func (s *ShellSorter) Name() string {
	return "ShellSort"
}

// Sort implements the ShellSort algorithm and sends steps.
// SIGNATURE UPDATED: Removed the 'time.Duration' return value.
func (s *ShellSorter) Sort(arr []int, steps chan<- SortStep) {
	start := time.Now()
	localArr := CopySlice(arr)
	n := len(localArr)

	// Start with a large gap, then reduce the gap
	// (Using Knuth's sequence: h = 3*h + 1)
	gap := 1
	for gap < n/3 {
		gap = 3*gap + 1
	}

	for gap >= 1 {
		// Do a gapped insertion sort for this gap size.
		for i := gap; i < n; i++ {
			temp := localArr[i]
			j := i

			for ; j >= gap; j -= gap {
				// Send comparison
				steps <- SortStep{Array: CopySlice(localArr), Comparison: []int{j - gap, j}}

				if localArr[j-gap] > temp {
					localArr[j] = localArr[j-gap]
					// Send "shift" as a swap
					steps <- SortStep{Array: CopySlice(localArr), Swap: []int{j, j - gap}}
				} else {
					break
				}
			}

			// put temp (the original localArr[i]) in its correct location
			if j != i {
				localArr[j] = temp
				steps <- SortStep{Array: CopySlice(localArr), Swap: []int{j, i}} // Visualizing final placement
			}
		}
		// reduce the gap
		gap = gap / 3
	}

	duration := time.Since(start)

	// FINAL STEP UPDATED:
	// 1. Added the 'Duration' field.
	// 2. Closed the channel.
	// 3. Removed the 'return' statement.
	steps <- SortStep{Array: CopySlice(localArr), IsDone: true, Duration: duration}
	close(steps)
}