package utilities

type PriorityElement[T comparable] struct {
	Value    T
	Priority int
}

type MaxPriorityQueue[T comparable] struct {
	priorityQueue[T]
}

func (mpq MaxPriorityQueue[_]) Less(i, j int) bool {
	return mpq.priorityQueue[i].Priority > mpq.priorityQueue[j].Priority
}

type MinPriorityQueue[T comparable] struct {
	priorityQueue[T]
}

func (mpq MinPriorityQueue[_]) Less(i, j int) bool {
	return mpq.priorityQueue[i].Priority < mpq.priorityQueue[j].Priority
}

type priorityQueue[T comparable] []PriorityElement[T]

func (pq priorityQueue[_]) Len() int {
	return len(pq)
}

func (pq priorityQueue[_]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue[T]) Push(x any) {
	item := x.(PriorityElement[T])
	*pq = append(*pq, item)
}

func (pq *priorityQueue[T]) Pop() any {
	size := len(*pq)
	item := (*pq)[size-1]
	*pq = (*pq)[:size-1]
	return item
}
