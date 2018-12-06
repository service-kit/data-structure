package queue

import "errors"

type Queue struct {
	capacity int
	size     int
	front    int
	rear     int
	array    []interface{}
}

func (q *Queue) Init() {
	q.capacity = 100
	q.array = make([]interface{}, q.capacity)
	q.front = 0
	q.rear = 0
	q.size = 0
}

func (q *Queue) ensureCapacity() {
	q.capacity *= 2
	newArray := make([]interface{}, q.capacity)
	copy(newArray, q.array[q.front:])
	copy(newArray[q.size-q.rear:], q.array[:q.rear])
	q.array = newArray
}

func (q *Queue) Enqueue(data interface{}) {
	if q.size == q.capacity {
		q.ensureCapacity()
	}
	q.array[q.rear] = data
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

func (q *Queue) Dequeue() (data interface{}, err error) {
	if q.size == 0 {
		return nil, errors.New("queue empty")
	}
	data = q.array[q.front]
	q.array[q.front] = nil
	q.front = (q.front + 1) % q.capacity
	q.size--
	return
}

func (q *Queue) Front() interface{} {
	return q.array[q.front]
}

func (q *Queue) Rear() interface{} {
	return q.array[(q.rear-1)%q.capacity]
}

func (q *Queue) Len() int {
	return q.size
}

func (q *Queue) Capacity() int {
	return q.capacity
}
