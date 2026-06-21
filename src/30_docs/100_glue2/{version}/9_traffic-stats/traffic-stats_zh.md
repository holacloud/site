# GET /v0/stats

返回网关的实时流量统计信息。

## 描述

此端点提供通过 Glue2 网关的请求的聚合指标，包括请求速率、延迟百分位数和状态码分布。

## 身份验证

无。此端点为公开端点。

## 请求

无需请求体。

## 示例

```bash
curl -X GET "https://api.hola.cloud/v0/stats"
```

## 响应

```json
{
  "total_requests": 154203,
  "requests_per_second": 42.5,
  "latency_p50_ms": 12,
  "latency_p95_ms": 45,
  "latency_p99_ms": 120,
  "status_codes": {
    "2xx": 148000,
    "4xx": 5200,
    "5xx": 1003
  },
  "active_connections": 87,
  "uptime_seconds": 86400
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 成功返回统计信息 |
