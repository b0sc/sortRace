package algorithms

// BubbleSort sorts an integer slice using the Bubble Sort algorithm.
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Flag to check if a swap was made
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no elements were swapped, the array is already sorted
		if !swapped {
			break
		}
	}
}
