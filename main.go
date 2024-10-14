package main

import (
	"fmt"
	"time"

	"github.com/b0sc/sortRace/internal/algorithms"
)

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", arr)

	// start timer
	start := time.Now()
	algorithms.BubbleSort(arr)
	elapsed := time.Since(start)

	fmt.Println("Sorted array:", arr)
	fmt.Printf("BubbleSort took %s\n", elapsed)
}
