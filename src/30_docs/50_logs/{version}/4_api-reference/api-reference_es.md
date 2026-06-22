# Referencia API de InstantLogs

URL base: `https://api.hola.cloud`

## Autenticación

Los endpoints de gestión usan `X-Glue-Authentication`. Los endpoints de logger aceptan un owner Glue del logger o una API key del logger con `Api-Key` y `Api-Secret`.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/v1/loggers` | Listar loggers propios |
| POST | `/v1/loggers` | Crear logger |
| GET | `/v1/loggers/{id}` | Obtener logger con owners/API keys |
| DELETE | `/v1/loggers/{id}` | Eliminar logger |
| POST | `/v1/loggers/{id}/ingest` | Ingerir bytes crudos, devuelve `{ "n": number }` |
| GET | `/v1/loggers/{id}/filter?regex=...` | Transmitir logs filtrados |
| GET | `/v1/loggers/{id}/events` | Transmitir eventos; `follow` es un flag de presencia |
| GET | `/v1/loggers/{id}/stats` | Estadísticas de ejecución |
| POST | `/v1/loggers/{id}/apiKeys` | Crear API key del logger |
| DELETE | `/v1/loggers/{id}/apiKeys/{apiKey}` | Eliminar API key del logger |
| POST | `/v1/loggers/{id}/owners` | Agregar owner del logger |
| DELETE | `/v1/loggers/{id}/owners/{ownerId}` | Eliminar owner del logger |
| POST | `/v1/ingest/events` | Ingerir eventos logframe con `project_id`, devuelve `events` y `bytes` |
