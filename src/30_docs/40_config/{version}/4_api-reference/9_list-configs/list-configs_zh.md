# GET /v0/configs

列出所有配置。此端点为公开访问（管理 API）。

## 身份验证

无需身份验证。此为公开端点。

## 查询参数

| 参数 | 类型 | 默认值 | 描述 |
|-----------|------|---------|-------------|
| `limit` | integer | 100 | 返回的最大配置数量 |
| `offset` | integer | 0 | 跳过的配置数量 |
| `project` | string | "" | 按项目名称筛选配置 |

## 请求

```bash
curl "https://api.hola.cloud/v0/configs?limit=10&offset=0"
```

## 响应

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "configs": [
    {
      "id": "cfg_abc123",
      "project": "my-app",
      "environment": "production",
      "data": {
        "database": {
          "host": "db.example.com",
          "port": 5432
        }
      },
      "createdAt": "2026-06-21T10:00:00Z",
      "updatedAt": "2026-06-21T10:00:00Z"
    }
  ],
  "total": 1
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 500 | Internal Server Error | 发生意外错误 |
