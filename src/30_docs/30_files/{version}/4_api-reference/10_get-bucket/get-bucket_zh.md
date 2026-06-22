# 获取 Bucket

为认证用户按 ID 获取 bucket。

## 认证

需要 `X-Glue-Authentication`。

## 请求

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## 响应

响应是 bucket 对象，字段为 `id`, `project_id`, `created_timestamp`, `owners`, `name`, `description`。
