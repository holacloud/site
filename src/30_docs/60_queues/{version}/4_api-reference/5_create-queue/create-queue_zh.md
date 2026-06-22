# 创建队列

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{"name":"my-queue"}'
```

响应：`201 Created`，响应体为空。
