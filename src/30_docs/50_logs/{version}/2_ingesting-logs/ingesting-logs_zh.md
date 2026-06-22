# 摄入 Logs

直接摄入会将请求体原始内容写入 logger。

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  --data-binary $'line one\nline two\n'
```

```json
{ "n": 18 }
```

帧摄入在 `POST /v1/ingest/events` 使用 logframe 数据。每个 frame 的 metadata 必须包含 `project_id`；响应报告接收的 `events` 和写入的 `bytes`。
