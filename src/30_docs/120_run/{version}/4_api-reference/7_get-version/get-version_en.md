# GET /version

Returns the current version of the Run service.

## Description

Provides version and build information about the running Run service instance.

## Authentication

None. This endpoint is public.

## Request

No request body required.

## Example

```bash
curl -X GET "https://api.hola.cloud/version"
```

## Response

```json
{
  "service": "run",
  "version": "1.5.2",
  "commit": "e5f6g7h8",
  "build_time": "2026-06-20T12:00:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Version information returned successfully |
