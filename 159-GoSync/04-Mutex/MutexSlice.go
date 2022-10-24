package main

import "sync"

type SliceQueue struct {
	dada []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{dada: make([]interface{}, 0, n)}
}

// Enqueue 把值放在队尾
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.dada = append(q.dada, v)
	q.mu.Unlock()
}

// Dequeue 移去队头并返回
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.dada) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.dada[0]
	q.dada = q.dada[1:]
	q.mu.Unlock()
	return v
}

func main() {

}
