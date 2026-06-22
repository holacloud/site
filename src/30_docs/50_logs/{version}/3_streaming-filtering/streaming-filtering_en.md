# Streaming and Filtering

Use `events` to stream logger output. Add `follow` as a presence flag to keep following new data.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?follow" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

Use `regex` to filter streamed data.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?regex=error|critical" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```
