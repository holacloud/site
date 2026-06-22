# InstantLogs

InstantLogs proporciona loggers HTTP para ingesta de bytes crudos y streaming/filtrado.

## Autenticación

La gestión de loggers usa autenticación Glue con `X-Glue-Authentication`. Los endpoints específicos de un logger pueden usarse con un propietario Glue del logger o con una API key del logger usando `Api-Key` y `Api-Secret`.

## Resumen de API

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/v1/loggers` | Listar loggers visibles para el usuario Glue |
| POST | `/v1/loggers` | Crear un logger para el usuario Glue |
| GET | `/v1/loggers/{id}` | Obtener detalles del logger, incluidos owners y API keys |
| DELETE | `/v1/loggers/{id}` | Eliminar un logger |
| POST | `/v1/loggers/{id}/ingest` | Ingerir bytes crudos del request, devuelve `{ "n": bytes }` |
| GET | `/v1/loggers/{id}/filter?regex=...` | Transmitir entradas que coinciden con una regex |
| GET | `/v1/loggers/{id}/events` | Transmitir eventos; `follow` se activa por presencia del flag |
| GET | `/v1/loggers/{id}/stats` | Obtener estadísticas de ejecución |
| POST | `/v1/loggers/{id}/apiKeys` | Crear una API key del logger |
| DELETE | `/v1/loggers/{id}/apiKeys/{apiKey}` | Eliminar una API key del logger |
| POST | `/v1/ingest/events` | Ingerir eventos logframe con `project_id`, devuelve `events` y `bytes` |
