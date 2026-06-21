
# InstantLogs API 参考

基础 URL：`https://api.hola.cloud`

## 认证

InstantLogs 根据端点类型使用两种认证模式。

### 管理认证

管理端点（日志记录器 CRUD、API 密钥管理）需要两个请求头：

| 请求头 | 描述 |
|--------|------|
| `Api-Key` | 您的 API 密钥标识符 |
| `Api-Secret` | 您的 API 密钥密码 |

### 数据认证

数据端点（接收、过滤、事件、统计）使用日志记录器密钥进行认证。请按以下方式提供：

| 请求头 | 描述 |
|--------|------|
| `X-Instantlogs-Event-Secret` | 日志记录器密钥值 |
| `Authorization: Bearer <secret>` | 包含日志记录器密钥的 Bearer 令牌 |

## 端点

| 方法 | 路径 | 描述 | 认证 |
|--------|------|------|------|
| GET | `/v1/loggers` | 列出所有日志记录器 | 管理 |
| POST | `/v1/loggers` | 创建新的日志记录器 | 管理 |
| GET | `/v1/loggers/{id}` | 获取日志记录器详情 | 管理 |
| DELETE | `/v1/loggers/{id}` | 删除日志记录器 | 管理 |
| POST | `/v1/loggers/{id}/ingest` | 接收日志条目 | 数据 |
| GET | `/v1/loggers/{id}/filter` | 流式传输和过滤日志 | 数据 |
| GET | `/v1/loggers/{id}/events` | 流式传输事件 | 数据 |
| GET | `/v1/loggers/{id}/stats` | 获取日志记录器统计信息 | 数据 |
| POST | `/v1/loggers/{id}/apiKeys` | 创建 API 密钥 | 管理 |
| DELETE | `/v1/loggers/{id}/apiKeys/{key}` | 删除 API 密钥 | 管理 |
| POST | `/v1/ingest/events` | 接收帧事件 | 数据 |

## 常见错误代码

| 代码 | 描述 |
|------|------|
| 400 | 错误请求 — JSON 格式错误或参数无效 |
| 401 | 未授权 — 缺少或无效的凭据 |
| 403 | 禁止 — 凭据无权访问 |
| 404 | 未找到 — 请求的资源不存在 |
| 409 | 冲突 — 资源已存在 |
| 429 | 请求过多 — 超出速率限制 |
| 500 | 服务器内部错误 |
