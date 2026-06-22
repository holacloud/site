# 保存环境变量

保存某个仓库的环境变量。

```json
{
  "repository": "my-project/my-app",
  "env": [
    {"key": "LOG_LEVEL", "desired_value": "info"},
    {"key": "DATABASE_URL", "desired_value": "postgres://user:pass@host:5432/db"}
  ]
}
```

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "env": [{"key": "LOG_LEVEL", "desired_value": "info"}]}'
```
