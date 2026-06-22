# 列出Keys

列出集合中的键，支持分页和前缀过滤。

## 身份验证

需要内部身份验证。通过 `X-Glue-Authentication` 头部或 `apikey` 和 `secret` 头部传递凭据。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| col | string | 集合名称 |

## 查询参数

| 参数 | 类型 | 描述 |
|------|------|------|
| limit | integer | 返回的最大键数（默认：100） |
| prefix | string | 按前缀过滤键 |

## 请求示例

```bash
curl "https://api.hola.cloud/v1/collections/users/keys?prefix=user:&limit=10" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "items": [
    {
      "key": "user:1001",
      "value": { "name": "Alice" },
      "version": 1,
      "updatedAt": "2025-06-20T15:30:00Z"
    },
    {
      "key": "user:1002",
      "value": { "name": "Bob" },
      "version": 1,
      "updatedAt": "2025-06-21T08:00:00Z"
    }
  ],
  "next": "user:1002"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | 未找到集合 |
| 500 | internal_error | 服务器内部错误 |
