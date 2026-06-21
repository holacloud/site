# GET /v0/status

返回所有后端服务的健康状态。

## 描述

此端点对每个注册的后端服务执行健康检查并报告其当前状态。监控系统和控制台仪表板使用此端点。

## 身份验证

无。此端点为公开端点。

## 请求

无需请求体。

## 示例

```bash
curl -X GET "https://api.hola.cloud/v0/status"
```

## 响应

```json
{
  "services": [
    {
      "name": "inceptiondb",
      "status": "healthy",
      "latency_ms": 5,
      "last_checked": "2026-06-21T10:00:00Z"
    },
    {
      "name": "lambda",
      "status": "healthy",
      "latency_ms": 12,
      "last_checked": "2026-06-21T10:00:00Z"
    },
    {
      "name": "files",
      "status": "degraded",
      "latency_ms": 350,
      "last_checked": "2026-06-21T10:00:00Z"
    },
    {
      "name": "kvnode",
      "status": "healthy",
      "latency_ms": 3,
      "last_checked": "2026-06-21T10:00:00Z"
    }
  ],
  "overall": "degraded"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 成功返回状态信息 |
