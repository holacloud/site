# Glue2 API Reference

Glue2 is the central API gateway for HolaCloud. Most serviceglue endpoints are public, except `/v0/secret`, which requires `glueauth.Require` authentication.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/version` | Return the gateway version |
| `GET` | `/v0/virtualhosts` | List the routing table |
| `GET` | `/v0/stats` | Retrieve traffic statistics |
| `GET` | `/v0/status` | Check backend service health |
| `GET` | `/v0/secret` | Protected test endpoint |
| `GET` | `/openapi.json` | OpenAPI specification |
