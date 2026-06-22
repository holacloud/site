# Tailon API 参考

基础 URL：`https://api.hola.cloud`

Tailon 是一个简单的内存队列服务。

| 方法 | 路径 | 描述 |
|--------|------|-------------|
| POST | `/v1/queues` | 创建队列；`201` 且响应体为空 |
| GET | `/v1/queues` | 以 `[]string` 列出队列 ID |
| GET | `/v1/queues/{queue_id}` | 返回 `name`、`len`、`reads`、`writes` |
| POST | `/v1/queues/{queue_id}:write` | 写入消息 |
| GET | `/v1/queues/{queue_id}:read` | 以 NDJSON 读取消息 |

delete handler 未实现。
