# GET /fakeapi/*

Wildcard endpoint that serves mock data for all Console backend services during development and demo mode.

## Description

The `/fakeapi/*` route matches any path under `/fakeapi/` and returns simulated responses. This allows the Console UI to function without a real backend. The path structure mirrors the real API layout so the UI code remains identical between development and production.

## Sub-routes

| Sub-path | Description |
|----------|-------------|
| `/fakeapi/auth/me` | Returns the current mock user session |
| `/fakeapi/inceptionapi` | Simulates InceptionDB API responses |
| `/fakeapi/lambdasapi` | Simulates Lambda function management |
| `/fakeapi/projectsapi` | Simulates project settings and listing |
| `/fakeapi/filesapi` | Simulates file storage operations |

## Authentication

None. These mock endpoints are unauthenticated.

## Example

```bash
curl -X GET "https://api.hola.cloud/fakeapi/auth/me"
```

```json
{
  "user": "demo@holacloud.com",
  "name": "Demo User",
  "project": "default",
  "role": "admin",
  "authenticated": true
}
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/projectsapi"
```

```json
{
  "projects": [
    {
      "id": "default",
      "name": "Default Project",
      "plan": "free",
      "region": "us-east"
    }
  ]
}
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/lambdasapi"
```

```json
{
  "functions": [
    {
      "id": "fn-001",
      "name": "hello-world",
      "runtime": "nodejs18",
      "status": "active"
    }
  ]
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Mock data returned successfully |
| 404 | Path under `/fakeapi/` not recognized |
