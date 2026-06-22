
# Filter Logs

Transmite y filtra entradas de log en tiempo real usando un patrón de expresión regular.

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
| `pattern` | string | — | Patrón regex para filtrar entradas de log por contenido del mensaje |
| `follow` | bool | `false` | Si es `true`, mantiene la conexión abierta y transmite nuevas entradas |

## Solicitud

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=error&follow=true" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

```http
GET /v1/loggers/logger_xyz789/filter?pattern=error&follow=true HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```

## Respuesta

Devuelve las entradas de log coincidentes como NDJSON:

```
{"message":"Error de conexión a la base de datos","level":"error","service":"db","timestamp":"2025-06-20T14:22:31Z","id":"evt_002"}
{"message":"Tiempo de espera excedido","level":"error","service":"web","timestamp":"2025-06-20T14:22:32Z","id":"evt_003"}
```

Con `follow=true`, la conexión permanece abierta y las nuevas entradas coincidentes se transmiten a medida que llegan.

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Patrón regex inválido |
| 401 | Secreto de evento faltante o inválido |
| 404 | Logger no encontrado |
