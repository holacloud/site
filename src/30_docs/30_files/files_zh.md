# Files

HolaCloud Files 提供 REST API，用 bucket 保存文件。文件路径位于 `/files/` 之后，列表接口使用 `/list/` 之后的路径作为前缀过滤。

## 认证

所有 `/v1` endpoint 都需要 `X-Glue-Authentication` header：

```http
X-Glue-Authentication: {"user":{"id":"user-123"}}
```

## 数据模型

Bucket 字段：`id`, `project_id`, `created_timestamp`, `owners`, `name`, `description`。

File 字段：`id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256`, `mime_type`。

## API 参考

所有 endpoint 使用 base URL `https://api.hola.cloud`。

| 方法 | 路径 | 说明 |
|--------|------|-------------|
| GET | `/v1/buckets` | 列出认证用户的 buckets |
| POST | `/v1/buckets` | 创建 bucket |
| GET | `/v1/buckets/{bucket_id}` | 获取 bucket 详情 |
| PATCH | `/v1/buckets/{bucket_id}` | 未实现 |
| DELETE | `/v1/buckets/{bucket_id}` | 删除 bucket |
| GET | `/v1/buckets/{bucket_id}/list/*` | 列出文件，可按前缀过滤 |
| PUT | `/v1/buckets/{bucket_id}/files/*` | 上传文件 |
| GET | `/v1/buckets/{bucket_id}/files/*` | 下载文件，或使用 `?metadata` 返回 JSON metadata |
| DELETE | `/v1/buckets/{bucket_id}/files/*` | 删除文件 |
| HEAD | `/v1/buckets/{bucket_id}/files/*` | 返回文件 headers |
