# OpenAPI 规范

返回 Lambda API 的 OpenAPI 规范文档。

## 认证

需要 `X-Glue-Authentication`。

## HTTP 请求

```http
GET /api/v0/openapi HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## 示例

```bash
curl -X GET "https://api.hola.cloud/api/v0/openapi" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## 响应

返回完整的 OpenAPI 规范 JSON 文档。

```json
{
  "openapi": "3.0.3",
  "info": {
    "title": "Lambda API",
    "version": "1.0.0"
  },
  "paths": {
    "/api/v0/lambdas": {
      "get": {
        "summary": "List lambdas",
        "responses": {
          "200": {
            "description": "List of lambdas"
          }
        }
      }
    }
  }
}
```

## 错误码

| 代码 | 描述 |
|------|------|
| 401 | 认证缺失或无效 |
