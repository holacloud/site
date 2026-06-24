# OpenAPI Specification

Returns the OpenAPI specification for the KVNode API.

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
  "openapi": "3.0.0",
  "info": {
    "title": "KVNode API",
    "version": "1.0.0"
  },
  "paths": {
    "/healthz": {
      "get": {
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "Node is healthy"
          }
        }
      }
    }
  }
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 500 | internal_error | Internal server error |
