# 写入消息

使用 `POST /v1/queues/{queue_id}:write` 写入消息。

```bash
curl -X POST "https://api.hola.cloud/v1/queues/queue_1:write" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary $'{"message":"hello"}\n'
```
