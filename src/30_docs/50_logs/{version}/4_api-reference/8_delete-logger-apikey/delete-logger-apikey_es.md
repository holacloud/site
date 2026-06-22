
# Delete Logger API Key

Elimina una clave API de un logger.

## Autenticación

Requiere credenciales de gestión:

- `Api-Key` — Tu clave API
- `Api-Secret` — Tu secreto de API

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único del logger |
| `key` | El ID de la clave API a eliminar |

## Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/v1/loggers/logger_xyz789/apiKeys/ak_123456" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret"
```

```http
DELETE /v1/loggers/logger_xyz789/apiKeys/ak_123456 HTTP/1.1
Host: api.hola.cloud
Api-Key: tu-api-key
Api-Secret: tu-api-secret
```

## Respuesta

HTTP `204 Sin Contenido`.

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Credenciales de API faltantes o inválidas |
| 403 | Las credenciales no tienen acceso a este logger |
| 404 | Logger o clave API no encontrados |
