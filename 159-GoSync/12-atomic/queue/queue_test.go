package queue

import "testing"

func TestLKQueue_Dequeue(t *testing.T) {
	q := NewLKQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(4)
	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
}
