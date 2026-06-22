# Getting Started

Create a logger with Glue auth, then write raw bytes to it.

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "X-Glue-Authentication: YOUR_GLUE_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"app"}'
```

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  --data-binary $'hello\nerror line\n'
```

Response:

```json
{ "n": 17 }
```

Filter with `regex`:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?regex=error" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```
