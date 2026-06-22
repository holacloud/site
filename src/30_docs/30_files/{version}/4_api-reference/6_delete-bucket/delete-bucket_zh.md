# 删除 Bucket

按 ID 删除 bucket。

## 认证

需要 `X-Glue-Authentication`。

## 请求

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## 响应

响应 body 是被删除的 bucket 对象。
