
# GET /v1/loggers/{id}/events

Transmite eventos de log en tiempo real. Soporta formatos NDJSON y texto plano.

## Autenticación

Requiere credenciales de datos:

- `X-Instantlogs-Event-Secret: <secret>` o `Authorization: Bearer <secret>`

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único del logger |

## Parámetros de Consulta

| Parámetro | Tipo | Por Defecto | Descripción |
|-----------|------|-------------|-------------|
| `regex` | string | — | Patrón regex para filtrar eventos por contenido del mensaje |
| `follow` | bool | `true` | Si es `true`, mantiene la conexión abierta y transmite nuevos eventos |
| `format` | string | `ndjson` | Formato de salida: `ndjson` o `text` |

## Solicitud

```bash
# Transmitir todos los eventos como NDJSON
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"

# Transmitir como texto plano
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?format=text" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"

# Filtrar con regex
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?regex=error&format=ndjson" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

```http
GET /v1/loggers/logger_xyz789/events?format=ndjson HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```

## Respuesta

Formato NDJSON (por defecto):

```
{"message":"Inicio de sesión exitoso","level":"info","service":"web","timestamp":"2025-06-20T14:22:30Z","id":"evt_001"}
{"message":"Error de conexión a la base de datos","level":"error","service":"db","timestamp":"2025-06-20T14:22:31Z","id":"evt_002"}
```

Formato texto:

```
14:22:30 [info] [web] Inicio de sesión exitoso
14:22:31 [error] [db] Error de conexión a la base de datos
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Parámetro de consulta inválido |
| 401 | Secreto de evento faltante o inválido |
| 404 | Logger no encontrado |
