package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	Queue []string
	cond  *sync.Cond
}

func main() {
	q := Queue{
		Queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		q.Dequeue()
		time.Sleep(time.Second)
	}
}

func (q *Queue) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.Queue = append(q.Queue, item)
	fmt.Printf("putting %s to queue, notify all\n", item)
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.Queue) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	result := q.Queue[0]
	q.Queue = q.Queue[1:]
	return result
}
