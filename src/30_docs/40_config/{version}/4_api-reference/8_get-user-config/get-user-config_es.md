# GET /v1/config

Obtener la configuración del usuario actual. Este endpoint requiere autenticación mediante API key.

## Autenticación

Requiere los encabezados `Api-Key` y `Api-Secret`.

## Solicitud

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

## Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "project": "my-app",
  "environment": "production",
  "services": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "cache": {
      "host": "redis.example.com",
      "port": 6379
    }
  },
  "features": {
    "new-checkout": false
  }
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 500 | Internal Server Error | Ocurrió un error inesperado |
