# 创建 Lambda

创建带有源码或静态内容以及路由元数据的 lambda。

## 认证

需要 `X-Glue-Authentication`。

## 请求体

| 字段 | 类型 | 描述 |
|------|------|------|
| `name` | string | 可读的 lambda 名称 |
| `language` | string | `javascript`、`static-html`、`static-css` 或 `static-js` |
| `code` | string | 源码或静态内容 |
| `method` | string | lambda 路由的 HTTP 方法 |
| `path` | string | lambda 路由的 HTTP path |

## HTTP 请求

```http
POST /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
Content-Type: application/json

{
  "name": "hello-world",
  "language": "javascript",
  "method": "GET",
  "path": "/hello-world",
  "code": "export default (req) => ({ body: { message: 'Hello, World!' } })"
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "language": "javascript",
    "method": "GET",
    "path": "/hello-world",
    "code": "export default (req) => ({ body: { message: \"Hello, World!\" } })"
  }'
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
| 400 | 必填字段缺失或无效 |
| 401 | 认证缺失或无效 |
| 409 | 已存在同名 lambda |
