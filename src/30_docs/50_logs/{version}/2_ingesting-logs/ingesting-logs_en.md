# Ingesting Logs

Direct ingest writes the raw request body to the logger.

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  --data-binary $'line one\nline two\n'
```

```json
{ "n": 18 }
```

Framed ingest uses logframe data on `POST /v1/ingest/events`. Each frame metadata must include `project_id`; the response reports accepted `events` and written `bytes`.
