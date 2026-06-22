# 获取 Lambda

按 ID 获取指定 lambda。

## 认证

需要 `X-Glue-Authentication`。

## Path 参数

| 参数 | 类型 | 描述 |
|------|------|------|
| `lambda_id` | string | Lambda 标识符 |

## HTTP 请求

```http
GET /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## 示例

```bash
curl -X GET "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## 响应

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "created_timestamp": 1751378400,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-world",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Hello, World!\" } })",
  "method": "GET",
  "path": "/hello-world"
}
```

## 错误码

| 代码 | 描述 |
|------|------|
| 401 | 认证缺失或无效 |
| 404 | Lambda 未找到 |
