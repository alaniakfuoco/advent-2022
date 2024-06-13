package util

import "fmt"

type Stack[T interface{}] struct {
	Size     int
	capacity int
	data     []T
}

func NewStack[T interface{}](cap int) *Stack[T] {
	return &Stack[T]{
		Size:     0,
		capacity: cap,
		data:     make([]T, cap),
	}
}

func (q *Stack[T]) Push(item T) {
	if q.capacity == q.Size {
		q.increaseSize()
	}

	q.data[q.Size] = item
	q.Size++
}

func (q *Stack[T]) Pop() T {
	result := q.data[q.Size-1]
	q.Size--
	return result
}

func (q *Stack[T]) increaseSize() {
	newData := make([]T, q.capacity*2) // create array with double the size
	for i := 0; i < q.Size; i++ {
		newData[i] = q.data[i] // copy data over
	}
	q.data = newData            // assign new array
	q.capacity = q.capacity * 2 // increase capacity
}

func (q *Stack[T]) Print() {
	fmt.Println(q.data[:q.Size])
}
