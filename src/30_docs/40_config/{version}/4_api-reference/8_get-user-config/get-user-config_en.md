# Get User Config

Retrieve the authenticated user's config entries map.

## Authentication

Requires `X-Glue-Authentication`.

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "X-Glue-Authentication: YOUR_GLUE_TOKEN"
```

```json
{
  "entries": {
    "feature.newCheckout": true
  }
}
```
