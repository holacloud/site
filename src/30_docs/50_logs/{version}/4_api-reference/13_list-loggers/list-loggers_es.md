
# GET /v1/loggers

Lista todos los loggers asociados con tu cuenta.

## Autenticación

Requiere credenciales de gestión:

- `Api-Key` — Tu clave API
- `Api-Secret` — Tu secreto de API

## Solicitud

```bash
curl "https://api.hola.cloud/v1/loggers" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret"
```

```http
GET /v1/loggers HTTP/1.1
Host: api.hola.cloud
Api-Key: tu-api-key
Api-Secret: tu-api-secret
```

## Respuesta

```json
[
  {
    "id": "logger_xyz789",
    "name": "my-app-logger",
    "created_at": "2025-06-20T14:22:00Z"
  },
  {
    "id": "logger_abc123",
    "name": "staging-logger",
    "created_at": "2025-06-19T10:00:00Z"
  }
]
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Credenciales de API faltantes o inválidas |
| 403 | Las credenciales no tienen permisos de listado |
