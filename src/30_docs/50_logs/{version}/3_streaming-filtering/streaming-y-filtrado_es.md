# Streaming y Filtrado

Usa `events` para transmitir la salida del logger. Agrega `follow` como flag de presencia para seguir datos nuevos.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?follow" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

Usa `regex` para filtrar datos transmitidos.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?regex=error|critical" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```
