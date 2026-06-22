# Patch Config

Patch a v0 config Thing's entries.

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":false}}'
```

```json
{
  "id": "cfg_abc123",
  "entries": {
    "feature.newCheckout": false
  }
}
```
