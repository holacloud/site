
# Create Logger API Key

Crea una nueva clave API para un logger. Las claves API permiten acceso de gestión al logger.

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
| `name` | string | sí | Una etiqueta para identificar esta clave API |

```json
{
  "name": "clave-ci-cd"
}
```

## Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/apiKeys" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "clave-ci-cd"
  }'
```

```http
POST /v1/loggers/logger_xyz789/apiKeys HTTP/1.1
Host: api.hola.cloud
Api-Key: tu-api-key
Api-Secret: tu-api-secret
Content-Type: application/json

{
  "name": "clave-ci-cd"
}
```

## Respuesta

```json
{
  "id": "ak_123456",
  "name": "clave-ci-cd",
  "key": "lgs_api_key_value",
  "created_at": "2025-06-20T15:00:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Falta el campo `name` o es inválido |
| 401 | Credenciales de API faltantes o inválidas |
| 403 | Las credenciales no tienen acceso a este logger |
| 404 | Logger no encontrado |
