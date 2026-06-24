# Secret

Protected test endpoint that validates authentication.

## Description

This endpoint requires a valid authentication token. It is used to verify that the authentication layer is working correctly.

## Authentication

Requires `glueauth.Require` authentication via a valid API key or session.

## Request

No request body required.

## Example

```bash
curl -X GET "https://api.hola.cloud/v0/secret" \
  -H "Authorization: Bearer <api-key>"
```

## Response

```text
OK
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Authentication successful |
| 401 | Missing or invalid authentication |
