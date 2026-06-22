# Documentation

Welcome to the HolaCloud documentation. Here you will find guides and API references for the services in the platform.

## Getting Started

If you are new to HolaCloud, start with the [Console](/docs/console/) and the individual service guides. Some services are still evolving, so the API reference should be read together with the current implementation-backed notes.

Each service below has its own getting-started guide, conceptual documentation, and API reference with endpoint details, request examples, and error codes.

## Services

### [InceptionDB](/docs/inceptiondb/)
NoSQL document database. Store and query JSON documents with dynamic schema, collections, and indexes. Ideal for applications that need flexible data models.

### [Lambda](/docs/lambda/)
Function execution platform. Deploy JavaScript functions and static assets behind HTTP endpoints.

### [Files](/docs/files/)
File storage API. Upload, download, list, inspect, and delete files organized in buckets.

### [Config](/docs/config/)
Configuration API. Store and update user-scoped JSON configuration entries.

### [InstantLogs](/docs/logs/)
Real-time log ingestion and streaming. Collect raw or framed logs, filter with regular expressions, and stream events.

### [Tailon](/docs/queues/)
Message queue API. Create queues, write JSON messages, and read newline-delimited JSON from named queues.

### [Scheduler](/docs/scheduler/)
Delayed task queue. Enqueue one-off delayed tasks, reserve them with leases, and stream scheduler snapshots.

### [KVNode](/docs/kvnode/)
Key-value API with collections, keys, node status, metrics, and root/replica synchronization.

### [Console](/docs/console/)
Web-based console surface for HolaCloud services, including the current fake API endpoints used by the UI.

### [Glue2](/docs/glue2/)
Host-based gateway and service glue layer with virtual hosts, authentication context injection, status, and traffic stats.

### [Holamail](/docs/holamail/)
Basic SMTP test server. Accepts and logs SMTP messages for testing and development.

### [Run](/docs/run/)
Container runtime controls with a built-in push-oriented Docker registry subset. Start, stop, and roll back repository runtimes.

## API Reference

Each service includes an API reference section with:

- Endpoint descriptions and HTTP methods
- Authentication requirements
- Request and response examples in JSON
- Error code tables
- `curl` examples for every endpoint

Browse the sidebar to navigate to any service or use the product links in the top navigation.
