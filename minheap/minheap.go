package minheap

func New[T any]() Heap[T] {
	return Heap[T]{}
}

type Node[T any] struct {
	Score int
	Value T
}

type Heap[T any] struct {
	Tree []Node[T]
}

func (h Heap[T]) Len() int {
	return len(h.Tree)
}

func (h *Heap[T]) Extract() (T, bool) {
	if len(h.Tree) == 0 {
		var t T
		return t, false
	}

	out := h.Tree[0].Value

	h.Tree[0] = h.Tree[len(h.Tree)-1] // set last value to first
	h.Tree = h.Tree[:len(h.Tree)-1]   // remove last value

	downHeap(h.Tree, 0)

	return out, true
}

func downHeap[T any](t []Node[T], i int) {
	left := i*2 + 1
	right := i*2 + 2
	smallest := i

	if len(t) > left && t[left].Score < t[smallest].Score {
		smallest = left
	}

	if len(t) > right && t[right].Score > t[smallest].Score {
		smallest = right
	}

	if smallest != i {
		t[i], t[smallest] = t[smallest], t[i]
		downHeap(t, smallest)
	}
}

func (h *Heap[T]) Insert(score int, value T) {
	h.Tree = append(h.Tree, Node[T]{score, value})

	upHeap(*h, len(h.Tree)-1)
}

func upHeap[T any](h Heap[T], i int) {
	if i == 0 {
		return
	}

	parent := (i - 1) / 2

	if h.Tree[parent].Score > h.Tree[i].Score {
		h.Tree[parent], h.Tree[i] = h.Tree[i], h.Tree[parent]
		upHeap(h, parent)
	}
}
