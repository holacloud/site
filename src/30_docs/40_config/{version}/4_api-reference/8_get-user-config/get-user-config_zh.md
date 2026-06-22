# 获取用户配置

检索已认证用户配置的 entries 映射。

## 认证

需要 `X-Glue-Authentication`。

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
