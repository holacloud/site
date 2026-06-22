# Config API Reference

Base URL: `https://api.hola.cloud`

## Authentication

`/v1/config` requires `X-Glue-Authentication`. `/v0/configs` manages config Things directly.

## Endpoints

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | `/v0/configs` | List config Things | Public |
| POST | `/v0/configs` | Create a config Thing with `entries` | Public |
| GET | `/v0/configs/{id}` | Get a config Thing by ID | Public |
| DELETE | `/v0/configs/{id}` | Delete a config Thing | Public |
| PATCH | `/v0/configs/{id}` | Patch a config Thing's `entries` | Public |
| GET | `/v1/config` | Get the authenticated user's config `entries` map | Glue auth |
| PATCH | `/v1/config` | Update the authenticated user's config `entries` map | Glue auth |

## Shapes

v0 config Thing:

```json
{
  "id": "cfg_abc123",
  "entries": {
    "feature.newCheckout": true
  }
}
```

v1 user config:

```json
{
  "entries": {
    "feature.newCheckout": true
  }
}
```
