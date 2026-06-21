# POST /schedulers/{id}/tasks/reserve

预留一个可用的任务进行处理。任务进入由 `worktime` 定义的租约期。

## 身份验证

需要身份验证。通过 `X-API-Key` 或 `Authorization: Bearer` 头部传递 API 密钥。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| id | string | 调度器的唯一标识符 |

## 请求体

| 字段 | 类型 | 描述 |
|------|------|------|
| worktime | integer | 租约持续时间（秒） |

```json
{
  "worktime": 30
}
```

## 请求示例

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 30
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
  "payload": {
    "type": "send_email",
    "to": "user@example.com",
    "template": "welcome"
  },
  "lease_expires_at": "2025-06-21T12:02:01Z"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 401 | unauthorized | 缺少或无效的 API 密钥 |
| 404 | not_found | 未找到调度器 |
| 409 | conflict | 没有可用于预留的任务 |
| 500 | internal_error | 服务器内部错误 |
