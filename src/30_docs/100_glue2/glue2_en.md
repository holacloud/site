# Glue2

Glue2 is the central API gateway for HolaCloud. It authenticates every incoming request, routes it to the correct backend service, and collects access logs. All traffic to HolaCloud services passes through Glue2.

## Authentication

Glue2 supports three authentication mechanisms:

### Sessions
Browser-based sessions authenticated via HolaAuth. After login, the session cookie is used to authenticate subsequent requests.

### API Keys
Machine-to-machine access using an API key and secret pair. These are passed as the `apikey` and `secret` HTTP headers and are validated against the project's credentials.

### OAuth
OAuth 2.0 integration allows authentication through external providers. Glue2 validates tokens and injects user identity information into proxied requests.

## Routing by Virtual Host

Glue2 routes requests based on the `Host` header. Each project in HolaCloud gets a unique subdomain (`project.hola.cloud`), and Glue2 maps incoming requests to the appropriate backend service for that project.

```
┌─────────────┐     ┌──────────┐     ┌──────────────┐
│   Client    │────▶│  Glue2   │────▶│  Project Svc  │
└─────────────┘     └──────────┘     └──────────────┘
                           │
                           ├────▶ InceptionDB
                           ├────▶ Lambda
                           ├────▶ KVNode
                           └────▶ Files / Config / etc.
```

## Injected Headers

When Glue2 forwards a request to a backend service, it injects the following headers:

| Header | Description |
|--------|-------------|
| `X-Glue-Authentication` | Plain JSON authentication context |
| `X-Holacloud-Project-Id` | The project ID extracted from the virtual host |
| `X-Forwarded-For` | The original client IP address |

Backend services use these headers for authorization and tenant isolation.

## Access Logs

All requests passing through Glue2 are logged with:

- Timestamp and request duration
- Client IP and user agent
- Request method, path, and status code
- Project ID and authentication method
- Backend service target and response time

Logs are streamed to InstantLogs for real-time monitoring and analysis.
