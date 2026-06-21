# POST /schedulers/{id}/tasks/{task}/reschedule

使用新的延迟或未来时间重新调度任务。

## 身份验证

需要身份验证。通过 `X-API-Key` 或 `Authorization: Bearer` 头部传递 API 密钥。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| id | string | 调度器的唯一标识符 |
| task | string | 任务的唯一标识符 |

## 请求体

| 字段 | 类型 | 描述 |
|------|------|------|
| future | string | ISO 8601 时间戳，指定任务可用的时间 |
| delay | integer | 从现在开始的延迟秒数（future 的替代方案） |

```json
{
  "delay": 120
}
```

## 请求示例

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-x1y2z3/reschedule" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "delay": 120
  }'
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "task-x1y2z3",
  "state": "pending",
  "available_at": "2025-06-21T12:03:01Z"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | invalid_request | 缺少或无效的 future/delay |
| 401 | unauthorized | 缺少或无效的 API 密钥 |
| 404 | not_found | 未找到调度器或任务 |
| 500 | internal_error | 服务器内部错误 |
