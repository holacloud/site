# 读取消息

使用 `GET /v1/queues/{queue_id}:read` 读取消息。响应为 NDJSON。

```bash
curl "https://api.hola.cloud/v1/queues/queue_1:read"
```

```json
{"message":"hello"}
```
