
# Agregar Owner del Logger

Agrega un owner a un logger. Los owners tienen acceso de gestión al logger.

## Autenticación

Requiere credenciales de gestión:

- `Api-Key` — Tu clave API
- `Api-Secret` — Tu secreto de API

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único del logger |

## Cuerpo de la Solicitud

| Campo | Tipo | Requerido | Descripción |
|-------|------|-----------|-------------|
| `user_id` | string | sí | El ID de usuario a agregar como owner |

```json
{
  "user_id": "user_abc123"
}
```

## Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/owners" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_abc123"
  }'
```

```http
POST /v1/loggers/logger_xyz789/owners HTTP/1.1
Host: api.hola.cloud
Api-Key: LOGGER_API_KEY
Api-Secret: LOGGER_API_SECRET
Content-Type: application/json

{
  "user_id": "user_abc123"
}
```

## Respuesta

Devuelve la lista actualizada de owners.

```json
["user_def456", "user_abc123"]
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Falta el campo `user_id` o es inválido |
| 401 | Credenciales de API faltantes o inválidas |
| 403 | Las credenciales no tienen acceso a este logger |
| 404 | Logger no encontrado |
