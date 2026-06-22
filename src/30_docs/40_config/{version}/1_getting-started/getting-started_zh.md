# 入门

本指南展示当前 Config API 结构。v1 用户端点需要 Glue 认证。

## 读取用户配置

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "X-Glue-Authentication: YOUR_GLUE_TOKEN"
```

响应：

```json
{
  "entries": {
    "feature.newCheckout": true,
    "database.host": "db.example.com"
  }
}
```

## 更新用户配置

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

响应：

```json
{
  "entries": {
    "feature.newCheckout": false
  }
}
```

v0 API 直接管理配置 Thing。每个 Thing 都有 `id` 和 `entries`。
