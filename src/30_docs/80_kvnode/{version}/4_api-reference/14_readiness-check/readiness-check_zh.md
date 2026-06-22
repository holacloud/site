# 就绪检查Check

返回 KVNode 的就绪状态，包括其父连接的运行状况。

## 身份验证

此端点为公开，无需身份验证。

## 请求示例

```bash
curl "https://api.hola.cloud/readyz"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "status": "ok",
  "parent_connected": true
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 503 | unavailable | 节点未就绪 |
