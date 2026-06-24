# OpenAPI Specification

Returns the OpenAPI specification for the Scheduler API.

## Authentication

This endpoint is public. No authentication required.

## Example Request

```bash
curl "https://api.hola.cloud/openapi.json"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "openapi": "3.1.0",
  "info": {
    "title": "Scheduler API",
    "version": "1.0.0"
  },
  "paths": {}
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 500 | internal_error | Internal server error |
