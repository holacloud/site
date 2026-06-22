# Tailon 队列

Tailon 是一个简单的内存队列服务。队列位于进程内存中，并提供用于创建、列出、查看、写入和读取的小型 HTTP API。

## API 概览

| 方法 | 路径 | 描述 |
|--------|------|-------------|
| POST | `/v1/queues` | 创建队列；返回 `201` 且响应体为空 |
| GET | `/v1/queues` | 以 `[]string` 列出队列 ID |
| GET | `/v1/queues/{queue_id}` | 返回 `name`、`len`、`reads` 和 `writes` |
| POST | `/v1/queues/{queue_id}:write` | 写入消息 |
| GET | `/v1/queues/{queue_id}:read` | 以 NDJSON 读取消息 |

delete handler 未实现。
