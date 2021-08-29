package datastructures

import "fmt"

func NewQueue() queue {
	return queue{}
}

func (q *queue) Enqueue(entry int) {
	q.entries = append(q.entries, entry)
}

func (q *queue) Dequeue() (int, bool) {

	if q.Empty() {
		fmt.Println("\nEmpty Queue")
		return 0, false
	}

	toDequeue := q.entries[0]
	q.entries = q.entries[1:]
	return toDequeue, true
}

func (q *queue) GetFront() (int, bool) {
	if q.Empty() {
		fmt.Println("\nEmpty Queue")
		return 0, false
	}
	return q.entries[0], true
}

func (q *queue) GetRear() (int, bool) {
	if q.Empty() {
		fmt.Println("\nEmpty Queue")
		return 0, false
	}
	return q.entries[len(q.entries)-1], true
}

func (q *queue) Clear() {
	q.entries = []int{}
}

func (q *queue) Empty() bool {
	return len(q.entries) == 0
}

func (q *queue) Size() int {
	return len(q.entries)
}

func (q *queue) PrintAllValues() {
	if q.Empty() {
		fmt.Println("\nEmpty Queue")
		return
	}
	fmt.Print("\nQueue :")
	for j := len(q.entries) - 1; j >= 0; j-- {
		fmt.Printf(" %v", q.entries[j])
	}
	fmt.Println()
}

type queue struct {
	entries []int
}
