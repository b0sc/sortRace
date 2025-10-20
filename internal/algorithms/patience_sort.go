// in internal/algorithms/patience_sort.go
package algorithms

import (
	"container/heap"
	"sort"
	"time"
)

// Name returns the algorithm's display name.
func (s *PatienceSorter) Name() string {
	return "Patience Sort"
}

// Sort implements the Patience Sort algorithm.
// This is a "black box" implementation for racing.
func (s *PatienceSorter) Sort(arr []int, steps chan<- SortStep) {
	localArr := CopySlice(arr)
	
	// Send initial state
	steps <- SortStep{Array: CopySlice(localArr)}
	start := time.Now()

	// 1. Build the piles
	var piles [][]int
	for _, val := range localArr {
		// Find the first pile where val <= pile's top
		pileIndex := sort.Search(len(piles), func(i int) bool {
			return piles[i][len(piles[i])-1] >= val
		})

		if pileIndex < len(piles) {
			// Add to existing pile
			piles[pileIndex] = append(piles[pileIndex], val)
		} else {
			// Start a new pile
			piles = append(piles, []int{val})
		}
	}

	// 2. Merge the piles using a min-heap
	sortedArr := make([]int, 0, len(localArr))
	pq := make(pileHeap, 0, len(piles))

	// Initialize heap with the top (last) card of each pile
	for i := 0; i < len(piles); i++ {
		heap.Push(&pq, &pileItem{
			val:       piles[i][len(piles[i])-1],
			pileIndex: i,
		})
		// Pop the card from the pile
		piles[i] = piles[i][:len(piles[i])-1]
	}

	// 3. Repeatedly pull from the heap
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pileItem)
		sortedArr = append(sortedArr, item.val)
		
		// If the item's original pile is not empty, push its new top
		if len(piles[item.pileIndex]) > 0 {
			pile := piles[item.pileIndex]
			heap.Push(&pq, &pileItem{
				val:       pile[len(pile)-1],
				pileIndex: item.pileIndex,
			})
			piles[item.pileIndex] = pile[:len(pile)-1]
		}
	}

	duration := time.Since(start)

	// Send final state
	steps <- SortStep{Array: sortedArr, IsDone: true, Duration: duration}
	close(steps)
}


// --- Min-Heap Implementation for Patience Sort ---

// pileItem is an item in our priority queue
type pileItem struct {
	val       int // The value of the item
	pileIndex int // The pile this item came from
	index     int // The index of the item in the heap
}

// pileHeap implements heap.Interface
type pileHeap []*pileItem

func (ph pileHeap) Len() int { return len(ph) }

func (ph pileHeap) Less(i, j int) bool {
	// Min-heap, so we want the smallest value
	return ph[i].val < ph[j].val
}

func (ph pileHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
	ph[i].index = i
	ph[j].index = j
}

func (ph *pileHeap) Push(x interface{}) {
	n := len(*ph)
	item := x.(*pileItem)
	item.index = n
	*ph = append(*ph, item)
}

func (ph *pileHeap) Pop() interface{} {
	old := *ph
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*ph = old[0 : n-1]
	return item
}