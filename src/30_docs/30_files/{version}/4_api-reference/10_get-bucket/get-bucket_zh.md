# 获取存储桶

获取指定存储桶的详细信息。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `id` | string | 存储桶 ID（例如 `bkt_abc123`） |

## 请求

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

## 响应

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "bkt_abc123",
  "name": "my-first-bucket",
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 1048576,
  "fileCount": 5,
  "public": false
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 404 | Not Found | 指定的存储桶不存在 |
| 500 | Internal Server Error | 发生意外错误 |
