# 设置键

设置集合中键的值。路径中的通配符 `*` 将被替换为键名。

## 身份验证

需要内部身份验证。通过 `X-Glue-Authentication` 头部或 `apikey` 和 `secret` 头部传递凭据。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| col | string | 集合名称 |

## 请求体

| 字段 | 类型 | 描述 |
|------|------|------|
| value | any | 要存储的 JSON 值 |

```json
{
  "value": {
    "name": "Alice",
    "email": "alice@example.com",
    "role": "admin"
  }
}
```

## 请求示例

```bash
curl -X POST "https://api.hola.cloud/v1/collections/users/keys/user:1001" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "value": {
      "name": "Alice",
      "email": "alice@example.com",
      "role": "admin"
    }
  }'
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
 {
  "ok": true,
  "seq": 1,
  "version": 1
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | invalid_json | 缺少或无效的值 |
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | 未找到集合 |
| 500 | internal_error | 服务器内部错误 |
