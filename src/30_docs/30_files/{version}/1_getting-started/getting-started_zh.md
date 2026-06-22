# 入门

本指南演示创建 bucket、上传文件、列出文件、下载和删除。

## 前提条件

- 本地安装 `curl`
- 有效的 `X-Glue-Authentication` header

```http
X-Glue-Authentication: {"user":{"id":"user-123"}}
```

所有请求使用 `https://api.hola.cloud`。

## 第 1 步：创建 Bucket

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: application/json" \
  -d '{"name":"my-first-bucket","description":"First test bucket"}'
```

响应是 bucket 对象，字段为 `id`, `project_id`, `created_timestamp`, `owners`, `name`, `description`。

## 第 2 步：上传文件

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/hello.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: text/plain" \
  --data-binary "Hello, HolaCloud Files!"
```

响应是 file 对象，字段为 `id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256`, `mime_type`。

## 第 3 步：列出文件

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/list/*" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

响应是 file 对象的 JSON array。

## 第 4 步：下载文件

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/hello.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## 第 5 步：删除文件

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/hello.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## 第 6 步：删除 Bucket

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```
