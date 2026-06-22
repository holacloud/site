# Get Version

Returns the current version of the Glue2 gateway.

## Description

The version endpoint returns the configured Glue2 version as plain text.

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
2.3.1
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Version information returned successfully |
