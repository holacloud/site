# Primeros pasos con InstantLogs

Esta guía te ayudará a crear un logger, ingresar una entrada de log y transmitir eventos en tiempo real.

## Paso 1: Crear un Logger

Primero, crea un logger usando la API de gestión. Necesitas una clave API con permisos de gestión.

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "logger-de-mi-app"
  }'
```

Respuesta esperada:

```json
{
  "id": "logger_xyz789",
  "name": "logger-de-mi-app",
  "secret": "lgs_abc123def456",
  "created_at": "2025-06-20T14:22:00Z"
}
```

Guarda el valor de `secret` de forma segura — se utiliza para autenticar operaciones de datos.

## Paso 2: Ingresar una Entrada de Log

Usa el Logger Secret para enviar una entrada de log a tu logger:

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "¡Hola desde mi app!",
    "level": "info",
    "service": "web",
    "timestamp": "2025-06-20T14:22:30Z"
  }'
```

Respuesta esperada:

```json
{
  "ingested": 1
}
```

## Paso 3: Transmitir Logs en Tiempo Real

Abre una conexión de streaming para ver los logs a medida que llegan:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

El servidor mantendrá la conexión abierta y enviará nuevas entradas de log como NDJSON:

```
{"message":"¡Hola desde mi app!","level":"info","service":"web","timestamp":"2025-06-20T14:22:30Z","id":"evt_001"}
```

Con el encabezado `Accept: text/plain`, puedes recibir líneas de texto plano en su lugar.

## Paso 4: Filtrar el Flujo

Agrega un patrón de expresión regular para filtrar eventos:

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=error" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

Solo se transmitirán las entradas de log cuyo mensaje coincida con el patrón `error`.

## Resumen

Creaste un logger, ingresaste un log, transmitiste eventos en vivo y aplicaste un filtro con expresión regular. InstantLogs facilita la centralización y observación de los logs de tu aplicación en tiempo real.
