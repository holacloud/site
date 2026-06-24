# OpenAPI 规范

返回 KVNode API 的 OpenAPI 规范。

## 身份验证

此端点为公开，无需身份验证。

## 请求示例

```bash
curl "https://api.hola.cloud/openapi.json"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "openapi": "3.0.0",
  "info": {
    "title": "KVNode API",
    "version": "1.0.0"
  },
  "paths": {
    "/healthz": {
      "get": {
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "Node is healthy"
          }
        }
      }
    }
  }
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 500 | internal_error | 服务器内部错误 |
