# Get Version

Returns the current version of the Console service.

## Description

This endpoint returns the configured Console version as plain text. It requires no authentication.

## Authentication

None. This endpoint is public.

## Request

No request body required.

## Example

```bash
curl -X GET "https://api.hola.cloud/version"
```

## Response

```text
1.0.0
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Version information returned successfully |
