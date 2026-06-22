# 管理 Buckets

Bucket 是文件容器。已实现的操作是创建、列出、获取和删除。

## Bucket 输入

`POST /v1/buckets` 接受 `name` 和 `description`。

- `name` 会被 trim，可以为空，可包含字母、数字、`_` 和 `-`，最长 64 个字符。
- `description` 最长 4096 个字符。
- 如果 `name` 为空，会使用生成的 bucket ID 作为名称。

## 创建 Bucket

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: application/json" \
  -d '{"name":"assets","description":"Application assets"}'
```

## 列出 Buckets

```bash
curl "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

响应是 JSON array，包含 `id`, `name`, `description`, `created_timestamp`, `created_h`, `owners`。

## 获取详情

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

响应是 bucket 对象，包含 `id`, `project_id`, `created_timestamp`, `owners`, `name`, `description`。

## 修改 Bucket

`PATCH /v1/buckets/{bucket_id}` 已在 router 注册，但 handler 未实现。

## 删除 Bucket

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

响应 body 是被删除的 bucket 对象。
