# 列出 Buckets

列出认证用户拥有的 buckets。

## 认证

需要 `X-Glue-Authentication`。

## 请求

```bash
curl "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## 响应

响应是 JSON array，包含 `id`, `name`, `description`, `created_timestamp`, `created_h`, `owners`。
