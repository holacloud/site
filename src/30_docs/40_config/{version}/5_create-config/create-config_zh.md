# POST /v0/configs

创建新的配置。此端点为公开访问（管理 API）。

## 身份验证

无需身份验证。此为公开端点。

## 请求体

```json
{
  "project": "my-app",
  "environment": "production",
  "data": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "features": {
      "new-checkout": false
    }
  }
}
```

| 字段 | 类型 | 必填 | 描述 |
|-------|------|----------|-------------|
| `project` | string | 是 | 项目名称标识符 |
| `environment` | string | 是 | 环境名称（例如 production、staging、development） |
| `data` | object | 是 | 配置数据，JSON 对象 |

## 请求

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{
    "project": "my-app",
    "environment": "production",
    "data": {
      "database": {
        "host": "db.example.com",
        "port": 5432
      },
      "features": {
        "new-checkout": false
      }
    }
  }'
```

## 响应

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "data": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "features": {
      "new-checkout": false
    }
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "updatedAt": "2026-06-21T10:00:00Z"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 400 | Bad Request | 缺少必填字段或 JSON 无效 |
| 409 | Conflict | 相同项目和环境的配置已存在 |
| 500 | Internal Server Error | 发生意外错误 |
