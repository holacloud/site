# 下载文件

从存储桶下载文件。文件路径在 URL 中 `/files/` 之后指定。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `id` | string | 存储桶 ID（例如 `bkt_abc123`） |

## 查询参数

| 参数 | 类型 | 默认值 | 描述 |
|-----------|------|---------|-------------|
| `metadata` | boolean | false | 设置为 `true` 以返回 JSON 格式的文件元数据而非文件内容 |

## 请求

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/logo.png" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -o logo.png
```

## 响应

```http
HTTP/1.1 200 OK
Content-Type: image/png
Content-Length: 24576
Last-Modified: Sun, 21 Jun 2026 12:00:00 GMT
ETag: "abc123def456"
```

响应体为原始文件内容。

### 元数据响应

当指定 `?metadata=true` 时，响应返回 JSON 而非文件内容：

```json
{
  "path": "images/logo.png",
  "size": 24576,
  "contentType": "image/png",
  "modifiedAt": "2026-06-21T12:00:00Z",
  "etag": "\"abc123def456\""
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 404 | Not Found | 指定的存储桶或文件不存在 |
| 500 | Internal Server Error | 发生意外错误 |
