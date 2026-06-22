# 入门

Tailon 队列是简单的内存队列。

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{"name":"my-queue"}'
```

创建返回 `201`，响应体为空。

```bash
curl -X POST "https://api.hola.cloud/v1/queues/queue_1:write" \
  --data-binary $'{"message":"hello"}\n'
```

```bash
curl "https://api.hola.cloud/v1/queues/queue_1:read"
```

读取返回 NDJSON。
