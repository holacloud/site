# 入门

使用 Glue 认证创建 logger，然后写入原始字节。

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "X-Glue-Authentication: YOUR_GLUE_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"app"}'
```

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  --data-binary $'hello\nerror line\n'
```

响应：

```json
{ "n": 17 }
```

使用 `regex` 过滤：

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?regex=error" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```
