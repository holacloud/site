# POST /v1/buckets

创建新的存储桶。存储桶名称在所有账户中必须全局唯一。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 请求体

```json
{
  "name": "my-new-bucket",
  "public": false
}
```

| 字段 | 类型 | 必填 | 描述 |
|-------|------|----------|-------------|
| `name` | string | 是 | 全局唯一的存储桶名称（3-63个字符，小写字母、数字和连字符） |
| `public` | boolean | 否 | 此存储桶中的文件是否可以无需身份验证访问（默认：`false`） |

## 请求

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-new-bucket",
    "public": false
  }'
```

## 响应

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "bkt_xyz789",
  "name": "my-new-bucket",
  "createdAt": "2026-06-21T12:00:00Z",
  "size": 0,
  "fileCount": 0,
  "public": false
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 400 | Bad Request | 存储桶名称无效或缺少必填字段 |
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 409 | Conflict | 此名称的存储桶已存在 |
| 500 | Internal Server Error | 发生意外错误 |
