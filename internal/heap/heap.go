package heap

import "container/heap"

type heapable interface{ Less(any) bool }
type generic[T heapable] []T

// boilerplate methods required by heap package
func (h generic[T]) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h generic[T]) Len() int           { return len(h) }
func (h generic[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *generic[T]) Push(x any)        { *h = append(*h, x.(T)) }
func (h *generic[T]) Pop() any          { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

// wrapper for generic heap, hides internal methods and exposses cleaner api
type Heap[T heapable] generic[T]

func MakeHeap[T heapable](elems ...T) Heap[T] {
	h := Heap[T](elems)
	h.Init()
	return h
}
func (h *Heap[T]) Init()          { heap.Init((*generic[T])(h)) }
func (h *Heap[T]) Pop() T         { return heap.Pop((*generic[T])(h)).(T) }
func (h *Heap[T]) Push(value T)   { heap.Push((*generic[T])(h), value) }
func (h *Heap[T]) Remove(i int) T { return heap.Remove((*generic[T])(h), i).(T) }
func (h *Heap[T]) Fix(i int)      { heap.Fix((*generic[T])(h), i) }
