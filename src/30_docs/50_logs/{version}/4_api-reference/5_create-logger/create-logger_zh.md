# 创建 Logger

创建归已认证 Glue 用户所有的 logger。

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "X-Glue-Authentication: YOUR_GLUE_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"app"}'
```

```json
{
  "id": "logger_xyz789",
  "name": "app",
  "owners": ["user_123"],
  "api_keys": [],
  "usage": {}
}
```
