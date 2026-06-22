# Enviando Logs

La ingesta directa escribe el cuerpo crudo del request en el logger.

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  --data-binary $'line one\nline two\n'
```

```json
{ "n": 18 }
```

La ingesta enmarcada usa datos logframe en `POST /v1/ingest/events`. La metadata de cada frame debe incluir `project_id`; la respuesta informa `events` aceptados y `bytes` escritos.
