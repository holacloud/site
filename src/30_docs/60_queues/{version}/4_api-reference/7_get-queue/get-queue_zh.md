
# Get Queue

根据 ID 获取特定队列的详细信息。

## 认证

公开 — 无需认证。

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 队列的唯一标识符 |

## 请求

```bash
curl "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

## 响应

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 5,
  "reads": 10,
  "writes": 15
}
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 404 | 队列未找到 |
