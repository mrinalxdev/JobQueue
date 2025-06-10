package worker

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/mrinalxdev/job-queue/job"
	"github.com/mrinalxdev/job-queue/queue"
)

type Worker struct {
	ID    int
	Queue *queue.Queue
}

func (w *Worker) Start() {
	go func(){
		for {
			j := w.Queue.Dequeue()
			if j != nil {
				fmt.Printf("Worker %d processing job ID : %s \n", w.ID, j.ID)

				err := handleJob(j)
				if err != nil {
					log.Printf("Job %s failed : %v\n", j.ID, err)

					j.RetryCount++
					if j.RetryCount <= j.MaxRetries {
						delay := time.Duration(math.Pow(2, float64(j.RetryCount))) * time.Second
						log.Printf("Retrying job %s in %v \n", j.ID, delay)

						go func(jobCopy job.Job){
							time.Sleep(delay)
							w.Queue.Enqueue(jobCopy)
						}(*j)
					} else {
						log.Printf("Job %s moved dto dead-letter queue \n", j.ID)
						w.Queue.PushToDeadLetter(*j)
					}
				}
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func handleJob(j *job.Job) error {
	time.Sleep(500 * time.Millisecond)

	if j.Payload["to"] == "user5@examples.com"{
		return fmt.Errorf("simulated error")
	}

	fmt.Printf("Handled job : %s for %s\n", j.ID, j.Payload["to"])
	return nil
}
