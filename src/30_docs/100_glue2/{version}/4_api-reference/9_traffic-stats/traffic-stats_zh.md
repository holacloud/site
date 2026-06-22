# 流量统计

返回按 host 分组的统计数组。每个元素使用 Glue2 的 `prettyStats` 字段，并包含 `*` 和 `-noroute-` 聚合行。

```bash
curl -X GET "https://api.hola.cloud/v0/stats"
```

```json
[
  {
    "host": "my-project.hola.cloud",
    "served_requests": 154203,
    "serving_time": "32.4s",
    "latency_avg": "210µs",
    "uptime": "24h0m0s",
    "start_timestamp": "2026-06-21T10:00:00Z"
  }
]
```
