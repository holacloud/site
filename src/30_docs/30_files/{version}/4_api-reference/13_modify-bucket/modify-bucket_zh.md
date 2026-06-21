# PATCH /v1/buckets/{id}

修改现有存储桶的元数据。仅更新请求体中提供的字段。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `id` | string | 存储桶 ID（例如 `bkt_abc123`） |

## 请求体

```json
{
  "name": "renamed-bucket",
  "public": true
}
```

| 字段 | 类型 | 必填 | 描述 |
|-------|------|----------|-------------|
| `name` | string | 否 | 新的全局唯一存储桶名称 |
| `public` | boolean | 否 | 设置为 `true` 以允许公开访问文件 |

## 请求

```bash
curl -X PATCH "https://api.hola.cloud/v1/buckets/bkt_abc123" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "renamed-bucket",
    "public": true
  }'
```

## 响应

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "bkt_abc123",
  "name": "renamed-bucket",
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 1048576,
  "fileCount": 5,
  "public": true
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 400 | Bad Request | 请求体无效 |
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 404 | Not Found | 指定的存储桶不存在 |
| 409 | Conflict | 新名称已被占用 |
| 500 | Internal Server Error | 发生意外错误 |
