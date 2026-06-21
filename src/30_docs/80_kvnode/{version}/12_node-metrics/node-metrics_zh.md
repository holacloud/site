# GET /v1/metrics

返回 KVNode 的操作指标，包括写入和读取计数。

## 身份验证

需要内部身份验证。通过 `X-Glue-Authentication` 头部或 `apikey` 和 `secret` 头部传递凭据。

## 请求示例

```bash
curl "https://api.hola.cloud/v1/metrics" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "writes_total": 15000,
  "reads_total": 98000,
  "collections": 3,
  "keys_total": 4500,
  "uptime_seconds": 86400
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 401 | unauthorized | 缺少或无效的身份验证 |
| 500 | internal_error | 服务器内部错误 |
