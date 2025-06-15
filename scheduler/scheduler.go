package scheduler

import (
	"container/heap"
	"sync"
	"time"

	"github.com/mrinalxdev/job-queue/job"
	"github.com/mrinalxdev/job-queue/queue"
)

type ScheduledJob struct {
	Job job.Job
	ScheduleAt time.Time
	index int
}

type JobHeap []*ScheduledJob

func (h JobHeap) Len() int { return len(h) }
func (h JobHeap) Less(i, j int) bool {
	return h[i].ScheduleAt.Before(h[j].ScheduleAt)
}

func (h JobHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *JobHeap) Push(x interface {}){
	n := len(*h)
	item := x.(*ScheduledJob)
	item.index = n
	*h = append(*h, item)
}

func (h *JobHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.index = -1
	*h = old[0 : n-1]
	return item
}

type Scheduler struct {
	mu sync.Mutex
	jobs JobHeap
	queue *queue.Queue
}

func NewScheduler(q *queue.Queue) *Scheduler{
	h :=  make(JobHeap, 0)
	heap.Init(&h)

	s := &Scheduler{jobs : h, queue : q}
	go s.poller()
	return s
}

func (s *Scheduler) Scheduler(j job.Job, delay time.Duration){
	s.mu.Lock()
	defer s.mu.Unlock()


	scheduled := &ScheduledJob{
		Job : j,
		ScheduleAt: time.Now().Add(delay),
	}

	heap.Push(&s.jobs, scheduled)
}


func (s *Scheduler) poller() {
	for {
		s.mu.Lock()
		if s.jobs.Len() > 0 {
			next := s.jobs[0]
			if time.Now().After(next.ScheduleAt){
				heap.Pop(&s.jobs)
				s.queue.Enqueue(next.Job)
			}
		}

		s.mu.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}