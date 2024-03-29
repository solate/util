// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates an integer heap built using the heap interface.
package heap

import (
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(v int) {
	*h = append(*h, v)
}

func (h *IntHeap) Pop() int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Example_intHeap() {
	myHeap := &IntHeap{2, 1, 5}
	var h Interface[int] = myHeap
	Init(h)
	Push(h, 3)
	fmt.Printf("minimum: %d\n", (*myHeap)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}
