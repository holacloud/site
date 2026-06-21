# Streaming y Filtrado de Logs

InstantLogs proporciona endpoints de streaming HTTP de larga duración que entregan eventos de log en tiempo casi real. Puedes filtrar flujos por patrón de expresión regular y elegir entre formatos de salida NDJSON y texto plano.

## Endpoints de Streaming

### Flujo de Eventos

El endpoint `GET /v1/loggers/{id}/events` transmite todas las entradas de log a medida que se ingieren. Por defecto, los eventos se devuelven como NDJSON (JSON delimitado por nuevas líneas).

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

Cada línea del cuerpo de la respuesta es un objeto JSON que representa una entrada de log:

```
{"id":"evt_001","message":"Solicitud iniciada","level":"info","service":"api-gateway","timestamp":"2025-06-20T15:00:00Z"}
{"id":"evt_002","message":"Solicitud completada","level":"info","service":"api-gateway","timestamp":"2025-06-20T15:00:05Z"}
```

Para recibir salida en texto plano, configura el encabezado `Accept` como `text/plain`:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Accept: text/plain"
```

### Flujo Filtrado

El endpoint `GET /v1/loggers/{id}/filter` aplica un filtro de expresión regular antes de transmitir. Solo se entregan las entradas de log cuyo mensaje coincide con el patrón.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=error|critical" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

Esto transmite solo las entradas que contienen `error` o `critical` en el campo message. El parámetro pattern acepta cualquier expresión regular válida de Go.

### Combinar Filtro y Formato

Puedes combinar el filtrado por expresión regular con tu formato de salida preferido:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=^\\[WARN\\]" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Accept: text/plain"
```

## Estadísticas del Logger

El endpoint `GET /v1/loggers/{id}/stats` devuelve estadísticas sobre un logger sin consumir el flujo de eventos.

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/stats" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

Respuesta esperada:

```json
{
  "logger_id": "logger_xyz789",
  "total_events": 15234,
  "events_last_hour": 342,
  "storage_bytes": 1048576,
  "created_at": "2025-06-20T14:22:00Z"
}
```

## Referencia de Solicitudes HTTP

```http
GET /v1/loggers/{id}/events HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456

GET /v1/loggers/{id}/filter?pattern=error HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456

GET /v1/loggers/{id}/stats HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```
