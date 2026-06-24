# Service Metadata

These endpoints expose service metadata and diagnostics.

## Route Tree

```http
GET /doc HTTP/1.1
Host: api.hola.cloud
```

Returns a text tree of the routes registered by the service.

## API Root

```http
GET /v1 HTTP/1.1
Host: api.hola.cloud
```

Returns the v1 welcome response.

## Release

```http
GET /release HTTP/1.1
Host: api.hola.cloud
```

Returns release metadata for the deployed service.

## OpenAPI

```http
GET /openapi.json HTTP/1.1
Host: api.hola.cloud
```

Returns the OpenAPI document served by the API.

## Debug Databases

```http
GET /v1/debug/dbs HTTP/1.1
Host: api.hola.cloud
```

Returns a debug database list. This endpoint is intended for diagnostics, not application traffic.
