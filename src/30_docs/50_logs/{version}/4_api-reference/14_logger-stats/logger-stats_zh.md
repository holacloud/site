
# Logger Stats

获取日志记录器的统计信息，包括接收的事件总数和存储使用量。

## 认证

需要数据凭据：

- `X-Instantlogs-Event-Secret: <secret>` 或 `Authorization: Bearer <secret>`

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 日志记录器的唯一标识符 |

## 请求

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/stats" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

```http
GET /v1/loggers/logger_xyz789/stats HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```

## 响应

```json
{
  "total_events": 15234,
  "storage_bytes": 4194304,
  "events_last_hour": 234,
  "events_last_day": 5601
}
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 401 | 缺少或无效的事件密钥 |
| 404 | 日志记录器未找到 |
