# OpenAPI 规范

返回 Run 服务的 OpenAPI 规范。

## 请求

```http
GET /openapi.json
```

## 示例

```bash
curl "https://api.hola.cloud/openapi.json"
```

## 响应

以 JSON 格式返回 OpenAPI 规范。

```json
{
  "openapi": "3.0.0",
  "info": {
    "title": "HolaCloud Run API",
    "version": "1.0.0"
  },
  "paths": {}
}
```
