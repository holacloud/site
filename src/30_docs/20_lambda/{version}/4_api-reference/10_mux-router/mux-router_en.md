# Mux Router

Routes public requests through an owner-scoped path. No authentication is required.

The mux route is `/mux/{owner_id}/*`. The remaining path is forwarded to the owner's lambda routing logic.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `owner_id` | string | Owner identifier |
| `*` | path | Remaining path forwarded to the owner scope |

## HTTP Request

```http
GET /mux/user_123/hello-world HTTP/1.1
Host: api.hola.cloud
```

## Example

```bash
curl -X GET "https://api.hola.cloud/mux/user_123/hello-world"
```

## Response

The response is produced by the lambda selected by the owner route.

```json
{
  "body": {
    "message": "Hello from mux",
    "path": "/hello-world"
  }
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 404 | Owner or lambda route not found |
| 500 | Lambda execution error |
