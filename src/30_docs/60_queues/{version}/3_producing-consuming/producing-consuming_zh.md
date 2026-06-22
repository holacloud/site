# 生产和消费

使用 `POST /v1/queues/{queue_id}:write` 写入，使用 `GET /v1/queues/{queue_id}:read` 读取。

```bash
curl -X POST "https://api.hola.cloud/v1/queues/queue_1:write" \
  --data-binary $'{"message":"hello"}\n'
```

```bash
curl "https://api.hola.cloud/v1/queues/queue_1:read"
```

读取响应为 NDJSON。
