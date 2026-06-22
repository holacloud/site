# 创建 Config

创建 v0 配置 Thing。

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":true}}'
```

响应：

```json
{
  "id": "cfg_abc123",
  "entries": {
    "feature.newCheckout": true
  }
}
```
