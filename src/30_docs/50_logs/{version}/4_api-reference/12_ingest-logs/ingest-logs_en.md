# Ingest Logs

`POST /v1/loggers/{id}/ingest` ingests the raw request body bytes.

Logger access requires either a logger owner via `X-Glue-Authentication` or a logger API key with `Api-Key` and `Api-Secret`.

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  --data-binary $'line one\nline two\n'
```

```json
{ "n": 18 }
```
