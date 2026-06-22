# Delete Collection

Deletes a collection and all its keys.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| col | string | The name of the collection |

## Example Request

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/users" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "ok": true,
  "collection": "users"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | Collection not found |
| 500 | internal_error | Internal server error |
