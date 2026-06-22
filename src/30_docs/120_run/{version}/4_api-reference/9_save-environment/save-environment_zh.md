# 保存环境变量

保存容器的环境变量。

## 描述

更新指定容器的环境变量配置。需要重新启动容器才能使更改生效。

## 身份验证

无。此端点为公开端点。

## 请求体

```json
{
  "container_id": "abc123def456",
  "env": {
    "LOG_LEVEL": "info",
    "DATABASE_URL": "postgres://user:pass@host:5432/db",
    "REDIS_URL": "redis://host:6379"
  }
}
```

## 示例

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Content-Type: application/json" \
  -d '{
    "container_id": "abc123def456",
    "env": {
      "LOG_LEVEL": "info",
      "DATABASE_URL": "postgres://user:pass@host:5432/db"
    }
  }'
```

## 响应

```json
{
  "container_id": "abc123def456",
  "env": {
    "LOG_LEVEL": "info",
    "DATABASE_URL": "postgres://user:pass@host:5432/db",
    "REDIS_URL": "redis://host:6379"
  },
  "updated": true
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 环境变量保存成功 |
| 400 | 请求体无效 |
| 404 | 容器未找到 |
