# GettingStarted

本指南将带您了解如何使用 Config API 检索和更新配置。开始之前，您需要拥有具有适当权限的 API 密钥。

## 第一步：获取配置

使用 `GET /v1/config` 端点获取当前配置：

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret"
```

预期响应：

```json
{
  "project": "my-app",
  "environment": "production",
  "services": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "cache": {
      "host": "redis.example.com",
      "port": 6379
    }
  },
  "features": {
    "new-checkout": false
  }
}
```

## 第二步：更新特定值

使用 `PATCH /v1/config` 发送部分 JSON 负载，仅更新您需要的字段。例如，启用新的结账功能：

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "features": {
      "new-checkout": true
    }
  }'
```

预期响应：

```json
{
  "project": "my-app",
  "environment": "production",
  "services": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "cache": {
      "host": "redis.example.com",
      "port": 6379
    }
  },
  "features": {
    "new-checkout": true
  }
}
```

## 第三步：验证更新

再次调用 `GET /v1/config` 确认更改已生效：

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret"
```

响应现在应显示 `"new-checkout": true`。

## 总结

只需几个步骤，您就完成了配置的检索、部分更新和验证。Config API 让您无需重新部署应用即可轻松管理运行时设置。
