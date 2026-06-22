# 列出Collections

返回 KVNode 中所有集合的列表。

## 身份验证

需要内部身份验证。通过 `X-Glue-Authentication` 头部或 `apikey` 和 `secret` 头部传递凭据。

## 请求示例

```bash
curl "https://api.hola.cloud/v1/collections" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "collections": ["users", "sessions"]
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 403 | forbidden | Missing authentication headers |
| 500 | internal_error | 服务器内部错误 |
