# 启动检查

返回 KVNode 的启动状态。

## 身份验证

此端点为公开，无需身份验证。

## 请求示例

```bash
curl "https://api.hola.cloud/startupz"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "ok": true,
  "node": "node-abc123",
  "role": "leader"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 500 | internal_error | 节点不健康 |
