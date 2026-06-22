# Ingerir Logs

`POST /v1/loggers/{id}/ingest` ingiere los bytes crudos del cuerpo del request.

El acceso al logger requiere un owner vía `X-Glue-Authentication` o una API key del logger con `Api-Key` y `Api-Secret`.

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  --data-binary $'line one\nline two\n'
```

```json
{ "n": 18 }
```
