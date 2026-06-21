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
| `X-Glue-Authentication` | Base64-encoded JWT containing user ID, session info, roles, and auth method |
| `X-Holacloud-Project-Id` | The UUID of the project derived from the virtual host |
| `X-Holacloud-Tenant-Project-Id` | Tenant-scoped project ID for multi-tenant isolation |
| `X-Forwarded-For` | Original client IP address |
| `X-Forwarded-Proto` | Original protocol (`http` or `https`) |

Backend services use these headers to identify the caller and enforce authorization. The `X-Glue-Authentication` JWT is signed with a shared secret known only to Glue2 and trusted backends.

## V0 Admin Endpoints

Glue2 exposes a set of admin endpoints under `/v0/`:

### List Virtual Hosts

```bash
curl "https://api.hola.cloud/v0/virtualhosts" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

```json
{
  "virtual_hosts": [
    {
      "host": "my-project.hola.cloud",
      "project_id": "p3b2c1a0-1234-5678-9abc-def012345678",
      "backend": "svc-1"
    }
  ]
}
```

### Traffic Stats

```bash
curl "https://api.hola.cloud/v0/stats" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

### Backend Health Status

```bash
curl "https://api.hola.cloud/v0/status" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

```json
{
  "backends": [
    {
      "name": "inceptiondb",
      "status": "healthy",
      "latency_ms": 12
    }
  ]
}
```

## Load Balancing & Failover

Glue2 distributes traffic across multiple backend instances for each service. If an instance is unhealthy (fails health checks or returns errors), Glue2 removes it from the rotation and retries on a healthy instance. Health checks run every 10 seconds. You can monitor backend health via the `/v0/status` endpoint.

## Next Steps

- Learn about authentication mechanisms in [Authentication](../2_authentication/authentication_en.md).
- Learn how to manage API keys in [Managing API Keys](../3_api-keys/api-keys_en.md).
