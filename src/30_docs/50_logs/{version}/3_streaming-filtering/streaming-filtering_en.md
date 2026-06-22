# Streaming Filtering

InstantLogs provides long-lived HTTP streaming endpoints that deliver log events in near real-time. You can filter streams by regex pattern and choose between NDJSON and plain-text output formats.

## Streaming Endpoints

### Events Stream

The `GET /v1/loggers/{id}/events` endpoint streams all log entries as they are ingested. By default, events are returned as NDJSON (Newline-Delimited JSON).

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

Each line of the response body is a JSON object representing one log entry:

```
{"id":"evt_001","message":"Request started","level":"info","service":"api-gateway","timestamp":"2025-06-20T15:00:00Z"}
{"id":"evt_002","message":"Request completed","level":"info","service":"api-gateway","timestamp":"2025-06-20T15:00:05Z"}
```

To receive plain-text output, set the `Accept` header to `text/plain`:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Accept: text/plain"
```

### Filtered Stream

The `GET /v1/loggers/{id}/filter` endpoint applies a regex filter before streaming. Only log entries whose message matches the pattern are delivered.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=error|critical" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

This streams only entries containing `error` or `critical` in the message field. The pattern parameter accepts any valid Go regular expression.

### Combining Filter and Format

You can combine regex filtering with your preferred output format:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=^\\[WARN\\]" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Accept: text/plain"
```

## Logger Stats

The `GET /v1/loggers/{id}/stats` endpoint returns statistics about a logger without consuming the event stream.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/stats" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

Expected response:

```json
{
  "logger_id": "logger_xyz789",
  "total_events": 15234,
  "events_last_hour": 342,
  "storage_bytes": 1048576,
  "created_at": "2025-06-20T14:22:00Z"
}
```

## HTTP Request Reference

```http
GET /v1/loggers/{id}/events HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456

GET /v1/loggers/{id}/filter?pattern=error HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456

GET /v1/loggers/{id}/stats HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```
