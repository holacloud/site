# Transmitir Eventos

Transmite eventos del logger.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?follow" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

`follow` se activa cuando el flag está presente. Usa `regex` para filtrar por expresión regular.
