# Health Check

Returns the service health status.

## Request

```http
GET /healthz
```

## Example

```bash
curl "https://api.hola.cloud/healthz"
```

## Response

Returns `ok` with a 200 status code when the service is healthy.

```text
ok
```
