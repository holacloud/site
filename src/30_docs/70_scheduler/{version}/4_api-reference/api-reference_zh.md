# 调度器 API 参考

基础 URL：`https://api.hola.cloud`

| 方法 | 路径 | 描述 | 认证 |
|------|------|------|------|
| GET | /schedulers | 列出调度器 | 公开 |
| POST | /schedulers | 创建调度器 | 私有 |
| GET | /schedulers/{id} | 获取调度器 | 公开 |
| PUT | /schedulers/{id} | 更新调度器 | 私有 |
| PATCH | /schedulers/{id} | 部分更新调度器 | 私有 |
| DELETE | /schedulers/{id} | 删除调度器 | 私有 |
| GET | /schedulers/{id}/health | 健康检查 | 公开 |
| POST | /schedulers/{id}/tasks | 入队任务 | 私有 |
| GET | /schedulers/{id}/tasks | 列出任务 | 公开 |
| GET | /schedulers/{id}/tasks/stream | SSE 任务流 | 公开 |
| POST | /schedulers/{id}/tasks/reserve | 预留任务 | 私有 |
| DELETE | /schedulers/{id}/tasks/{task} | 确认任务 | 私有 |
| POST | /schedulers/{id}/tasks/{task}/extend | 延长租约 | 私有 |
| POST | /schedulers/{id}/tasks/{task}/reschedule | 重新调度任务 | 私有 |

## 身份验证

私有端点需要 API 密钥。请使用以下任一方式传递：

- `X-API-Key` 头部：`X-API-Key: YOUR_API_KEY`
- `Authorization` 头部：`Authorization: Bearer YOUR_API_KEY`

只读端点（GET）为公开，无需身份验证。

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | validation_error | 请求体或参数无效 |
| 401 | unauthorized | 缺少或无效的 API 密钥 |
| 404 | not_found | 未找到调度器或任务 |
| 409 | conflict | 任务已被预留或租约冲突 |
| 500 | internal_error | 服务器内部错误 |
