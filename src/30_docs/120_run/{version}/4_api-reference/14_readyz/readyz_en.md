# Readiness Check

Returns the service readiness status.

## Request

```http
GET /readyz
```

## Example

```bash
curl "https://api.hola.cloud/readyz"
```

## Response

Returns `ok` with a 200 status code when the service is ready to accept traffic.

```text
ok
```
