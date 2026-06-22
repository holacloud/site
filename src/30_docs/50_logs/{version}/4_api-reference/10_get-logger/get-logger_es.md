
# Get Logger

Obtiene los detalles de un logger específico por su ID.

## Autenticación

Requiere credenciales de gestión:

- `Api-Key` — Tu clave API
- `Api-Secret` — Tu secreto de API

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único del logger |

## Solicitud

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret"
```

```http
GET /v1/loggers/logger_xyz789 HTTP/1.1
Host: api.hola.cloud
Api-Key: tu-api-key
Api-Secret: tu-api-secret
```

## Respuesta

```json
{
  "id": "logger_xyz789",
  "name": "my-app-logger",
  "secret": "lgs_abc123def456",
  "created_at": "2025-06-20T14:22:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Credenciales de API faltantes o inválidas |
| 403 | Las credenciales no tienen acceso a este logger |
| 404 | Logger no encontrado |
