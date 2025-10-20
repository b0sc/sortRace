// in internal/algorithms/tree_sort.go
package algorithms

import "time"

// Node defines a node in the Binary Search Tree
type Node struct {
	val   int
	left  *Node
	right *Node
}

// insert adds a new value to the BST
func insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{val: val}
	}
	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		// Insert duplicates to the right
		root.right = insert(root.right, val)
	}
	return root
}

// inOrder traverses the tree and populates the sorted array
func inOrder(root *Node, sortedArr *[]int) {
	if root != nil {
		inOrder(root.left, sortedArr)
		*sortedArr = append(*sortedArr, root.val)
		inOrder(root.right, sortedArr)
	}
}

// Name returns the algorithm's display name.
func (s *TreeSorter) Name() string {
	return "Tree Sort"
}

// Sort implements the Tree Sort algorithm.
// This is a "black box" implementation for racing.
func (s *TreeSorter) Sort(arr []int, steps chan<- SortStep) {
	localArr := CopySlice(arr)
	
	// Send initial state
	steps <- SortStep{Array: CopySlice(localArr)}

	start := time.Now()

	// 1. Build the tree
	var root *Node
	for _, val := range localArr {
		root = insert(root, val)
	}

	// 2. Perform in-order traversal to get sorted list
	sortedArr := make([]int, 0, len(localArr))
	inOrder(root, &sortedArr)

	duration := time.Since(start)

	// Send final state
	steps <- SortStep{Array: sortedArr, IsDone: true, Duration: duration}
	close(steps)
}