# 任务流式传输

使用服务器发送事件（SSE）实时流式传输任务事件。

## 身份验证

此端点为公开，无需身份验证。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| id | string | 调度器的唯一标识符 |

## 请求示例

```bash
curl -N "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/stream"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: text/event-stream
Cache-Control: no-cache
Connection: keep-alive
```

```
event: snapshot
data: {"scheduler_id":"sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890","generated_at":"2025-06-21T12:00:01Z","scheduled":[{"id":"task-x1y2z3","future":"2025-06-21T12:01:01Z","labels":["priority:high"]}],"inflight":[],"scheduled_meta":{"page":1,"per_page":25,"total":1,"total_pages":1},"inflight_meta":{"page":1,"per_page":25,"total":0,"total_pages":0},"health":{"status":"ok","ready":true,"scheduled":1,"inflight":0,"scheduler_id":"sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 404 | not_found | 未找到调度器 |
| 500 | internal_error | 服务器内部错误 |
