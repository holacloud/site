
# POST /v1/loggers/{id}/ingest

Ingresa entradas de log en un logger.

## Autenticación

Requiere credenciales de datos:

- `X-Instantlogs-Event-Secret: <secret>` o `Authorization: Bearer <secret>`

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único del logger |

## Cuerpo de la Solicitud

Entrada de log individual:

| Campo | Tipo | Requerido | Descripción |
|-------|------|-----------|-------------|
| `message` | string | sí | El contenido del mensaje de log |
| `level` | string | no | Nivel de log (ej., `info`, `warn`, `error`) |
| `service` | string | no | Nombre del servicio de origen |
| `timestamp` | string | no | Marca de tiempo ISO 8601 (por defecto la del servidor) |

```json
{
  "message": "Inicio de sesión exitoso",
  "level": "info",
  "service": "web",
  "timestamp": "2025-06-20T14:22:30Z"
}
```

## Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "Inicio de sesión exitoso",
    "level": "info",
    "service": "web",
    "timestamp": "2025-06-20T14:22:30Z"
  }'
```

```http
POST /v1/loggers/logger_xyz789/ingest HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
Content-Type: application/json

{
  "message": "Inicio de sesión exitoso",
  "level": "info",
  "service": "web",
  "timestamp": "2025-06-20T14:22:30Z"
}
```

## Respuesta

```json
{
  "ingested": 1
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Cuerpo de solicitud faltante o inválido |
| 401 | Secreto de evento faltante o inválido |
| 404 | Logger no encontrado |
| 413 | Carga útil demasiado grande |
