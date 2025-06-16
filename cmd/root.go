package cmd

import (
	"fmt"
	"os"
	// "strconv"
	"time"

	"github.com/google/uuid"
	"github.com/mrinalxdev/job-queue/job"
	"github.com/mrinalxdev/job-queue/queue"
	"github.com/mrinalxdev/job-queue/scheduler"
	"github.com/mrinalxdev/job-queue/worker"
	"github.com/urfave/cli/v2"
)

var q = queue.NewQueue()
var s = scheduler.NewScheduler(q)

func StartCLI() {
	app := &cli.App{
		Name:  "Job Queue CLI",
		Usage: "Manage jobs, workers, and queues",
		Commands: []*cli.Command{
			{
				Name:  "enqueue",
				Usage: "Enqueue a job",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "to", Required: true},
					&cli.StringFlag{Name: "priority", Value: "low"},
					&cli.IntFlag{Name: "retries", Value: 3},
					&cli.IntFlag{Name: "delay", Value: 0, Usage: "Delay in seconds"},
				},
				Action: func(c *cli.Context) error {
					priority := job.Low
					if c.String("priority") == "high" {
						priority = job.High
					}

					j := job.Job{
						ID:         uuid.New().String(),
						Type:       "email",
						Payload:    map[string]string{"to": c.String("to")},
						Priority:   priority,
						MaxRetries: c.Int("retries"),
						CreatedAt:  time.Now(),
					}

					delay := c.Int("delay")
					if delay > 0 {
						s.Scheduler(j, time.Duration(delay)*time.Second)
						fmt.Println("Scheduled job:", j.ID)
					} else {
						q.Enqueue(j)
						fmt.Println("Enqueued job:", j.ID)
					}
					return nil
				},
			},
			{
				Name:  "start",
				Usage: "Start worker(s)",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "count", Value: 1},
				},
				Action: func(c *cli.Context) error {
					count := c.Int("count")
					for i := 1; i <= count; i++ {
						w := worker.Worker{ID: i, Queue: q}
						go w.Start()
					}
					fmt.Printf("Started %d worker(s)\n", count)
					select {} // block forever
				},
			},
			{
				Name:  "dlq",
				Usage: "Show dead-letter queue",
				Action: func(c *cli.Context) error {
					dlq := q.GetAllDeadLetters()
					if len(dlq) == 0 {
						fmt.Println("No failed jobs")
						return nil
					}
					fmt.Println("Dead-letter jobs:")
					for _, j := range dlq {
						fmt.Printf("- %s (%s)\n", j.ID, j.Payload["to"])
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error:", err)
	}
}
