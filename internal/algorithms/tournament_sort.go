// in internal/algorithms/tournament_sort.go
package algorithms

import (
	"container/heap"
	"time"
)

// Name returns the algorithm's display name.
func (s *TournamentSorter) Name() string {
	return "Tournament Sort"
}

// Sort implements the Tournament Sort algorithm.
// This is a "black box" implementation for racing.
func (s *TournamentSorter) Sort(arr []int, steps chan<- SortStep) {
	localArr := CopySlice(arr)
	
	// Send initial state
	steps <- SortStep{Array: CopySlice(localArr)}
	start := time.Now()

	// 1. Build the min-heap
	pq := make(tournamentHeap, len(localArr))
	for i, val := range localArr {
		pq[i] = val
	}
	heap.Init(&pq)

	// 2. Repeatedly pull the "winner" (smallest item)
	sortedArr := make([]int, 0, len(localArr))
	for pq.Len() > 0 {
		smallest := heap.Pop(&pq).(int)
		sortedArr = append(sortedArr, smallest)
	}

	duration := time.Since(start)

	// Send final state
	steps <- SortStep{Array: sortedArr, IsDone: true, Duration: duration}
	close(steps)
}

// --- Min-Heap Implementation for Tournament Sort ---
// (We just need a simple int heap)

type tournamentHeap []int

func (h tournamentHeap) Len() int           { return len(h) }
func (h tournamentHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h tournamentHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *tournamentHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *tournamentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}