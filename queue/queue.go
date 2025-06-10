package queue

import (
	"container/list"
	"sync"

	"github.com/mrinalxdev/job-queue/job"
)

type Queue struct {
	mu     sync.Mutex
	queues map[job.Priority]*list.List
}

func NewQueue() *Queue {
	return &Queue{
		queues: map[job.Priority]*list.List{
			job.High: list.New(),
			job.Low:  list.New(),
		},
	}
}

func (q *Queue) Enqueue(j job.Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queues[j.Priority].PushBack(j)
}

func (q *Queue) Dequeue() *job.Job {
	q.mu.Lock()
	defer q.mu.Unlock()

	for _, p := range []job.Priority{job.High, job.Low} {
		queue := q.queues[p]
		if queue.Len() > 0 {
			front := queue.Front()
			j := front.Value.(job.Job)
			queue.Remove(front)
			return &j
		}
	}
	return nil
}
