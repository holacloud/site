# 保存存储卷

保存容器的卷配置。

## 描述

更新指定容器的持久卷配置。卷由 HolaCloud 的分布式存储支持，并在容器重启后仍然存在。

## 身份验证

无。此端点为公开端点。

## 请求体

```json
{
  "container_id": "abc123def456",
  "volumes": [
    {
      "container_path": "/data",
      "size_gb": 20,
      "mount_path": "/mnt/data"
    },
    {
      "container_path": "/config",
      "size_gb": 5,
      "mount_path": "/mnt/config"
    }
  ]
}
```

## 示例

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Content-Type: application/json" \
  -d '{
    "container_id": "abc123def456",
    "volumes": [
      {"container_path": "/data", "size_gb": 20}
    ]
  }'
```

## 响应

```json
{
  "container_id": "abc123def456",
  "volumes": [
    {
      "volume_id": "vol-001",
      "container_path": "/data",
      "size_gb": 20,
      "mount_path": "/mnt/data",
      "status": "active"
    }
  ],
  "updated": true
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 卷配置保存成功 |
| 400 | 请求体无效 |
| 404 | 容器未找到 |
