
# Create Queue

创建新队列。

## 认证

公开 — 无需认证。

## 请求体

| 字段 | 类型 | 必填 | 描述 |
|------|------|------|-------------|
| `name` | string | 是 | 队列的名称 |

```json
{
  "name": "my-queue"
}
```

## 请求

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-queue"
  }'
```

```http
POST /v1/queues HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "name": "my-queue"
}
```

## 响应

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 0,
  "reads": 0,
  "writes": 0
}
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | `name` 字段缺失或无效 |
| 409 | 此名称的队列已存在 |
