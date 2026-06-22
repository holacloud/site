# 下载文件

从 bucket 下载文件。文件路径位于 `/files/` 后面。

## 认证

需要 `X-Glue-Authentication`。

## Query 参数

| 参数 | 说明 |
|-----------|-------------|
| `metadata` | 如果存在，返回 file 对象 JSON，而不是文件 body。 |

## 请求

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -o logo.png
```

## 响应

响应 body 是文件内容。`Content-Type` 来自 `mime_type`。

## Metadata 响应

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png?metadata" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

返回包含已实现字段的 file 对象。
