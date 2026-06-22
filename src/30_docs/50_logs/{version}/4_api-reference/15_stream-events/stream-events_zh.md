# 流式传输事件

流式传输 logger 事件。

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?follow" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

`follow` 在查询 flag 存在时启用。使用 `regex` 按正则表达式过滤。
