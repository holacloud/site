# Scheduler

The Scheduler is an in-memory delayed task queue that is part of the HolaCloud ecosystem. It provides time-based task enqueuing, lease-based reservation for worker safety, and snapshot monitoring via Server-Sent Events (SSE). It is designed for short-lived delayed work, not recurring scheduling, automatic redelivery, or shared persistence.

## Key Features

### In-Memory Task Storage
Tasks are held in memory for fast scheduling and dispatch. The service is optimized for high-throughput, low-latency workloads where tasks are processed shortly after their scheduled time.

### Time-Based Scheduling
Each task carries a `future` (or `delay`) field that controls when it becomes available. Tasks are dormant until their scheduled time elapses, at which point they become visible to workers.

### Lease-Based Reservation
When a worker reserves a task, the task enters a lease period defined by `worktime`. If the worker completes and acknowledges the task within the lease, it is removed. If the lease expires, the task becomes available for reservation by another worker. This guarantees at-least-once processing without manual intervention.

### SSE Streaming
The `/schedulers/{id}/tasks/stream` endpoint delivers SSE `snapshot` events. Workers and monitoring dashboards can subscribe to receive the current scheduled and inflight task lists.

### Labeling and Filtering
Tasks can be annotated with `labels` as an array of strings at enqueue time. List operations support filtering by label, making it easy to inspect tasks by project, priority, or worker group.

## Use Cases

### Delayed Dispatch
Schedule one-off delayed work. Each enqueue call creates a task that becomes available at the specified time.

### Delayed Job Processing
Enqueue jobs with a delay, for example sending a reminder email 24 hours after a user signs up. The Scheduler holds the task in memory until its time comes.

### Worker Coordination
Multiple workers can reserve and process tasks from the same scheduler. The lease mechanism prevents concurrent processing while the lease is active; if the lease expires, the task becomes available again.

### Manual Rescheduling
When work must run later, reschedule the task with a new future time or delay string.

The Scheduler by HolaCloud is a lightweight, developer-friendly delayed queue for cloud-native applications.
