package worker

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mrinalxdev/job-queue/job"
	"github.com/mrinalxdev/job-queue/queue"
)

func TestWorkerRetriesAndDLQ(t *testing.T){
	q := queue.NewQueue()

	j := job.Job{
		ID: uuid.New().String(),
		Type: "email",
		Payload: map[string]string{"to":"user5@example.com"},
		Priority: job.High,
		MaxRetries: 2,
	}

	w := Worker{ID : 1, Queue: q}
	w.Start()

	q.Enqueue(j)
	time.Sleep(8 * time.Second)

	dlq := q.GetAllDeadLetters()
	if len(dlq) != 1 {
		t.Fatalf("Expected 1 job in DLQ, got %d", len(dlq))
	}
	if dlq[0].ID != j.ID {
		t.Errorf("Expected job %s in DLQ, got %s", j.ID, dlq[0].ID)
	}
}