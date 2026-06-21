# Scheduler

The Scheduler is a distributed task scheduling service that is part of the HolaCloud ecosystem. It provides time-based task enqueuing, lease-based reservation for worker safety, and real-time monitoring via Server-Sent Events (SSE). Designed as a lightweight alternative to traditional cron systems, the Scheduler integrates seamlessly with the rest of your HolaCloud services.

## Key Features

### In-Memory Task Storage
Tasks are held in memory for fast scheduling and dispatch. The service is optimized for high-throughput, low-latency workloads where tasks are processed shortly after their scheduled time.

### Time-Based Scheduling
Each task carries a `future` (or `delay`) field that controls when it becomes available. Tasks are dormant until their scheduled time elapses, at which point they become visible to workers.

### Lease-Based Reservation
When a worker reserves a task, the task enters a lease period defined by `worktime`. If the worker completes and acknowledges the task within the lease, it is removed. If the lease expires, the task becomes available for reservation by another worker. This guarantees at-least-once processing without manual intervention.

### SSE Streaming
The `/schedulers/{id}/tasks/stream` endpoint delivers a real-time SSE stream of task events. Workers and monitoring dashboards can subscribe to receive new tasks, completions, and lease expirations as they happen.

### Labeling and Filtering
Tasks can be annotated with arbitrary `labels` (key-value pairs) at enqueue time. List operations support filtering by labels, making it easy to inspect tasks by project, priority, or worker group.

## Use Cases

### Cron Replacement
Schedule recurring or one-off jobs without managing crontab files. Each scheduled call creates a task that becomes available at the specified time.

### Delayed Job Processing
Enqueue jobs with a delay — for example, sending a reminder email 24 hours after a user signs up. The Scheduler holds the task until its time comes.

### Distributed Workers
Multiple workers can reserve and process tasks from the same scheduler. The lease mechanism ensures that a task is processed exactly once under normal conditions, and automatically retried if a worker fails.

### Retry Queues
When a task fails, reschedule it with an increased delay. The Scheduler's `reschedule` endpoint makes it simple to implement exponential backoff.

The Scheduler by HolaCloud is a lightweight, developer-friendly alternative to traditional job schedulers, built for cloud-native applications.
