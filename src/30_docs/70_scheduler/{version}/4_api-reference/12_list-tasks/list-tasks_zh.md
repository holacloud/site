# GET /schedulers/{id}/tasks

列出调度器的任务，已计划和运行中的任务分别分页。

## 身份验证

此端点为公开，无需身份验证。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| id | string | 调度器的唯一标识符 |

## 查询参数

| 参数 | 类型 | 描述 |
|------|------|------|
| scheduled_page | integer | 已计划任务的页码（默认：1） |
| scheduled_per_page | integer | 已计划任务每页条数（默认：20） |
| inflight_page | integer | 运行中任务的页码（默认：1） |
| inflight_per_page | integer | 运行中任务每页条数（默认：20） |
| search | string | 按负载内容搜索任务 |
| label | string | 按标签过滤（格式：key:value） |

## 请求示例

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?label=priority:high"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "scheduled": {
    "tasks": [
      {
        "id": "task-x1y2z3",
        "state": "pending",
        "available_at": "2025-06-21T12:01:01Z",
        "labels": { "project": "onboarding", "priority": "high" }
      }
    ],
    "total": 1,
    "page": 1,
    "per_page": 20
  },
  "inflight": {
    "tasks": [],
    "total": 0,
    "page": 1,
    "per_page": 20
  }
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 404 | not_found | 未找到调度器 |
| 500 | internal_error | 服务器内部错误 |
