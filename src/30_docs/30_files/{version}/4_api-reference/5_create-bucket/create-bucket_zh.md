# 创建 Bucket

为认证用户创建 bucket。

## 认证

需要 `X-Glue-Authentication`。

## 请求 Body

```json
{
  "name": "my-bucket",
  "description": "Optional description"
}
```

字段：`name` 和 `description`。

## 请求

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: application/json" \
  -d '{"name":"my-bucket","description":"Optional description"}'
```

## 响应

响应是 bucket 对象，字段为 `id`, `project_id`, `created_timestamp`, `owners`, `name`, `description`。
