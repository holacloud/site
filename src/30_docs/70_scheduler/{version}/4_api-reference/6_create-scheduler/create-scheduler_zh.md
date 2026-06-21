# POST /schedulers

创建一个新的调度器。

## 身份验证

需要身份验证。通过 `X-API-Key` 或 `Authorization: Bearer` 头部传递 API 密钥。

## 请求体

```json
{
  "display_name": "my-scheduler"
}
```

## 请求示例

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "my-scheduler"
  }'
```

## 响应示例

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "my-scheduler",
  "task_count": 0,
  "status": "active",
  "created_at": "2025-06-21T10:00:00Z"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | invalid_request | 缺少或无效的 display_name |
| 401 | unauthorized | 缺少或无效的 API 密钥 |
| 500 | internal_error | 服务器内部错误 |
