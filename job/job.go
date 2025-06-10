package job

import "time"

type Job struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Payload    map[string]string `json:"payload"`
	Priority   Priority          `json:"priority"`
	RetryCount int               `json:"retry_count"`
	MaxRetries int               `json:"max_retries"`
	CreatedAt  time.Time         `json:"created_at"`
}
