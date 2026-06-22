# Update User Config

Update the authenticated user's config entries map.

## Authentication

Requires `X-Glue-Authentication`.

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "X-Glue-Authentication: YOUR_GLUE_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":false}}'
```

```json
{
  "entries": {
    "feature.newCheckout": false
  }
}
```
