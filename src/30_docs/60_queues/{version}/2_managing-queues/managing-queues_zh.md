# 管理队列

本文档涵盖队列的生命周期操作：创建（含可选配置）、列表、查看详情、客户端监控和删除。

## 创建队列

通过发送POST请求创建队列：

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-queue"
  }'
```

预期响应：

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 0,
  "reads": 0,
  "writes": 0
}
```

对应的HTTP请求如下：

```http
POST /v1/queues HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "name": "my-queue"
}
```

## 列出所有队列

检索您账户中的所有队列：

```bash
curl "https://api.hola.cloud/v1/queues"
```

预期响应：

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

## 获取队列详情

通过ID查看单个队列的当前长度、总读取数和总写入数：

```bash
curl "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

预期响应：

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 5,
  "reads": 10,
  "writes": 15
}
```

## 监控活跃客户端

列出所有当前连接的流式客户端：

```bash
curl "https://api.hola.cloud/v1/clients"
```

```http
GET /v1/clients HTTP/1.1
Host: api.hola.cloud
```

预期响应：

```json
[
  {
    "id": "client-001",
    "queue_id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "connected_at": "2025-06-21T12:00:00Z",
    "user_agent": "curl/7.88.1"
  }
]
```

## 删除队列

永久删除队列及其所有待处理消息：

```bash
curl -X DELETE "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
DELETE /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

预期响应：HTTP `204 No Content`。

删除后，该队列ID将不再有效。正在进行的读取器将在下一次请求时收到错误。
