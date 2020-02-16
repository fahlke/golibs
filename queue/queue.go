package queue

import (
	"errors"
	"sync"
)

// Queue represents a queue implementation
type Queue struct {
	mutex sync.RWMutex
	data  []interface{}
}

// ErrEmptyQueue will be returned on an invalid operation on an empty queue.
var ErrEmptyQueue = errors.New("empty queue")

// Empty checks whether the underlying queue is empty.
func (q *Queue) Empty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return !(len(q.data) > 0)
}

// Size returns the number of elements in the underlying queue.
func (q *Queue) Size() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return len(q.data)
}

// Push inserts an element at the end.
func (q *Queue) Push(item interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.data = append(q.data, item)
}

// Pop removes the first element from the queue.
func (q *Queue) Pop() (interface{}, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.data) > 0 {
		item := q.data[0]
		q.data = q.data[1:]

		return item, nil
	}

	return nil, ErrEmptyQueue
}

// Front returns the first element in the queue without removing it.
func (q *Queue) Front() (interface{}, error) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	if len(q.data) > 0 {
		return q.data[0], nil
	}

	return nil, ErrEmptyQueue
}

// Back returns the last element in the queue without removing it.
func (q *Queue) Back() (interface{}, error) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	if len(q.data) > 0 {
		return q.data[len(q.data)-1], nil
	}

	return nil, ErrEmptyQueue
}
