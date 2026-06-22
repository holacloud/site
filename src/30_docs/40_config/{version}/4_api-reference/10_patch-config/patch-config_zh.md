# 修改配置

部分更新配置。仅修改请求体中提供的字段。此端点为公开访问（管理 API）。

## 身份验证

无需身份验证。此为公开端点。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `id` | string | 配置 ID（例如 `cfg_abc123`） |

## 请求体

```json
{
  "data": {
    "features": {
      "new-checkout": true
    }
  }
}
```

| 字段 | 类型 | 必填 | 描述 |
|-------|------|----------|-------------|
| `project` | string | 否 | 新的项目名称 |
| `environment` | string | 否 | 新的环境名称 |
| `data` | object | 否 | 要合并的部分配置数据 |

## 请求

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{
    "data": {
      "features": {
        "new-checkout": true
      }
    }
  }'
```

## 响应

```http
HTTP/1.1 200 OK
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
      "new-checkout": true
    }
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "updatedAt": "2026-06-21T11:00:00Z"
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 400 | Bad Request | 请求体无效 |
| 404 | Not Found | 指定的配置不存在 |
| 409 | Conflict | 更新与现有配置冲突 |
| 500 | Internal Server Error | 发生意外错误 |
