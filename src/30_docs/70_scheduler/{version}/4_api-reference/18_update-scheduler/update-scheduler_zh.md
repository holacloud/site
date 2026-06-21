# PUT /schedulers/{id}

替换现有调度器的配置。

## 身份验证

需要身份验证。通过 `X-API-Key` 或 `Authorization: Bearer` 头部传递 API 密钥。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| id | string | 调度器的唯一标识符 |

## 请求体

```json
{
  "display_name": "renamed-scheduler"
}
```

## 请求示例

```bash
curl -X PUT "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "renamed-scheduler"
  }'
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "renamed-scheduler",
  "task_count": 5,
  "status": "active",
  "created_at": "2025-06-20T10:00:00Z",
  "updated_at": "2025-06-21T09:00:00Z"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | invalid_request | 请求体无效 |
| 401 | unauthorized | 缺少或无效的 API 密钥 |
| 404 | not_found | 未找到调度器 |
| 500 | internal_error | 服务器内部错误 |
