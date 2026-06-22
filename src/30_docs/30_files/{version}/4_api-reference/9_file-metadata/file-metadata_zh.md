# 文件 Headers

使用 `HEAD /v1/buckets/{bucket_id}/files/*` 返回文件 headers。

## 认证

需要 `X-Glue-Authentication`。

## 请求

```bash
curl -I "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## 响应

handler 设置 `Last-Modified` 和 `Content-Length`。
