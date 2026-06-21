# GET /openapi.json

Returns the OpenAPI specification for the Glue2 gateway.

## Description

This endpoint serves the OpenAPI 3.0 specification document that describes all routes, request formats, and response schemas exposed by the gateway. The specification is auto-generated from the route definitions.

## Authentication

None. This endpoint is public.

## Request

No request body required.

## Example

```bash
curl -X GET "https://api.hola.cloud/openapi.json"
```

## Response

```json
{
  "openapi": "3.0.3",
  "info": {
    "title": "Glue2 API Gateway",
    "version": "2.3.1",
    "description": "Central API gateway for HolaCloud services"
  },
  "paths": {
    "/version": {
      "get": {
        "summary": "Get gateway version",
        "responses": {
          "200": {
            "description": "Version information"
          }
        }
      }
    },
    "/v0/virtualhosts": {
      "get": {
        "summary": "List virtual hosts",
        "responses": {
          "200": {
            "description": "Routing table"
          }
        }
      }
    },
    "/v0/stats": {
      "get": {
        "summary": "Get traffic stats",
        "responses": {
          "200": {
            "description": "Traffic statistics"
          }
        }
      }
    },
    "/v0/status": {
      "get": {
        "summary": "Backend health status",
        "responses": {
          "200": {
            "description": "Service status"
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
| 200 | OpenAPI specification returned successfully |
