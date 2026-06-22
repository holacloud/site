# 修改 Lambda

更新现有 lambda 的支持字段。

## 认证

需要 `X-Glue-Authentication`。

## Path 参数

| 参数 | 类型 | 描述 |
|------|------|------|
| `lambda_id` | string | Lambda 标识符 |

## 请求体

| 字段 | 类型 | 描述 |
|------|------|------|
| `name` | string | 新的 lambda 名称 |
| `language` | string | `javascript`、`static-html`、`static-css` 或 `static-js` |
| `code` | string | 新源码或静态内容 |
| `method` | string | 新 HTTP 方法 |
| `path` | string | 新 HTTP path |

## HTTP 请求

```http
PATCH /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
Content-Type: application/json

{
  "name": "hello-updated",
  "method": "POST",
  "path": "/hello-updated",
  "code": "export default (req) => ({ body: { message: 'Updated lambda', data: req.body } })"
}
```

## 示例

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-updated",
    "method": "POST",
    "path": "/hello-updated",
    "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })"
  }'
```

## 响应

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "created_timestamp": 1751378400,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-updated",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })",
  "method": "POST",
  "path": "/hello-updated"
}
```

## 错误码

| 代码 | 描述 |
|------|------|
| 400 | 请求体无效 |
| 401 | 认证缺失或无效 |
| 404 | Lambda 未找到 |
