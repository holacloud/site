
# List Lambdas

列出已认证账户的所有 Lambda 函数。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 请求头。

## HTTP 请求

```http
GET /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## 示例

```bash
curl -X GET "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf"
```

## 响应

```json
[
  {
    "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
    "name": "hello-world",
    "runtime": "javascript",
    "active": true,
    "created_at": "2025-05-10T12:00:00Z"
  },
  {
    "id": "a2b3c4d5-e6f7-8901-bcde-f12345678901",
    "name": "process-orders",
    "runtime": "python",
    "active": true,
    "created_at": "2025-06-20T08:30:00Z"
  }
]
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 401 | 缺少或无效的身份验证标头 |
