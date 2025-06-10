package worker

import (
	"fmt"
	"log"
	"time"

	"github.com/mrinalxdev/job-queue/job"
	"github.com/mrinalxdev/job-queue/queue"
)

type Worker struct {
	ID    int
	Queue *queue.Queue
}

func (w *Worker) Start() {
	go func() {
		for {
			j := w.Queue.Dequeue()
			if j != nil {
				fmt.Printf("Worker %d processing job ID: %s\n", w.ID, j.ID)
				err := handleJob(j)
				if err != nil {
					log.Printf("Job %s failed: %v\n", j.ID, err)
					// Add retry logic here later
				}
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func handleJob(j *job.Job) error {
	time.Sleep(500 * time.Millisecond)

	if j.Type == "email" {
		fmt.Printf("Sending email to %s\n", j.Payload["to"])
		return nil
	}
	return fmt.Errorf("unknown job type: %s", j.Type)
}
