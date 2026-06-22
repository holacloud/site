# Traffic Stats

Returns real-time traffic statistics for the gateway.

## Description

This endpoint provides aggregated metrics about requests passing through the Glue2 gateway, including request rates, latency percentiles, and status code distributions.

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
{
  "total_requests": 154203,
  "requests_per_second": 42.5,
  "latency_p50_ms": 12,
  "latency_p95_ms": 45,
  "latency_p99_ms": 120,
  "status_codes": {
    "2xx": 148000,
    "4xx": 5200,
    "5xx": 1003
  },
  "active_connections": 87,
  "uptime_seconds": 86400
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Statistics returned successfully |
