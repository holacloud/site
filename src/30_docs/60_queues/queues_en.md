# Tailon Queues

Tailon is a lightweight, in-memory message queue service that is part of the HolaCloud ecosystem. Designed for speed and simplicity, Tailon enables asynchronous communication between microservices, event distribution, and workload buffering without the overhead of a full-blown message broker.

## Key Features

### In-Memory Performance
Messages are stored in memory for ultra-fast write and read operations. Tailon is optimized for high-throughput, low-latency scenarios where persistence is not the primary concern.

### Long-Poll Consumption
Consumers retrieve messages via long-poll HTTP requests. This reduces the need for constant polling while keeping connections lightweight. The optional `Limit` header controls how many messages are returned per request.

### Streaming Writes
Producers can stream multiple messages in a single request using newline-delimited JSON (NDJSON), making batch ingestion efficient and straightforward.

### Multi-Producer / Multi-Consumer
Any number of producers can write to the same queue, and any number of consumers can read from it. Tailon handles concurrent access safely out of the box.

### Active Client Monitoring
Tailon exposes the `/v1/clients` endpoint so you can inspect which streaming consumers are currently connected to your queues.

## Use Cases

### Async Task Queuing
Offload heavy or slow operations — such as email sending, image processing, or report generation — to background workers that consume from a Tailon queue.

### Event Distribution
Broadcast events from a single producer to multiple consumer services. Each consumer reads independently, enabling fan-out patterns.

### Load Leveling
Buffer incoming requests during traffic spikes. Workers consume at their own pace, preventing downstream services from being overwhelmed.

### Decoupling Microservices
Replace direct HTTP calls between services with queue-based communication. Producers and consumers evolve independently as long as the message contract is stable.

Tailon by HolaCloud is a simple yet powerful queuing solution for modern applications that need fast, reliable, and scalable message passing.
