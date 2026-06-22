
# Mux Router

Routes incoming requests to the specified owner's lambda functions based on the sub-path. This is a public endpoint, no authentication required.

The Mux Router allows mapping custom domains or paths to specific lambda owners. The remaining path after `/mux/{owner}/` is forwarded to the owner's lambda routing logic.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| owner | string | The owner identifier (username or project name) |
| `*` | path | The remainder of the path forwarded to the owner's lambdas |

## HTTP Request

```http
GET /mux/acme-webapp/hello-world HTTP/1.1
Host: api.hola.cloud
```

## Example

```bash
curl -X GET "https://api.hola.cloud/mux/acme-webapp/hello-world"
```

## Response

The response depends entirely on the lambda function that handles the routed request.

```json
{
  "status": 200,
  "body": {
    "message": "Hello from acme-webapp's lambda!"
  }
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 404 | Owner or lambda route not found |
| 500 | Lambda execution error |
