
# Run Lambda (Public)

公开调用 Lambda 函数，无需身份验证。

此端点专为 Webhook 和公共 API 端点设计。Lambda ID 可从 Lambda 管理页面获取。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|------|
| id | uuid | 要运行的 Lambda 函数的唯一标识符 |

## 请求体

您想要传递给 Lambda 函数的任何 JSON 负载。

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

```json
{
  "status": 200,
  "body": {
    "processed": true,
    "event": "payment_received"
  }
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 400 | 无效的请求体 |
| 404 | 未找到 Lambda 函数或函数未激活 |
| 500 | Lambda 执行错误 |
