
# POST /api/v0/run/{id}

使用管理员凭证调用 Lambda 函数。需要身份验证。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 请求头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|------|
| id | uuid | 要运行的 Lambda 函数的唯一标识符 |

## 请求体

您想要传递给 Lambda 函数的任何 JSON 负载。Lambda 函数将其作为 `req` 参数接收。

## HTTP 请求

```http
POST /api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "name": "Alice"
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice"
  }'
```

## 响应

```json
{
  "status": 200,
  "body": {
    "message": "你好，Alice！"
  }
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 400 | 无效的请求体 |
| 401 | 缺少或无效的身份验证标头 |
| 404 | 未找到 Lambda 函数或函数未激活 |
| 500 | Lambda 执行错误 |
