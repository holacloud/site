# InstantLogs API 参考

基础 URL：`https://api.hola.cloud`

## 认证

管理端点使用 `X-Glue-Authentication`。Logger 端点接受该 logger 的 Glue owner，或使用 `Api-Key` 和 `Api-Secret` 的 logger API key。

## 端点

| 方法 | 路径 | 描述 |
|--------|------|-------------|
| GET | `/v1/loggers` | 列出拥有的 logger |
| POST | `/v1/loggers` | 创建 logger |
| GET | `/v1/loggers/{id}` | 获取包含 owners/API keys 的 logger |
| DELETE | `/v1/loggers/{id}` | 删除 logger |
| POST | `/v1/loggers/{id}/ingest` | 摄入原始字节，返回 `{ "n": number }` |
| GET | `/v1/loggers/{id}/filter?regex=...` | 流式传输过滤后的日志 |
| GET | `/v1/loggers/{id}/events` | 流式传输事件；`follow` 是存在即启用的 flag |
| GET | `/v1/loggers/{id}/stats` | 运行时统计 |
| POST | `/v1/loggers/{id}/apiKeys` | 创建 logger API key |
| DELETE | `/v1/loggers/{id}/apiKeys/{apiKey}` | 删除 logger API key |
| POST | `/v1/loggers/{id}/owners` | 添加 logger owner |
| DELETE | `/v1/loggers/{id}/owners/{ownerId}` | 删除 logger owner |
| POST | `/v1/ingest/events` | 摄入带 `project_id` 的 logframe 事件，返回 `events` 和 `bytes` |
