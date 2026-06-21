# Getting Started with InstantLogs

This guide will help you create a logger, ingest a log entry, and stream events in real time.

## Step 1: Create a Logger

First, create a logger using the management API. You need an API key with management permissions.

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-app-logger"
  }'
```

Expected response:

```json
{
  "id": "logger_xyz789",
  "name": "my-app-logger",
  "secret": "lgs_abc123def456",
  "created_at": "2025-06-20T14:22:00Z"
}
```

Keep the `secret` value safe — it is used to authenticate data operations.

## Step 2: Ingest a Log Entry

Use the Logger Secret to send a log entry to your logger:

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "Hello from my app!",
    "level": "info",
    "service": "web",
    "timestamp": "2025-06-20T14:22:30Z"
  }'
```

Expected response:

```json
{
  "ingested": 1
}
```

## Step 3: Stream Logs in Real Time

Open a streaming connection to see logs as they arrive:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

The server will keep the connection open and push new log entries as NDJSON:

```
{"message":"Hello from my app!","level":"info","service":"web","timestamp":"2025-06-20T14:22:30Z","id":"evt_001"}
```

With the `Accept: text/plain` header, you can receive plain-text lines instead.

## Step 4: Filter the Stream

Append a regex pattern to filter events:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=error" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

Only log entries whose message matches the pattern `error` will be streamed.

## Summary

You created a logger, ingested a log, streamed events live, and applied a regex filter. InstantLogs makes it easy to centralize and observe your application logs in real time.
