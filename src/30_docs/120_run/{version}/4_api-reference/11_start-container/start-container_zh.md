# 启动容器

从指定镜像启动新容器。

## 描述

从给定镜像创建并启动容器，可设置环境变量、卷和端口映射。

## 身份验证

无。此端点为公开端点。

## 请求体

```json
{
  "image": "my-project/my-app:latest",
  "env": {
    "LOG_LEVEL": "debug",
    "DATABASE_URL": "postgres://user:pass@host:5432/db"
  },
  "volumes": [
    {
      "container_path": "/data",
      "size_gb": 10
    }
  ],
  "ports": {
    "80/tcp": 8080
  }
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "my-project/my-app:latest",
    "env": {"LOG_LEVEL": "debug"},
    "ports": {"80/tcp": 8080}
  }'
```

## 响应

```json
{
  "container_id": "abc123def456",
  "image": "my-project/my-app:latest",
  "status": "running",
  "ports": {"80/tcp": 8080},
  "started_at": "2026-06-21T10:00:00Z"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 容器启动成功 |
| 400 | 请求体无效 |
| 404 | 镜像未找到 |
