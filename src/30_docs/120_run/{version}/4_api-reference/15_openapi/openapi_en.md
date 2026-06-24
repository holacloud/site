# OpenAPI Specification

Returns the OpenAPI specification for the Run service.

## Request

```http
GET /openapi.json
```

## Example

```bash
curl "https://api.hola.cloud/openapi.json"
```

## Response

Returns the OpenAPI specification as JSON.

```json
{
  "openapi": "3.0.0",
  "info": {
    "title": "HolaCloud Run API",
    "version": "1.0.0"
  },
  "paths": {}
}
```
