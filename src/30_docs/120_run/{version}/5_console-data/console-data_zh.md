# GET /api/console

返回控制台仪表板的仓库、镜像和正在运行的容器状态。

## 描述

此端点聚合了 HolaCloud 控制台显示已部署服务当前状态所需的所有容器相关数据。

## 身份验证

无。此端点为公开端点。

## 请求

无需请求体。

## 示例

```bash
curl -X GET "https://api.hola.cloud/api/console"
```

## 响应

```json
{
  "repositories": [
    {
      "name": "my-project/my-app",
      "tags": ["latest", "v1.0.0", "v1.1.0"],
      "pull_count": 142
    }
  ],
  "images": [
    {
      "repository": "my-project/my-app",
      "tag": "latest",
      "digest": "sha256:a1b2c3d4...",
      "size_bytes": 72450000,
      "created": "2026-06-20T10:00:00Z"
    }
  ],
  "containers": [
    {
      "id": "abc123def456",
      "image": "my-project/my-app:latest",
      "status": "running",
      "ports": {"80/tcp": 8080},
      "started_at": "2026-06-21T08:00:00Z",
      "env": {
        "LOG_LEVEL": "info",
        "DATABASE_URL": "postgres://..."
      }
    }
  ]
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 成功返回控制台数据 |
