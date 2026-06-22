# Managing Configs

The v0 API manages config Things directly. A config Thing has an `id` and an `entries` object.

## List

```bash
curl "https://api.hola.cloud/v0/configs"
```

```json
[
  {
    "id": "cfg_abc123",
    "entries": {
      "feature.newCheckout": true
    }
  }
]
```

## Create

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":true}}'
```

## Retrieve

```bash
curl "https://api.hola.cloud/v0/configs/cfg_abc123"
```

## Patch

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":false}}'
```

## Delete

```bash
curl -X DELETE "https://api.hola.cloud/v0/configs/cfg_abc123"
```
