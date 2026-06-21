# DELETE /schedulers/{id}/tasks/{task}

确认并从调度器中移除任务。工作进程成功处理预留的任务后应调用此接口。

## 身份验证

需要身份验证。通过 `X-API-Key` 或 `Authorization: Bearer` 头部传递 API 密钥。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| id | string | 调度器的唯一标识符 |
| task | string | 任务的唯一标识符 |

## 请求示例

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-x1y2z3" \
  -H "X-API-Key: YOUR_API_KEY"
```

## 响应示例

```http
HTTP/1.1 204 No Content
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 401 | unauthorized | 缺少或无效的 API 密钥 |
| 404 | not_found | 未找到调度器或任务 |
| 409 | conflict | 任务当前未被预留 |
| 500 | internal_error | 服务器内部错误 |
