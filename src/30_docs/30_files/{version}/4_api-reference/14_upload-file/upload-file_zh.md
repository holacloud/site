# 上传文件

上传文件到 bucket。文件路径位于 `/files/` 后面。

## 认证

需要 `X-Glue-Authentication`。

## 请求

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: image/png" \
  --data-binary @logo.png
```

## 响应

返回 file 对象，字段为 `id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256`, `mime_type`。
