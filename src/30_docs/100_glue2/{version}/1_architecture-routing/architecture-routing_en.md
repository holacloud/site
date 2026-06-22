<h1>Architecture & Routing</h1>

Glue2 is the central reverse proxy that sits in front of every HolaCloud service. All incoming traffic — whether from browser users, API clients, or internal services — passes through Glue2 before reaching its destination.

![Glue2 Architecture](img/glue2-architecture.png)

## Reverse Proxy

Glue2 operates as a Layer-7 reverse proxy. It terminates TLS, inspects each request, authenticates the caller, and then forwards the request to the appropriate backend service. No external traffic reaches HolaCloud services directly; Glue2 is the single entry point.

```
Client ──▶ Glue2 (auth check) ──▶ Backend Service
              │
              ├──▶ InceptionDB
              ├──▶ Lambda
              ├──▶ Files
              ├──▶ Config
              ├──▶ KVNode
              ├──▶ Scheduler
              └──▶ Logs
```

## Virtual Host Routing

Glue2 routes requests based on the `Host` header. Each HolaCloud project is assigned a unique subdomain (`<project>.hola.cloud`), and Glue2 maps that subdomain to the project's backend service.

Built-in platform services use reserved subdomains:

| Host | Service |
|------|---------|
| `inceptiondb.hola.cloud` | InceptionDB |
| `api.hola.cloud` | Public API (Lambda, Files, Config, etc.) |
| `console.hola.cloud` | HolaCloud Console |
| `auth.hola.cloud` | Authentication / Session service |
| `logs.hola.cloud` | InstantLogs |

When a request arrives for `my-project.hola.cloud`, Glue2 looks up the project with that slug, resolves its primary backend service, and forwards the request.

## The Proxy Flow

1. Client sends an HTTP request to `<project>.hola.cloud`.
2. Glue2 terminates the TLS connection and reads the `Host` header.
3. Glue2 checks authentication (session cookie, API key, or OAuth token). If the endpoint requires auth and none is provided, the request is rejected with `401 Unauthorized`.
4. Glue2 resolves the virtual host to a project ID and backend service target.
5. Glue2 injects authentication and project headers into the request.
6. The request is forwarded to the backend service.
7. The backend response is streamed back to the client, and Glue2 logs the transaction.

## Injected Headers

When forwarding to a backend, Glue2 injects:

| Header | Description |
|--------|-------------|
| `X-Glue-Authentication` | Plain JSON authentication context |
| `X-Holacloud-Project-Id` | The UUID of the project derived from the virtual host |
| `X-Holacloud-Tenant-Project-Id` | Tenant-scoped project ID for multi-tenant isolation |
| `X-Forwarded-For` | Original client IP address |
| `X-Forwarded-Proto` | Original protocol (`http` or `https`) |

Backend services use these headers to identify the caller and enforce authorization. `X-Glue-Authentication` is plain JSON.

## V0 Admin Endpoints

Glue2 exposes a set of admin endpoints under `/v0/`:

### List Virtual Hosts

```bash
curl "https://api.hola.cloud/v0/virtualhosts" \
```

```json
["my-project.hola.cloud", "api.my-project.hola.cloud"]
```

### Traffic Stats

```bash
curl "https://api.hola.cloud/v0/stats" \
```

### Backend Health Status

```bash
curl "https://api.hola.cloud/v0/status" \
```

```json
[
  {
    "id": "project-123",
    "name": "My Project",
    "host": "my-project.hola.cloud",
    "status": 200,
    "statusText": "200 OK"
  }
]
```

## Next Steps

- Learn about authentication mechanisms in [Authentication](../2_authentication/authentication_en.md).
- Learn how to manage API keys in [Managing API Keys](../3_api-keys/api-keys_en.md).
