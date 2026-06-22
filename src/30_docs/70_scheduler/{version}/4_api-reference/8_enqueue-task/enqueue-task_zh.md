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
| id | string | 必需的任务 ID |
| future | string | ISO 8601 时间戳，指定任务可用的时间 |
| delay | string | 从现在开始的 Go duration 字符串，例如 `60s` 或 `5m`（future 的替代） |
| payload | object | 传递给工作进程的任意 JSON 负载 |
| labels | string array | 用于过滤的可选标签 |

```json
{
  "payload": {
    "type": "send_email",
    "to": "user@example.com",
    "template": "welcome"
  },
  "delay": "60s",
  "labels": ["project:onboarding", "priority:high"]
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
    "delay": "60s",
    "labels": ["project:onboarding", "priority:high"]
  }'
```

## 响应示例

```http
HTTP/1.1 202 Accepted
```

响应体为空。

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | invalid_json | JSON 无效 |
| 400 | validation_error | 缺少 id、future/delay 无效或 labels 无效 |
| 401 | unauthorized | 缺少或无效的 API 密钥 |
| 409 | task_already_exists | 任务已存在 |
| 500 | internal_error | 服务器内部错误 |
