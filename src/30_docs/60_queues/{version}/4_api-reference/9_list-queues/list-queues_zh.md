
# List Queues

列出与您账户关联的所有队列。

## 认证

公开 — 无需认证。

## 请求

```bash
curl "https://api.hola.cloud/v1/queues"
```

```http
GET /v1/queues HTTP/1.1
Host: api.hola.cloud
```

## 响应

```json
[
  {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "my-queue",
    "length": 5,
    "reads": 10,
    "writes": 15
  },
  {
    "id": "b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "name": "events-queue",
    "length": 0,
    "reads": 42,
    "writes": 42
  }
]
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 500 | 服务器内部错误 |
