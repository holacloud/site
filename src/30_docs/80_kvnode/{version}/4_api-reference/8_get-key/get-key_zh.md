# 获取键

获取集合中键的值。路径中的通配符 `*` 将被替换为键名。

## 身份验证

需要内部身份验证。通过 `X-Glue-Authentication` 头部或 `apikey` 和 `secret` 头部传递凭据。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| col | string | 集合名称 |

## 请求示例

```bash
curl "https://api.hola.cloud/v1/collections/users/keys/user:1001" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "key": "user:1001",
  "collection": "users",
  "value": {
    "name": "Alice",
    "email": "alice@example.com",
    "role": "admin"
  },
  "created_at": "2025-06-21T10:00:00Z",
  "updated_at": "2025-06-21T10:00:00Z"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | 未找到集合或键 |
| 500 | internal_error | 服务器内部错误 |
