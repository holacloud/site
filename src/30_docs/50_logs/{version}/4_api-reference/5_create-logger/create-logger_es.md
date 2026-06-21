
# POST /v1/loggers

Crea un nuevo logger.

## Autenticación

Requiere credenciales de gestión:

- `Api-Key` — Tu clave API
- `Api-Secret` — Tu secreto de API

## Cuerpo de la Solicitud

| Campo | Tipo | Requerido | Descripción |
|-------|------|-----------|-------------|
| `name` | string | sí | Un nombre legible para el logger |

```json
{
  "name": "logger-de-mi-app"
}
```

## Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "logger-de-mi-app"
  }'
```

```http
POST /v1/loggers HTTP/1.1
Host: api.hola.cloud
Api-Key: tu-api-key
Api-Secret: tu-api-secret
Content-Type: application/json

{
  "name": "logger-de-mi-app"
}
```

## Respuesta

```json
{
  "id": "logger_xyz789",
  "name": "logger-de-mi-app",
  "secret": "lgs_abc123def456",
  "created_at": "2025-06-20T14:22:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Falta el campo `name` o es inválido |
| 401 | Credenciales de API faltantes o inválidas |
| 409 | Ya existe un logger con este nombre |
