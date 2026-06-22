# Documentation

Welcome to the HolaCloud documentation. Here you will find guides, API references, and operational manuals for every service in the platform.

## Getting Started

If you are new to HolaCloud, the best place to start is the [Console](/docs/console/) — the web-based management interface where you can provision resources, manage projects, and monitor your services.

Each service below has its own getting-started guide, conceptual documentation, and a complete API reference with endpoint details, request examples, and error codes.

## Services

### [InceptionDB](/docs/inceptiondb/)
Managed NoSQL document database. Store and query JSON documents with dynamic schema, indexing, and JavaScript execution. Ideal for applications that need flexible data models.

### [Lambda](/docs/lambda/)
Serverless function execution platform. Deploy code that runs on demand, scales automatically, and integrates with the rest of the HolaCloud ecosystem. No servers to manage.

### [Files](/docs/files/)
Secure object storage service. Upload, download, and organize files with configurable bucket policies. Suitable for media storage, backups, and content distribution.

### [Config](/docs/config/)
Centralized configuration management. Store, version, and distribute application configuration across environments. Supports JSON patches and user-scoped overrides.

### [InstantLogs](/docs/logs/)
Real-time log ingestion and streaming. Collect structured and unstructured log data from any source, filter in real time, and integrate with monitoring pipelines.

### [Tailon](/docs/queues/)
Message queue and async workload system. Reliable message delivery with long-poll consumers, client tracking, and support for high-throughput asynchronous processing.

### [Scheduler](/docs/scheduler/)
Distributed task scheduling service. Schedule one-off and recurring tasks, manage task leases, and stream task events. Built for time-sensitive and background workloads.

### [KVNode](/docs/kvnode/)
Replicated key-value store with configurable persistence backends. Low-latency reads and writes, health checking, and node-level metrics.

### [Console](/docs/console/)
Web-based management console for all HolaCloud services. Provision projects, manage resources, and monitor service health from a single interface.

### [Glue2](/docs/glue2/)
API Gateway and authentication layer. Central routing, authentication, and traffic management for the HolaCloud service mesh. Supports API key and bearer token auth.

### [Holamail](/docs/holamail/)
SMTP transactional email service. Send emails via SMTP relay with reliable delivery tracking. Suitable for notifications, password resets, and transactional messaging.

### [Run](/docs/run/)
Container execution service with a built-in Docker registry. Start, stop, and roll back containers. Manage environment variables, volumes, and deployment lifecycles.

## API Reference

Each service includes a complete API reference section with:

- Endpoint descriptions and HTTP methods
- Authentication requirements
- Request and response examples in JSON
- Error code tables
- `curl` examples for every endpoint

Browse the sidebar to navigate to any service or use the product links in the top navigation.
