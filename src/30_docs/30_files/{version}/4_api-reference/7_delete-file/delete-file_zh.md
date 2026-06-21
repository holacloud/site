# DELETE /v1/buckets/{id}/files/*

从存储桶中删除文件。文件路径在 URL 中 `/files/` 之后指定。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `id` | string | 存储桶 ID（例如 `bkt_abc123`） |

## 请求

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/logo.png" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

## 响应

```http
HTTP/1.1 204 No Content
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 404 | Not Found | 指定的存储桶或文件不存在 |
| 500 | Internal Server Error | 发生意外错误 |
