# Delete Key

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
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "ok": true,
  "seq": 2,
  "version": 2
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 403 | forbidden | Missing authentication headers |
| 404 | missing_collection | Collection does not exist |
| 404 | not_found | Key not found |
| 502 | parent_unavailable | Parent node is not reachable |
| 500 | internal_error | Internal server error |
