# Create Config

Create a v0 config Thing.

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":true}}'
```

Response:

```json
{
  "id": "cfg_abc123",
  "entries": {
    "feature.newCheckout": true
  }
}
```
