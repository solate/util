package heap

import (
	"math/rand"
	"testing"
)

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v int) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v int) {
	*h = append(*h, v)
}

func (h myHeap) verify(t *testing.T, i int) {
	t.Helper()
	n := h.Len()
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if h.Less(j1, i) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i], j1, h[j1])
			return
		}
		h.verify(t, j1)
	}
	if j2 < n {
		if h.Less(j2, i) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i], j1, h[j2])
			return
		}
		h.verify(t, j2)
	}
}

func TestInit0(t *testing.T) {
	my := new(myHeap)
	var h Interface[int] = my
	for i := 20; i > 0; i-- {
		h.Push(0) // all elements are the same
	}
	Init(h)
	my.verify(t, 0)

	for i := 1; h.Len() > 0; i++ {
		x := Pop(h)
		my.verify(t, 0)
		t.Log(h)
		if x != 0 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 0)
		}
	}
}

func Test(t *testing.T) {
	my := new(myHeap)
	var h Interface[int] = my
	my.verify(t, 0)

	for i := 20; i > 10; i-- {
		h.Push(i)
	}
	Init(h)
	t.Log(h)
	my.verify(t, 0)

	for i := 10; i > 0; i-- {
		Push(h, i)
		my.verify(t, 0)
	}

	for i := 1; h.Len() > 0; i++ {
		x := Pop(h)
		if i < 20 {
			Push(h, 20+i)
		}
		my.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestRemove0(t *testing.T) {
	my := new(myHeap)
	var h Interface[int] = my
	for i := 0; i < 10; i++ {
		h.Push(i)
	}
	my.verify(t, 0)

	for h.Len() > 0 {
		i := h.Len() - 1
		x := Remove(h, i).(int)
		if x != i {
			t.Errorf("Remove(%d) got %d; want %d", i, x, i)
		}
		my.verify(t, 0)
	}
}

func TestRemove1(t *testing.T) {
	my := new(myHeap)
	var h Interface[int] = my
	for i := 0; i < 10; i++ {
		h.Push(i)
	}
	my.verify(t, 0)

	for i := 0; h.Len() > 0; i++ {
		x := Remove(h, 0).(int)
		if x != i {
			t.Errorf("Remove(0) got %d; want %d", x, i)
		}
		my.verify(t, 0)
	}
}

func TestRemove2(t *testing.T) {
	N := 10

	my := new(myHeap)
	var h Interface[int] = my
	for i := 0; i < N; i++ {
		h.Push(i)
	}
	my.verify(t, 0)

	m := make(map[int]bool)
	for h.Len() > 0 {
		m[Remove(h, (h.Len()-1)/2).(int)] = true
		my.verify(t, 0)
	}

	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	for i := 0; i < len(m); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}

func BenchmarkDup(b *testing.B) {
	const n = 10000
	my := make(myHeap, 0, n)
	var h Interface[int] = &my
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			Push(h, 0) // all elements are the same
		}
		for h.Len() > 0 {
			Pop(h)
		}
	}
}

func TestFix(t *testing.T) {
	my := new(myHeap)
	var h Interface[int] = my
	my.verify(t, 0)

	for i := 200; i > 0; i -= 10 {
		Push(h, i)
	}
	my.verify(t, 0)

	if (*my)[0] != 10 {
		t.Fatalf("Expected head to be 10, was %d", (*my)[0])
	}
	(*my)[0] = 210
	Fix(h, 0)
	my.verify(t, 0)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(h.Len())
		if i&1 == 0 {
			(*my)[elem] *= 2
		} else {
			(*my)[elem] /= 2
		}
		Fix(h, elem)
		my.verify(t, 0)
	}
}
