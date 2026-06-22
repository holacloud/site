# Filtrar Logs

Transmite datos de log que coinciden con el query parameter `regex`.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?regex=error" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

`regex` acepta una expresión regular de Go.
