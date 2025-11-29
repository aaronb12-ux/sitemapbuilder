package bfs

import (
	"fmt"
)

type Queue struct {
	data []string
}

func (q *Queue) Enqueue(url string) {
	q.data = append(q.data, url)
}

func (q *Queue) Dequeue() string {
	if len(q.data) == 0 {
		fmt.Println("queue is empty")
		return "empty"
	}
	front := q.data[0]
	return front
}

func CreateQueue() *Queue {
	queue := &Queue{}
	return queue
}