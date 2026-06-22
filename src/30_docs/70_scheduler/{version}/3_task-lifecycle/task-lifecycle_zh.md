# 任务生命周期

Scheduler 维护一个内存延迟队列。任务使用 `id`、`future` 或 `delay`、`payload` 以及字符串数组 `labels` 入队。工作者通过 lease 预留任务，并通过删除任务来确认完成。

## 入队

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "task-001",
    "payload": { "action": "generate_report" },
    "delay": "1h",
    "labels": ["team:analytics", "env:production"]
  }'
```

响应：HTTP `202 Accepted`，响应体为空。

## 列表

```json
{
  "scheduled": [
    {
      "id": "task-001",
      "future": "2025-06-22T00:00:00Z",
      "labels": ["team:analytics"]
    }
  ],
  "inflight": [],
  "scheduled_meta": { "page": 1, "per_page": 25, "total": 1, "total_pages": 1 },
  "inflight_meta": { "page": 1, "per_page": 25, "total": 0, "total_pages": 0 }
}
```

## 预留

```json
{ "worktime": "60s" }
```

```json
{
  "id": "task-001",
  "payload": { "action": "generate_report" },
  "lease_expires_at": "2025-06-21T12:01:00Z",
  "labels": ["team:analytics"]
}
```

没有可用任务时，API 返回 `204 No Content`。

## 延长 Lease

```json
{ "extension": "30s" }
```

```json
{ "lease_expires_at": "2025-06-21T12:01:30Z" }
```

## 确认

工作者通过删除任务确认完成：

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-001" \
  -H "X-API-Key: YOUR_API_KEY"
```

响应：HTTP `204 No Content`。

## 重新调度

```json
{ "delay": "5m" }
```

```json
{
  "id": "task-001",
  "future": "2025-06-21T12:06:00Z"
}
```

## SSE

SSE stream 只发送 `snapshot` 事件，包含 `scheduled`、`inflight`、分页 meta 和 `health`。
