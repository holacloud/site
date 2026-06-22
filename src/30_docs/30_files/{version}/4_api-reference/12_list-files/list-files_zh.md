# 列出文件

列出 bucket 中的文件。`/list/` 后面的路径用作前缀过滤。

## 认证

需要 `X-Glue-Authentication`。

## 请求

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/list/images/" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## 响应

响应是包含已实现字段的 file 对象 JSON array。
