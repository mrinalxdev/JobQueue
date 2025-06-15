package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mrinalxdev/job-queue/job"
	"github.com/mrinalxdev/job-queue/queue"
	"github.com/mrinalxdev/job-queue/worker"
)

func main() {
	q := queue.NewQueue()

	for i := 1; i <= 2; i++ {
		w := worker.Worker{ID: i, Queue: q}
		w.Start()
	}

	for i := 0; i < 10; i++ {
		j := job.Job{
			ID:         uuid.New().String(),
			Type:       "email",
			Payload:    map[string]string{"to": fmt.Sprintf("user%d@example.com", i)},
			Priority:   job.High,
			MaxRetries: 3,
			CreatedAt:  time.Now(),
		}
		q.Enqueue(j)
	}

	time.AfterFunc(20*time.Second, func() {
		fmt.Println("\n--- Dead Letter Queue ---")
		for _, deadJob := range q.GetAllDeadLetters() {
			fmt.Printf("❌ JobID: %s | To: %s | Retries: %d\n", deadJob.ID, deadJob.Payload["to"], deadJob.RetryCount)
		}
	})

	select {}
}
