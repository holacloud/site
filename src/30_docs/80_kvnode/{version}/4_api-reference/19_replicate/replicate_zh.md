# 复制

从父节点复制数据到此 KVNode。

## 身份验证

需要内部身份验证。通过 `X-Glue-Authentication` 头部或 `apikey` 和 `secret` 头部传递凭据。

## 请求示例

```bash
curl -X POST "https://api.hola.cloud/v1/replicate" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "entries": [
      {
        "collection": "users",
        "key": "user:1001",
        "value": {"name": "Alice"},
        "seq": 42
      }
    ]
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
  "applied": 1
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | invalid_json | 无效的 JSON payload |
| 403 | forbidden | 缺少认证头部 |
| 502 | parent_unavailable | 父节点不可达 |
| 500 | internal_error | 服务器内部错误 |
