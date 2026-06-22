# 修补 Config

修补 v0 配置 Thing 的 entries。

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
