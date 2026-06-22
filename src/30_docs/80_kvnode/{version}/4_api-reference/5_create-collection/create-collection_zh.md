# 创建集合

在 KVNode 中创建一个新的集合。

## 身份验证

需要内部身份验证。通过 `X-Glue-Authentication` 头部或 `apikey` 和 `secret` 头部传递凭据。

## 请求体

```json
{
  "name": "users"
}
```

## 请求示例

```bash
curl -X POST "https://api.hola.cloud/v1/collections" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "users"
  }'
```

## 响应示例

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "name": "users",
  "key_count": 0,
  "created_at": "2025-06-21T10:00:00Z"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | invalid_request | 缺少或无效的集合名称 |
| 401 | unauthorized | 缺少或无效的身份验证 |
| 409 | conflict | 集合已存在 |
| 500 | internal_error | 服务器内部错误 |
