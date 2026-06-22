# ManagingConfigs

管理 API（`/v0/configs`）提供了完整的 CRUD 操作来管理配置。这些端点可以公开访问，无需 API 密钥。

## 列出所有配置

获取所有已存储配置的列表：

```bash
curl "https://api.hola.cloud/v0/configs"
```

预期响应：

```json
[
  {
    "id": "cfg_abc123",
    "project": "my-app",
    "environment": "production",
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-06-20T14:22:00Z"
  },
  {
    "id": "cfg_def456",
    "project": "my-app",
    "environment": "staging",
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-06-19T09:00:00Z"
  }
]
```

## 创建配置

创建新的配置条目：

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{
    "project": "my-app",
    "environment": "production",
    "config": {
      "database": {
        "host": "db.example.com",
        "port": 5432
      }
    }
  }'
```

预期响应：

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "config": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    }
  },
  "created_at": "2025-06-20T14:22:00Z",
  "updated_at": "2025-06-20T14:22:00Z"
}
```

## 获取配置

根据 ID 获取单个配置：

```bash
curl "https://api.hola.cloud/v0/configs/cfg_abc123"
```

预期响应：

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "config": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    }
  },
  "created_at": "2025-01-15T10:30:00Z",
  "updated_at": "2025-06-20T14:22:00Z"
}
```

## 部分更新配置

对现有配置进行部分更新：

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{
    "config": {
      "database": {
        "host": "db-new.example.com"
      }
    }
  }'
```

预期响应：

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "config": {
    "database": {
      "host": "db-new.example.com",
      "port": 5432
    }
  },
  "created_at": "2025-01-15T10:30:00Z",
  "updated_at": "2025-06-20T15:00:00Z"
}
```

## 删除配置

永久删除配置：

```bash
curl -X DELETE "https://api.hola.cloud/v0/configs/cfg_abc123"
```

预期响应：`HTTP 204 No Content`

## HTTP 请求参考

```http
GET /v0/configs HTTP/1.1
Host: api.hola.cloud

POST /v0/configs HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "project": "my-app",
  "environment": "production",
  "config": { ... }
}

GET /v0/configs/{id} HTTP/1.1
Host: api.hola.cloud

PATCH /v0/configs/{id} HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "config": { ... }
}

DELETE /v0/configs/{id} HTTP/1.1
Host: api.hola.cloud
```
