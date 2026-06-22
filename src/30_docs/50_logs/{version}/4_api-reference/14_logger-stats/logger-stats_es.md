
# Logger Stats

Obtiene estadísticas de un logger, incluyendo el total de eventos recibidos y el uso de almacenamiento.

## Autenticación

Requiere credenciales de datos:

- `X-Instantlogs-Event-Secret: <secret>` o `Authorization: Bearer <secret>`

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único del logger |

## Solicitud

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/stats" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

```http
GET /v1/loggers/logger_xyz789/stats HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```

## Respuesta

```json
{
  "total_events": 15234,
  "storage_bytes": 4194304,
  "events_last_hour": 234,
  "events_last_day": 5601
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Secreto de evento faltante o inválido |
| 404 | Logger no encontrado |
