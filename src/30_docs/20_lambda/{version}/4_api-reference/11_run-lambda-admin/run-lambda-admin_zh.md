# 运行 Lambda（管理）

通过经过认证的管理路由调用 lambda。该端点接受任何 HTTP 方法。

## 认证

需要 `X-Glue-Authentication`。

## Path 参数

| 参数 | 类型 | 描述 |
|------|------|------|
| `lambda_id` | string | Lambda 标识符 |

## 请求体

发送任何希望 lambda 作为 `req.body` 接收的 payload。

## HTTP 请求

```http
POST /api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
Content-Type: application/json

{
  "name": "Alice"
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice"
  }'
```

## 响应

响应就是 lambda 处理器返回的内容。

```json
{
  "body": {
    "message": "Hello, Alice!"
  }
}
```

## 错误码

| 代码 | 描述 |
|------|------|
| 400 | 请求体无效 |
| 401 | 认证缺失或无效 |
| 404 | Lambda 未找到 |
| 500 | Lambda 执行错误 |
