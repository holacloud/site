# WorkingWithFiles

本指南涵盖 HolaCloud Files 中的文件操作：上传、下载、列出、获取元数据和删除文件。

## 文件路径

文件路径在 URL 中 `/files/` 或 `/list/` 之后指定。路径可以使用正斜杠包含目录：

```
/v1/buckets/{id}/files/images/logo.png
/v1/buckets/{id}/files/backups/db-2026-06-21.sql.gz
```

无需预先创建目录 -- 文件路径被视为带有斜杠的平面键。

## 上传文件

使用 PUT 上传文件。将 `Content-Type` 标头设置为适当的 MIME 类型。

### 上传文本文件

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/notes/readme.txt" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: text/plain" \
  --data-binary @readme.txt
```

### 上传二进制文件（图片）

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/photo.jpg" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: image/jpeg" \
  --data-binary @photo.jpg
```

### 上传带有元数据

您可以使用 `File-Metadata` 标头为文件附加自定义元数据：

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/report.pdf" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: application/pdf" \
  -H "File-Metadata: {\"author\": \"张三\", \"department\": \"工程部\"}" \
  --data-binary @report.pdf
```

预期响应：

```json
{
  "path": "report.pdf",
  "size": 245760,
  "contentType": "application/pdf",
  "uploadedAt": "2026-06-21T12:00:00Z",
  "etag": "\"abc123def456\"",
  "metadata": {
    "author": "张三",
    "department": "工程部"
  }
}
```

## 下载文件

使用 GET 通过文件路径下载：

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/photo.jpg" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -o photo.jpg
```

HTTP 请求：

```http
GET /v1/buckets/bkt_abc123/files/images/photo.jpg HTTP/1.1
Host: api.hola.cloud
Api-Key: 您的API密钥
Api-Secret: 您的API密钥密码
```

响应包含文件内容以及适当的 `Content-Type` 和 `Content-Length` 标头。

### 部分下载（范围请求）

Files 支持 HTTP 范围请求以实现部分下载：

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/video.mp4" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Range: bytes=0-1023" \
  -o video_chunk.mp4
```

## 列出文件

使用可选的前缀过滤器列出存储桶中的文件。`/list/` 后的 `*` 是必需的，表示所有文件。使用特定路径按前缀过滤。

### 列出所有文件

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

### 列出目录中的文件

列出 `images/` 前缀下的文件：

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/images/*" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

响应：

```json
{
  "files": [
    {
      "path": "images/photo.jpg",
      "size": 245760,
      "contentType": "image/jpeg",
      "modifiedAt": "2026-06-21T12:00:00Z"
    },
    {
      "path": "images/banner.png",
      "size": 512000,
      "contentType": "image/png",
      "modifiedAt": "2026-06-21T11:30:00Z"
    }
  ],
  "prefix": "images/",
  "total": 2
}
```

### 分页列出

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*?limit=10&offset=0" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

## 获取文件元数据

使用 HEAD 获取文件元数据，无需下载内容：

```bash
curl -I "https://api.hola.cloud/v1/buckets/bkt_abc123/files/report.pdf" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

响应标头：

```http
HTTP/1.1 200 OK
Content-Type: application/pdf
Content-Length: 245760
ETag: "abc123def456"
Last-Modified: Sun, 21 Jun 2026 12:00:00 GMT
File-Metadata: {"author": "张三", "department": "工程部"}
```

## 删除文件

按路径删除文件：

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123/files/notes/readme.txt" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

预期响应：HTTP 204 无内容。

### 错误：文件未找到

```json
{
  "error": {
    "code": "FILE_NOT_FOUND",
    "message": "在存储桶 'bkt_abc123' 中未找到文件 'notes/readme.txt'。"
  }
}
```

## Content-Type 处理

上传文件时，您应将 `Content-Type` 标头设置为文件的正确 MIME 类型。如果省略，服务默认使用 `application/octet-stream`。

常见 MIME 类型：

| 扩展名 | Content-Type |
|-----------|-------------|
| .txt | text/plain |
| .html | text/html |
| .css | text/css |
| .js | application/javascript |
| .json | application/json |
| .jpg / .jpeg | image/jpeg |
| .png | image/png |
| .gif | image/gif |
| .pdf | application/pdf |
| .zip | application/zip |
| .gz | application/gzip |
| .mp4 | video/mp4 |

## 文件路径约定

- 文件路径区分大小写
- 不允许前导斜杠：`notes/readme.txt`（正确），`/notes/readme.txt`（错误）
- 路径最多可达 1024 个字符
- 使用正斜杠（`/`）作为目录分隔符
- 路径段中的特殊字符应进行 URL 编码

## HTTP 状态码汇总

| 状态码 | 描述 |
|--------|------|
| 200 | 成功 (GET, HEAD, PUT) |
| 204 | 无内容 (DELETE) |
| 400 | 错误请求 |
| 404 | 未找到 -- 文件或存储桶不存在 |
| 416 | 范围不满足要求 |
| 500 | 服务器内部错误 |
