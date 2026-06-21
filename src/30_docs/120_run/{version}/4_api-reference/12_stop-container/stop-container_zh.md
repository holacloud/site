# POST /api/console/stop

停止正在运行的容器。

## 描述

通过容器 ID 优雅地停止容器。在强制终止之前会给容器一个宽限期。

## 身份验证

无。此端点为公开端点。

## 请求体

```json
{
  "container_id": "abc123def456"
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/api/console/stop" \
  -H "Content-Type: application/json" \
  -d '{"container_id": "abc123def456"}'
```

## 响应

```json
{
  "container_id": "abc123def456",
  "status": "stopped",
  "stopped_at": "2026-06-21T11:00:00Z"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 容器停止成功 |
| 400 | 请求体无效 |
| 404 | 容器未找到 |
