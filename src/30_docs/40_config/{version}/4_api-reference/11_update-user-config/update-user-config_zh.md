# 更新用户配置

更新已认证用户配置的 entries 映射。

## 认证

需要 `X-Glue-Authentication`。

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
