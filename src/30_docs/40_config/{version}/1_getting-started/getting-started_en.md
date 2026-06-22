# Getting Started

This guide shows the current Config API shape. The v1 user endpoint requires Glue authentication.

## Read User Config

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "X-Glue-Authentication: YOUR_GLUE_TOKEN"
```

Response:

```json
{
  "entries": {
    "feature.newCheckout": true,
    "database.host": "db.example.com"
  }
}
```

## Update User Config

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "X-Glue-Authentication: YOUR_GLUE_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "entries": {
      "feature.newCheckout": false
    }
  }'
```

Response:

```json
{
  "entries": {
    "feature.newCheckout": false
  }
}
```

The v0 API manages config Things directly. Each Thing has an `id` and `entries`.
