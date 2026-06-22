<h1>Lambda 入门</h1>

本指南介绍如何创建 JavaScript lambda、列出它，并通过管理路由和公开路由调用它。

## 前提条件

- 一个带有 `X-Glue-Authentication` token 的 HolaCloud 账户。
- 本机已安装 [curl](https://curl.se/)。

## 步骤 1：创建 Lambda

创建一个简单的 `hello-world` lambda：

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "language": "javascript",
    "method": "GET",
    "path": "/hello-world",
    "code": "export default (req) => ({ body: { message: \"Hello, World!\", method: req.method, path: req.path } })"
  }'
```

预期响应：

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "created_timestamp": 1750507200,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-world",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Hello, World!\", method: req.method, path: req.path } })",
  "method": "GET",
  "path": "/hello-world"
}
```

保存返回的 `id`；它就是运行、获取、更新和删除端点使用的 `lambda_id`。

## 步骤 2：列出 Lambdas

确认 lambda 已创建：

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## 步骤 3：运行 Lambda

通过经过认证的管理路由调用它：

```bash
curl -X GET "https://api.hola.cloud/api/v0/run/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

预期响应：

```json
{
  "body": {
    "message": "Hello, World!",
    "method": "GET",
    "path": "/hello-world"
  }
}
```

也可以通过公开运行路由调用 lambda：

```bash
curl -X GET "https://api.hola.cloud/run/YOUR_LAMBDA_ID"
```

## 步骤 4：查看进行中的调用

列出当前正在运行的调用：

```bash
curl "https://api.hola.cloud/ongoing"
```

## 下一步

- 在 [管理 Lambda 函数](../2_managing-functions/managing-functions_zh.md) 中学习如何更新代码和路由字段。
- 在 [调用 Lambda 函数](../3_invoking-functions/invoking-functions_zh.md) 中了解公开运行路由、mux 路由、`/me` 和 `/openapi`。
