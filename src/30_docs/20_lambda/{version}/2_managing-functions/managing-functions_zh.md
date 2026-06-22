<h1>管理 Lambda 函数</h1>

创建 lambda 后，可以使用管理 API 查看它、更新支持的字段、列出账户中的所有 lambdas，或删除它。

## 函数结构

JavaScript lambdas 导出默认处理器。处理器接收请求对象，并返回希望 HolaCloud 发回的响应 body。

```javascript
export default (req) => {
  return {
    body: {
      method: req.method,
      path: req.path,
      headers: req.headers,
      data: req.body
    }
  };
};
```

静态 lambdas 使用一种静态语言模式：`static-html`、`static-css` 或 `static-js`。在这些模式中，`code` 是对应 lambda 提供的内容。

## 更新 Lambda

使用 `PATCH /api/v0/lambdas/{lambda_id}` 更新 `name`、`language`、`code`、`method` 或 `path`。

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-updated",
    "method": "POST",
    "path": "/hello-updated",
    "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })"
  }'
```

预期响应：

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "created_timestamp": 1750507200,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-updated",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })",
  "method": "POST",
  "path": "/hello-updated"
}
```

## 查看 Lambda 详情

按 ID 获取一个 lambda：

```bash
curl "https://api.hola.cloud/api/v0/lambdas/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## 列出所有 Lambdas

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## 删除 Lambda

永久删除一个 lambda：

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

删除成功后会返回已删除 lambda 的 API 响应或确认 payload，具体取决于已部署的 API 版本。
