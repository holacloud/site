
# Eliminar Owner del Logger

Elimina un owner de un logger. Los owners no pueden eliminarse a sí mismos.

## Autenticación

Requiere credenciales de gestión:

- `Api-Key` — Tu clave API
- `Api-Secret` — Tu secreto de API

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único del logger |
| `ownerId` | El ID de usuario del owner a eliminar |

## Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/v1/loggers/logger_xyz789/owners/user_abc123" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

```http
DELETE /v1/loggers/logger_xyz789/owners/user_abc123 HTTP/1.1
Host: api.hola.cloud
Api-Key: LOGGER_API_KEY
Api-Secret: LOGGER_API_SECRET
```

## Respuesta

Devuelve la lista actualizada de owners.

```json
["user_def456"]
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | El owner no puede eliminarse a sí mismo |
| 401 | Credenciales de API faltantes o inválidas |
| 403 | Las credenciales no tienen acceso a este logger |
| 404 | Logger no encontrado |
