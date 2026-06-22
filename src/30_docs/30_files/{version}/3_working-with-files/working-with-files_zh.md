# 使用文件

文件通过 `/files/` 后面的路径定位。列表接口使用 `/list/` 后面的路径作为前缀过滤。

## 上传文件

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: text/plain" \
  --data-binary @readme.txt
```

响应是包含已实现字段的 file 对象。

## 下载文件

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -o readme.txt
```

响应 body 是保存的文件内容。`Content-Type` 来自 `mime_type`。

## 获取 JSON Metadata

在下载 endpoint 使用 `?metadata` 可返回 file 对象 JSON。

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt?metadata" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## 列出文件

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/list/notes/" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

响应是 file 对象的 JSON array。

## HEAD 文件

```bash
curl -I "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

handler 设置 `Last-Modified` 和 `Content-Length`。

## 删除文件

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```
