package queue

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mrinalxdev/job-queue/job"
)

func TestEnqueueDequeue(t *testing.T){
	q := NewQueue()
	j := job.Job{
		ID : uuid.NewString(),
		Type : "email",
		Payload: map[string]string{"to":"test@example.com"},
		Priority: job.High,
	}

	q.Enqueue(j)
	result := q.Dequeue()
	if result == nil {
		t.Fatal("EXPECTED A JOB, GOT NIL")
	}

	if result.ID != j.ID {
		t.Errorf("Expected job ID %s, got %s", j.ID, result.ID)
	}
}

func TestDeadLetterQueue(t *testing.T) {
	q := NewQueue()

	j := job.Job{
		ID:       uuid.NewString(),
		Type:     "email",
		Payload:  map[string]string{"to": "fail@example.com"},
		Priority: job.High,
	}

	q.PushToDeadLetter(j)
	dlq := q.GetAllDeadLetters()

	if len(dlq) != 1 {
		t.Errorf("Expected 1 job in dead-letter queue, got %d", len(dlq))
	}
	if dlq[0].ID != j.ID {
		t.Errorf("Expected job ID %s in DLQ, got %s", j.ID, dlq[0].ID)
	}
}
 
func TestDequeueEmptyQueue(t *testing.T){
	q := NewQueue()
	job := q.Dequeue()

	if job != nil {
		t.Error("Expected nil from empty queue, got job")
	}
}
