# Obtener Usuario Actual

Devuelve información sobre el usuario autenticado.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud HTTP

```http
GET /api/v0/me HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: TU_TOKEN
```

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/api/v0/me" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

## Respuesta

```json
{
  "id": "user_123",
  "email": "user@example.com",
  "created_at": "2025-01-15T08:00:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Autenticación faltante o inválida |
