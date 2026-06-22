# Files API 参考

Base URL: `https://api.hola.cloud`

所有 `/v1` endpoint 都需要 `X-Glue-Authentication`。

```http
X-Glue-Authentication: {"user":{"id":"user-123"}}
```

## Endpoints

| 方法 | 路径 | 说明 |
|--------|------|-------------|
| GET | `/v1/buckets` | 列出 buckets |
| POST | `/v1/buckets` | 创建 bucket |
| GET | `/v1/buckets/{bucket_id}` | 获取 bucket 详情 |
| PATCH | `/v1/buckets/{bucket_id}` | 未实现 |
| DELETE | `/v1/buckets/{bucket_id}` | 删除 bucket |
| GET | `/v1/buckets/{bucket_id}/list/*` | 按前缀列出文件 |
| PUT | `/v1/buckets/{bucket_id}/files/*` | 上传文件 |
| GET | `/v1/buckets/{bucket_id}/files/*` | 下载文件 |
| DELETE | `/v1/buckets/{bucket_id}/files/*` | 删除文件 |
| HEAD | `/v1/buckets/{bucket_id}/files/*` | 返回文件 headers |

## 字段

Bucket 字段：`id`, `project_id`, `created_timestamp`, `owners`, `name`, `description`。

File 字段：`id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256`, `mime_type`。

## 常见错误

| 状态 | 说明 |
|--------|-------------|
| 401 | 缺少或无效的 `X-Glue-Authentication` |
| 404 | Bucket 或文件不存在 |
| 409 | Bucket 已存在 |
| 500 | 持久化或 filesystem 错误 |
