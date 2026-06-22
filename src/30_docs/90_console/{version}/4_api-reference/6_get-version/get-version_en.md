# Get Version

Returns the current version of the Console service.

## Description

This endpoint is used by the Console UI to display the version information in the application footer and diagnostics pages. It requires no authentication.

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
  "version": "1.0.0",
  "build": "20260601-abcdef",
  "mode": "development"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Version information returned successfully |
