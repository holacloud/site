
# POST /api/v0/lambdas

使用指定的代码和运行时配置创建新的 Lambda 函数。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 请求头。

## 请求体

| 字段 | 类型 | 描述 |
|-------|------|------|
| name | string | Lambda 函数的可读名称 |
| runtime | string | 运行环境：`javascript`、`python` 或 `go` |
| code | string | 函数的源代码 |
| active | boolean | Lambda 是否激活（可选，默认 `true`） |

## HTTP 请求

```http
POST /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "name": "hello-world",
  "runtime": "javascript",
  "active": true,
  "code": "export default async (req) => { return { status: 200, body: { message: 'Hello, World!' } }; }"
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "runtime": "javascript",
    "active": true,
    "code": "export default async (req) => { return { status: 200, body: { message: \"Hello, World!\" } }; }"
  }'
```

## 响应

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "name": "hello-world",
  "runtime": "javascript",
  "active": true,
  "created_at": "2025-07-01T14:00:00Z"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 400 | 缺少或无效的必填字段 |
| 401 | 缺少或无效的身份验证标头 |
| 409 | 同名 Lambda 函数已存在 |
