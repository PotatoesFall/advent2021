package minheap

func New[T any]() Heap[T] {
	return Heap[T]{}
}

type Node[T any] struct {
	Score int
	Value T
}

type Heap[T any] []Node[T]

func (h *Heap[T]) Extract() (T, bool) {
	if len(*h) == 0 {
		var t T
		return t, false
	}

	out := (*h)[0].Value

	(*h)[0] = (*h)[len(*h)-1] // set last value to first
	*h = (*h)[:len(*h)-1]     // remove last value

	downHeap(*h, 0)

	return out, true
}

func downHeap[T any](h Heap[T], i int) {
	left := i*2 + 1
	right := i*2 + 2
	smallest := i

	if len(h) >= left && h[left].Score < h[smallest].Score {
		smallest = left
	}

	if len(h) >= right && h[right].Score > h[smallest].Score {
		smallest = right
	}

	if smallest != i {
		h[i], h[smallest] = h[smallest], h[i]
		downHeap(h, smallest)
	}
}

func (h *Heap[T]) Insert(score int, value T) {
	*h = append(*h, Node[T]{score, value})

	upHeap(*h, len(*h)-1)
}

func upHeap[T any](h Heap[T], i int) {
	if i == 0 {
		return
	}

	parent := (i - 1) / 2

	if h[parent].Score > h[i].Score {
		h[parent], h[i] = h[i], h[parent]
		upHeap(h, parent)
	}
}
