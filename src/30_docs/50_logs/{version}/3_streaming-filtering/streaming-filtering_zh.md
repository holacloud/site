# 流式传输和过滤

使用 `events` 流式传输 logger 输出。添加 `follow` 作为存在即启用的 flag 来继续跟随新数据。

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?follow" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

使用 `regex` 过滤流式数据。

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?regex=error|critical" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```
