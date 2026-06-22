# 入队任务

将一个新任务入队，以便在预定时间执行。

## 身份验证

需要身份验证。通过 `X-API-Key` 或 `Authorization: Bearer` 头部传递 API 密钥。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| id | string | 调度器的唯一标识符 |

## 请求体

| 字段 | 类型 | 描述 |
|------|------|------|
| id | string | 客户端提供的任务 ID（可选） |
| future | string | ISO 8601 时间戳，指定任务可用的时间 |
| delay | integer | 从现在开始的延迟秒数（future 的替代方案） |
| payload | object | 传递给工作进程的任意 JSON 负载 |
| labels | object | 可选的键值对，用于过滤 |

```json
{
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
}
```

## 请求示例

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

## 响应示例

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "task-x1y2z3",
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

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | invalid_request | 请求体缺失或无效 |
| 401 | unauthorized | 缺少或无效的 API 密钥 |
| 404 | not_found | 未找到调度器 |
| 500 | internal_error | 服务器内部错误 |
