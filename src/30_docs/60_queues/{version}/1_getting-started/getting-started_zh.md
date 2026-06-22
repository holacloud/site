# GettingStarted

本指南将带您完成创建队列、写入消息、通过长轮询读取消息、检查队列统计信息以及清理资源的过程。

## 前提条件

- 一个 HolaCloud 账户。
- 在您的机器上安装 [curl](https://curl.se/)。

## 第一步：创建队列

通过发送带有名称的POST请求来创建新队列：

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

保存返回的 `id` — 您将在后续请求中使用它。

## 第二步：写入消息

Tailon接受每行一个JSON对象的流式数据（NDJSON）。写入三条消息：

```bash
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/x-ndjson" \
  -d '{"type": "greeting", "text": "Hello"}
{"type": "greeting", "text": "你好"}
{"type": "command", "action": "ping"}'
```

预期响应：

```json
{
  "written": 3
}
```

## 第三步：读取消息

使用长轮询GET请求检索消息。`Limit`头部控制返回的消息数量：

```bash
curl -X GET "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Limit: 2"
```

预期响应：

```json
[
  {
    "id": "m1",
    "type": "greeting",
    "text": "Hello",
    "timestamp": "2025-06-21T12:00:01Z"
  },
  {
    "id": "m2",
    "type": "greeting",
    "text": "你好",
    "timestamp": "2025-06-21T12:00:01Z"
  }
]
```

消息按FIFO顺序返回。重复请求以检索剩余消息。

## 第四步：检查队列统计信息

查看队列详情，包括当前长度和读写总数：

```bash
curl "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

预期响应：

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 1,
  "reads": 2,
  "writes": 3
}
```

## 第五步：删除队列

完成后，删除队列以释放资源：

```bash
curl -X DELETE "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

预期响应：HTTP `204 No Content`。

## 后续步骤

- 了解如何管理多个队列和监控活跃客户端，请参阅[管理队列](../2_managing-queues/managing-queues_zh.md)。
- 探索高级的生产和消费模式，请参阅[生产与消费消息](../3_producing-consuming/producing-consuming_zh.md)。
