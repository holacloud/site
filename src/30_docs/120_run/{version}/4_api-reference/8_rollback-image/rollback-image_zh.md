# POST /api/console/rollback

将容器回滚到之前的镜像版本。

## 描述

将正在运行或已停止容器的当前镜像替换为指定的先前版本并重新启动。

## 身份验证

无。此端点为公开端点。

## 请求体

```json
{
  "container_id": "abc123def456",
  "tag": "v1.0.0"
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/api/console/rollback" \
  -H "Content-Type: application/json" \
  -d '{
    "container_id": "abc123def456",
    "tag": "v1.0.0"
  }'
```

## 响应

```json
{
  "container_id": "abc123def456",
  "previous_image": "my-project/my-app:latest",
  "current_image": "my-project/my-app:v1.0.0",
  "status": "running",
  "rolled_back_at": "2026-06-21T12:00:00Z"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 回滚成功完成 |
| 400 | 请求体无效 |
| 404 | 容器或镜像标签未找到 |
