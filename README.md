# JobQueue

My main task for this project was not to use redis instead replicate few functionalities
A Redis-free background **job queue system** written in Go with support for:

- ✅ Multiple priority queues (High, Low)
- ✅ Background workers
- ✅ Job retries + dead-letter queue
- ✅ Delayed job scheduling
- ✅ Dynamic CLI to enqueue and manage jobs

## Features

| Feature               | Description                                        |
|----------------------|----------------------------------------------------|
| Priority Queues    | High / Low priority job queues                     |
| Worker Pool        | Workers poll queues concurrently                   |
| Delayed Scheduling | Supports scheduled jobs (simulates Redis ZSET)     |
| Retry & DLQ        | Retry with exponential backoff, move to DLQ on fail |
| CLI Support        | Enqueue, schedule, and start workers from CLI      |


