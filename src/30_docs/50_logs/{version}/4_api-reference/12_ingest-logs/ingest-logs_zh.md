# 摄入 Logs

`POST /v1/loggers/{id}/ingest` 摄入请求体的原始字节。

Logger 访问需要通过 `X-Glue-Authentication` 的 owner，或使用 `Api-Key` 和 `Api-Secret` 的 logger API key。

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  --data-binary $'line one\nline two\n'
```

```json
{ "n": 18 }
```
