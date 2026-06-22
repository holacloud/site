# 节点状态

返回 KVNode 的当前状态，包括集合、复制状态和运行时间。

## 身份验证

需要内部身份验证。通过 `X-Glue-Authentication` 头部或 `apikey` 和 `secret` 头部传递凭据。

## 请求示例

```bash
curl "https://api.hola.cloud/v1/status" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "node_id": "node-abc123",
  "collections": 3,
  "uptime_seconds": 86400,
  "replication": {
    "role": "primary",
    "connected_replicas": 2
  }
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 401 | unauthorized | 缺少或无效的身份验证 |
| 500 | internal_error | 服务器内部错误 |
