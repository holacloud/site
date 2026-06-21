# GET /openapi.json

返回 Glue2 网关的 OpenAPI 规范。

## 描述

此端点提供 OpenAPI 3.0 规范文档，描述网关暴露的所有路由、请求格式和响应模式。该规范从路由定义自动生成。

## 身份验证

无。此端点为公开端点。

## 请求

无需请求体。

## 示例

```bash
curl -X GET "https://api.hola.cloud/openapi.json"
```

## 响应

```json
{
  "openapi": "3.0.3",
  "info": {
    "title": "Glue2 API Gateway",
    "version": "2.3.1",
    "description": "HolaCloud 服务的中央 API 网关"
  },
  "paths": {
    "/version": {
      "get": {
        "summary": "获取网关版本",
        "responses": {
          "200": {
            "description": "版本信息"
          }
        }
      }
    },
    "/v0/virtualhosts": {
      "get": {
        "summary": "列出虚拟主机",
        "responses": {
          "200": {
            "description": "路由表"
          }
        }
      }
    },
    "/v0/stats": {
      "get": {
        "summary": "获取流量统计",
        "responses": {
          "200": {
            "description": "流量统计信息"
          }
        }
      }
    },
    "/v0/status": {
      "get": {
        "summary": "后端健康状态",
        "responses": {
          "200": {
            "description": "服务状态"
          }
        }
      }
    }
  }
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 成功返回 OpenAPI 规范 |
