# Traffic Stats

Returns real-time traffic statistics for the gateway.

## Description

This endpoint returns an array of per-host gateway statistics. Each item uses Glue2's `prettyStats` fields, with aggregate rows for `*` and `-noroute-`.

## Authentication

None. This endpoint is public.

## Request

No request body required.

## Example

```bash
curl -X GET "https://api.hola.cloud/v0/stats"
```

## Response

```json
[
  {
    "host": "my-project.hola.cloud",
    "served_requests": 154203,
    "serving_time": "32.4s",
    "latency_avg": "210µs",
    "uptime": "24h0m0s",
    "start_timestamp": "2026-06-21T10:00:00Z"
  }
]
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Statistics returned successfully |
