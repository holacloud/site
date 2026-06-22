# 管理队列

Tailon 提供内存队列。

使用 `POST /v1/queues` 创建，使用 `GET /v1/queues` 以 `[]string` 列出队列 ID，使用 `GET /v1/queues/{queue_id}` 获取 `name`、`len`、`reads` 和 `writes`。

delete handler 未实现。
