# DELETE /v1/collections/{col}/keys/*

Deletes a key from a collection. The wildcard `*` in the path is replaced with the key name.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| col | string | The name of the collection |

## Example Request

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/users/keys/user:1001" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## Example Response

```http
HTTP/1.1 204 No Content
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | unauthorized | Missing or invalid authentication |
| 404 | not_found | Collection or key not found |
| 500 | internal_error | Internal server error |
