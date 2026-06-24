# OpenAPI Specification

Returns the OpenAPI specification document for the Lambda API.

## Authentication

Requires `X-Glue-Authentication`.

## HTTP Request

```http
GET /api/v0/openapi HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## Example

```bash
curl -X GET "https://api.hola.cloud/api/v0/openapi" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## Response

Returns the full OpenAPI specification as a JSON document.

```json
{
  "openapi": "3.0.3",
  "info": {
    "title": "Lambda API",
    "version": "1.0.0"
  },
  "paths": {
    "/api/v0/lambdas": {
      "get": {
        "summary": "List lambdas",
        "responses": {
          "200": {
            "description": "List of lambdas"
          }
        }
      }
    }
  }
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication |
