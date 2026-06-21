# 任务生命周期

本文档详细介绍任务的完整生命周期：入队、列出、预留、延长租约、确认、重新调度以及通过SSE实时流。

## 任务状态

任务经历以下状态：

- **pending** — 等待计划时间。对工作者尚不可见。
- **available** — 计划时间已到；可被预留。
- **reserved** — 由工作者在租约下持有。
- **completed** — 已被工作者确认并移除。

如果租约到期未确认，任务将返回 **available** 状态。

## 入队任务

创建一个带有负载、未来时间（或延迟）和可选标签的新任务：

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "payload": {
      "action": "generate_report",
      "report_id": "rpt-5678"
    },
    "future": "2025-06-22T00:00:00Z",
    "labels": {
      "team": "analytics",
      "env": "production"
    }
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "payload": {
    "action": "generate_report",
    "report_id": "rpt-5678"
  },
  "future": "2025-06-22T00:00:00Z",
  "labels": {
    "team": "analytics",
    "env": "production"
  }
}
```

或者，使用 `delay`（距现在的秒数）代替 `future`：

```json
{
  "payload": { "action": "send_reminder" },
  "delay": 3600,
  "labels": { "type": "email" }
}
```

预期响应：

```json
{
  "id": "task-001",
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "payload": { "action": "generate_report", "report_id": "rpt-5678" },
  "state": "pending",
  "available_at": "2025-06-22T00:00:00Z",
  "labels": { "team": "analytics", "env": "production" }
}
```

## 分页列出任务

列出任务，支持可选的分页和标签过滤：

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?offset=0&limit=20"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?offset=0&limit=20 HTTP/1.1
Host: api.hola.cloud
```

预期响应：

```json
{
  "tasks": [
    {
      "id": "task-001",
      "state": "pending",
      "available_at": "2025-06-22T00:00:00Z",
      "labels": { "team": "analytics" }
    },
    {
      "id": "task-002",
      "state": "available",
      "available_at": "2025-06-21T12:00:00Z",
      "labels": { "team": "analytics" }
    }
  ],
  "total": 2,
  "offset": 0,
  "limit": 20
}
```

## 预留任务

工作者通过指定`worktime`（租约持续时间，以秒为单位）来预留可用任务：

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 60
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "worktime": 60
}
```

预期响应：

```json
{
  "id": "task-002",
  "payload": { "action": "send_reminder" },
  "lease_expires_at": "2025-06-21T12:01:00Z"
}
```

如果没有可用任务，端点返回HTTP `204 No Content`。

## 延长租约

如果工作者需要更多时间，请在租约到期前延长：

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/extend" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 30
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/extend HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "worktime": 30
}
```

预期响应：

```json
{
  "id": "task-002",
  "lease_expires_at": "2025-06-21T12:01:30Z"
}
```

## 确认（删除）任务

处理成功后，确认任务以将其移除：

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002" \
  -H "X-API-Key: YOUR_API_KEY"
```

```http
DELETE /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002 HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
```

预期响应：HTTP `204 No Content`。

## 重新调度任务

如果处理失败，使用新的延迟或未来时间重新调度任务：

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/reschedule" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "delay": 300
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/reschedule HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "delay": 300
}
```

预期响应：

```json
{
  "id": "task-002",
  "state": "pending",
  "available_at": "2025-06-21T12:06:00Z"
}
```

这非常适合实现带有指数退避的重试逻辑。

## 用于实时更新的SSE流

通过服务器发送事件订阅实时任务事件：

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/stream"
```

流以SSE格式发送事件：

```
event: task_available
data: {"id":"task-003","payload":{"action":"process_image"},"labels":{}}

event: task_completed
data: {"id":"task-003"}

event: task_expired
data: {"id":"task-003"}
```

| 事件 | 描述 |
|------|------|
| `task_available` | 任务已可用于预留 |
| `task_completed` | 工作者已确认任务 |
| `task_expired` | 租约已到期，任务再次可用 |

SSE对于构建响应式工作者池非常有用，无需轮询即可立即响应任务。
