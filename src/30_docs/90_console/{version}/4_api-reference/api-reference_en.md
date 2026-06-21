# Console API Reference

The Console exposes a set of mock API endpoints used by the web UI during development and demo mode. These endpoints simulate backend services and require no authentication.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/version` | Return the current Console version |
| `GET` | `/fakeapi/*` | Mock API endpoints for development |

### GET /version

Returns the version string for the Console service.

### GET /fakeapi/*

Wildcard endpoint that serves mock data for various backend services including authentication, InceptionDB, Lambda, Files, and project management.
