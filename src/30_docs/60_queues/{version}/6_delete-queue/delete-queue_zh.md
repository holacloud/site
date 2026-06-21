
# DELETE /v1/queues/{id}

永久删除队列及其所有待处理消息。

## 认证

公开 — 无需认证。

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 队列的唯一标识符 |

## 请求

```bash
curl -X DELETE "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
DELETE /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

## 响应

HTTP `204 无内容`。

## 错误代码

| 代码 | 描述 |
|------|------|
| 404 | 队列未找到 |
