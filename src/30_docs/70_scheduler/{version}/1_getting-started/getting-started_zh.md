# 快速开始

本指南创建 scheduler、入队延迟任务、列出任务、通过 lease 预留任务，并检查 health。

## 创建 Scheduler

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "display_name": "my-scheduler"
  }'
```

```json
{
  "scheduler": {
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "ready": true,
    "scheduled": 0,
    "inflight": 0,
    "display_name": "my-scheduler"
  }
}
```

## 入队任务

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "task-001",
    "payload": { "type": "send_email" },
    "delay": "60s",
    "labels": ["project:onboarding", "priority:high"]
  }'
```

响应：HTTP `202 Accepted`，响应体为空。

## 列出任务

```json
{
  "scheduled": [
    {
      "id": "task-001",
      "future": "2025-06-21T12:01:01Z",
      "labels": ["project:onboarding", "priority:high"]
    }
  ],
  "inflight": [],
  "scheduled_meta": { "page": 1, "per_page": 25, "total": 1, "total_pages": 1 },
  "inflight_meta": { "page": 1, "per_page": 25, "total": 0, "total_pages": 0 }
}
```

## 预留并确认

```json
{ "worktime": "30s" }
```

```json
{
  "id": "task-001",
  "payload": { "type": "send_email" },
  "lease_expires_at": "2025-06-21T12:02:01Z",
  "labels": ["project:onboarding", "priority:high"]
}
```

通过删除任务确认完成。响应为 HTTP `204 No Content`。

## Health

```json
{
  "status": "ok",
  "ready": true,
  "scheduled": 0,
  "inflight": 0,
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```
