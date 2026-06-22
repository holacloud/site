# InstantLogs

InstantLogs 提供用于原始字节摄入以及流式传输/过滤的 HTTP logger。

## 认证

Logger 管理使用 Glue 认证 `X-Glue-Authentication`。Logger 专用端点可以由该 logger 的 Glue owner 访问，也可以使用 logger API key 的 `Api-Key` 和 `Api-Secret` 访问。

## API 概览

| 方法 | 路径 | 描述 |
|--------|------|-------------|
| GET | `/v1/loggers` | 列出 Glue 用户可见的 logger |
| POST | `/v1/loggers` | 为 Glue 用户创建 logger |
| GET | `/v1/loggers/{id}` | 获取 logger 详情，包括 owners 和 API keys |
| DELETE | `/v1/loggers/{id}` | 删除 logger |
| POST | `/v1/loggers/{id}/ingest` | 摄入请求原始字节，返回 `{ "n": bytes }` |
| GET | `/v1/loggers/{id}/filter?regex=...` | 流式传输匹配 regex 的条目 |
| GET | `/v1/loggers/{id}/events` | 流式传输事件；`follow` 由 flag 是否存在启用 |
| GET | `/v1/loggers/{id}/stats` | 获取运行时统计 |
| POST | `/v1/loggers/{id}/apiKeys` | 创建 logger API key |
| DELETE | `/v1/loggers/{id}/apiKeys/{apiKey}` | 删除 logger API key |
| POST | `/v1/ingest/events` | 摄入带 `project_id` 的 logframe 事件，返回 `events` 和 `bytes` |
