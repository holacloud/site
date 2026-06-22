# 运行 Lambda（公开）

公开调用 lambda。不需要认证。该端点接受任何 HTTP 方法。

此路由适用于 webhook 和公开 API 调用。lambda ID 来自 lambda 管理端点。

## Path 参数

| 参数 | 类型 | 描述 |
|------|------|------|
| `lambda_id` | string | Lambda 标识符 |

## 请求体

发送任何希望 lambda 作为 `req.body` 接收的 payload。

## HTTP 请求

```http
POST /run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "event": "payment_received",
  "amount": 49.99,
  "currency": "USD"
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Content-Type: application/json" \
  -d '{
    "event": "payment_received",
    "amount": 49.99,
    "currency": "USD"
  }'
```

## 响应

响应就是 lambda 处理器返回的内容。

```json
{
  "body": {
    "processed": true,
    "event": "payment_received"
  }
}
```

## 错误码

| 代码 | 描述 |
|------|------|
| 400 | 请求体无效 |
| 404 | Lambda 未找到 |
| 500 | Lambda 执行错误 |
