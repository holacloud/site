# Enviando Logs

InstantLogs soporta dos métodos para enviar datos de log: ingesta directa en un logger e ingesta de eventos enmarcados. Ambos métodos requieren el secreto adecuado para la autenticación.

## Autenticación

Los endpoints de ingesta de datos utilizan el Logger Secret para la autenticación. Puedes proporcionarlo de dos maneras:

- **Encabezado**: `X-Instantlogs-Event-Secret: <tu-logger-secret>`
- **Bearer Token**: `Authorization: Bearer <tu-logger-secret>`

## Método 1: Ingesta Directa

Envía entradas de log directamente a un logger específico usando el endpoint `POST /v1/loggers/{id}/ingest`.

### Entrada Individual

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "Inicio de sesión exitoso",
    "level": "info",
    "service": "auth-service",
    "timestamp": "2025-06-20T15:00:00Z",
    "metadata": {
      "user_id": "usr_123",
      "ip": "192.168.1.1"
    }
  }'
```

Respuesta esperada:

```json
{
  "ingested": 1
}
```

### Entradas Múltiples

Puedes enviar múltiples entradas de log en una sola solicitud proporcionando un array:

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '[
    {
      "message": "Solicitud iniciada",
      "level": "info",
      "service": "api-gateway",
      "timestamp": "2025-06-20T15:00:00Z"
    },
    {
      "message": "Solicitud completada",
      "level": "info",
      "service": "api-gateway",
      "timestamp": "2025-06-20T15:00:05Z"
    }
  ]'
```

Respuesta esperada:

```json
{
  "ingested": 2
}
```

## Método 2: Ingesta de Eventos Enmarcados

Para escenarios de alto rendimiento o cuando se utilizan protocolos binarios, usa el endpoint de eventos enmarcados `POST /v1/ingest/events`. Este endpoint utiliza el Event Secret para autenticación y acepta un flujo de eventos enmarcados.

```bash
curl -X POST "https://api.hola.cloud/v1/ingest/events" \
  -H "X-Instantlogs-Event-Secret: evt_secret_789" \
  -H "Content-Type: application/json" \
  -d '{
    "logger_id": "logger_xyz789",
    "events": [
      {
        "message": "Evento enmarcado 1",
        "level": "warn",
        "timestamp": "2025-06-20T15:01:00Z"
      },
      {
        "message": "Evento enmarcado 2",
        "level": "error",
        "timestamp": "2025-06-20T15:01:01Z"
      }
    ]
  }'
```

## Referencia de Solicitudes HTTP

```http
POST /v1/loggers/{id}/ingest HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
Content-Type: application/json

{
  "message": "ejemplo de log",
  "level": "info",
  "service": "web",
  "timestamp": "2025-06-20T15:00:00Z"
}

POST /v1/ingest/events HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: evt_secret_789
Content-Type: application/json

{
  "logger_id": "logger_xyz789",
  "events": [ ... ]
}
```
