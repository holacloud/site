
# POST /v1/ingest/events

Ingest framed events using the logframe protocol. This endpoint accepts a binary protocol with metadata and message length prefixes.

## Authentication

Requires data credentials:

- `X-Instantlogs-Event-Secret: <secret>` or `Authorization: Bearer <secret>`

## Request Body

The request body uses the logframe protocol format:

```
len_metadata:len_message:metadata_json:message\n
```

Each frame consists of:
- `len_metadata` — Length of the metadata JSON (as ASCII number)
- `len_message` — Length of the message (as ASCII number)
- `metadata_json` — JSON object with event metadata
- `message` — The log message content
- `\n` — Newline delimiter

## Request

```bash
# Frame: metadata={"service":"web"} + message="User logged in"
# len_metadata=18, len_message=14
# Format: 18:14:{"service":"web"}User logged in\n
printf '18:14:{"service":"web"}User logged in\n' | \
curl -X POST "https://api.hola.cloud/v1/ingest/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @-
```

```http
POST /v1/ingest/events HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
Content-Type: application/octet-stream

18:14:{"service":"web"}User logged in\n
```

## Response

```json
{
  "ingested": 1
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Malformed frame — invalid length prefix or missing delimiter |
| 401 | Missing or invalid event secret |
| 413 | Payload too large |
