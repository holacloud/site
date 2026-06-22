# 获取调度器

返回特定调度器的详细信息。

## 身份验证

此端点为公开，无需身份验证。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| id | string | 调度器的唯一标识符 |

## 请求示例

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "my-scheduler",
  "ready": true,
  "scheduled": 5,
  "inflight": 0,
  "created_at": "2025-06-20T10:00:00Z",
  "updated_at": "2025-06-21T08:30:00Z"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 404 | not_found | 未找到调度器 |
| 500 | internal_error | 服务器内部错误 |
