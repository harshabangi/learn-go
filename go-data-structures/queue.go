package go_data_structures

type Queue[T any] []T

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Size() int {
	return len(*q)
}

func (q *Queue[T]) Dequeue() *T {
	if q.Empty() {
		return nil
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return &item
}

func (q *Queue[T]) Enqueue(s T) {
	*q = append(*q, s)
}
