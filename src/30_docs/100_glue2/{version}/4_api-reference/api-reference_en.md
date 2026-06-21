# Glue2 API Reference

Glue2 is the central API gateway for HolaCloud. Most endpoints are public, except those under `/v0/secret` which require glueauth.Require authentication.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/version` | Return the gateway version |
| `GET` | `/v0/virtualhosts` | List the routing table |
| `GET` | `/v0/stats` | Retrieve traffic statistics |
| `GET` | `/v0/status` | Check backend service health |
| `GET` | `/openapi.json` | OpenAPI specification |
