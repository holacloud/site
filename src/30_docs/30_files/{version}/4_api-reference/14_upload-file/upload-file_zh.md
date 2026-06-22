# 上传文件

上传文件到存储桶。文件路径在 URL 中 `/files/` 之后指定。最大文件大小为 5 GB。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `id` | string | 存储桶 ID（例如 `bkt_abc123`） |

## 请求标头

| 标头 | 必填 | 描述 |
|--------|----------|-------------|
| `Content-Type` | 是 | 上传文件的 MIME 类型 |
| `Content-Length` | 是 | 文件大小（字节） |

## 请求

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/logo.png" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: image/png" \
  --data-binary @logo.png
```

## 响应

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "path": "images/logo.png",
  "size": 24576,
  "contentType": "image/png",
  "uploadedAt": "2026-06-21T12:00:00Z",
  "etag": "\"abc123def456\""
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 400 | Bad Request | 请求无效或缺少 Content-Type |
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 404 | Not Found | 指定的存储桶不存在 |
| 413 | Payload Too Large | 文件超过 5 GB 大小限制 |
| 500 | Internal Server Error | 发生意外错误 |
