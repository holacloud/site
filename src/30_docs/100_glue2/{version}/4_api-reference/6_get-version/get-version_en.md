# GET /version

Returns the current version of the Glue2 gateway.

## Description

The version endpoint provides build and release information about the running gateway instance.

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
  "service": "glue2",
  "version": "2.3.1",
  "commit": "a1b2c3d4",
  "build_time": "2026-06-20T12:00:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Version information returned successfully |
