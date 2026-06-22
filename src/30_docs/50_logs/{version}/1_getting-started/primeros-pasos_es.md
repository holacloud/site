# Primeros Pasos

Crea un logger con auth Glue y luego escribe bytes crudos.

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "X-Glue-Authentication: TU_GLUE_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"app"}'
```

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  --data-binary $'hello\nerror line\n'
```

Respuesta:

```json
{ "n": 17 }
```

Filtra con `regex`:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?regex=error" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```
