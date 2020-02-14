package queue

import (
	"sync"
)

type Queue struct {
	mutex sync.RWMutex
	data  []interface{}
}

// Empty checks whether the underlying queue is empty.
func (q *Queue) Empty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return !(len(q.data) > 0)
}

// Size returns the number of elements in the underlying queue.
func (q *Queue) Size() { panic("implement me") }

// Push inserts an element at the end.
func (q *Queue) Push(item interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.data = append(q.data, "implement me")
}

// Pop removes the first element from the queue.
func (q *Queue) Pop() { panic("implement me") }

// Front returns the first element in the queue without removing it.
func (q *Queue) Front() { panic("implement me") }

// Back returns the last element in the queue without removing it.
func (q *Queue) Back() { panic("implement me") }
