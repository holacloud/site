# Stream Events

Stream logger events.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?follow" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

`follow` is enabled when the query flag is present. Use `regex` to filter by regular expression.
