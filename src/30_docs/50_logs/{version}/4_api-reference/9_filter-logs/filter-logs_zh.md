# 过滤 Logs

流式传输匹配 `regex` 查询参数的日志数据。

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?regex=error" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

`regex` 接受 Go 正则表达式。
