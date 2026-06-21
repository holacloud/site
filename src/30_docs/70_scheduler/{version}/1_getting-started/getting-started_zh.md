# 调度器快速入门

本指南将带您完成创建调度器、入队延迟任务、列出任务、预留和确认任务以及检查调度器运行状况的过程。

## 前提条件

- 一个带有API凭证的 HolaCloud 账户。
- 在您的机器上安装 [curl](https://curl.se/)。

## 身份验证

某些端点需要通过API密钥进行身份验证。使用 `X-API-Key` 头部或 `Authorization: Bearer <key>` 传递密钥。

## 第一步：创建调度器

创建一个新的调度器：

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "my-scheduler"
  }'
```

预期响应：

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "my-scheduler",
  "task_count": 0,
  "status": "active"
}
```

保存返回的 `id` — 您将在后续请求中使用它。

## 第二步：入队延迟任务

入队一个60秒后可用的任务：

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "payload": {
      "type": "send_email",
      "to": "user@example.com",
      "template": "welcome"
    },
    "delay": 60,
    "labels": {
      "project": "onboarding",
      "priority": "high"
    }
  }'
```

预期响应：

```json
{
  "id": "task-001",
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "payload": {
    "type": "send_email",
    "to": "user@example.com",
    "template": "welcome"
  },
  "state": "pending",
  "available_at": "2025-06-21T12:01:01Z",
  "labels": {
    "project": "onboarding",
    "priority": "high"
  }
}
```

## 第三步：列出任务

查看调度器中的所有任务：

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks"
```

预期响应：

```json
{
  "tasks": [
    {
      "id": "task-001",
      "state": "pending",
      "available_at": "2025-06-21T12:01:01Z",
      "labels": { "project": "onboarding", "priority": "high" }
    }
  ],
  "total": 1
}
```

## 第四步：预留和确认任务

延迟时间过去后，预留该任务：

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 30
  }'
```

预期响应：

```json
{
  "id": "task-001",
  "payload": { "type": "send_email", "to": "user@example.com", "template": "welcome" },
  "lease_expires_at": "2025-06-21T12:02:01Z"
}
```

处理任务，然后确认（删除）它：

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-001" \
  -H "X-API-Key: YOUR_API_KEY"
```

预期响应：HTTP `204 No Content`。

## 第五步：检查调度器运行状况

验证调度器是否正常运行：

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health"
```

预期响应：

```json
{
  "status": "ok",
  "uptime_seconds": 12345
}
```

## 后续步骤

- 了解如何管理调度器，包括更新和监控，请参阅[管理调度器](../2_managing-schedulers/managing-schedulers_zh.md)。
- 深入了解任务生命周期——包括租约延长、重新调度和SSE流——请参阅[任务生命周期](../3_task-lifecycle/task-lifecycle_zh.md)。
