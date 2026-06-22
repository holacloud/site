# 删除文件

从 bucket 删除文件。文件路径位于 `/files/` 后面。

## 认证

需要 `X-Glue-Authentication`。

## 请求

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## 响应

文件删除成功时返回空的成功响应。
