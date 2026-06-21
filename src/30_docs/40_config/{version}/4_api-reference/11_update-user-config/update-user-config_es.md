# PATCH /v1/config

Actualizar la configuración del usuario actual. Utiliza semántica de fusión profunda: solo se modifican los campos proporcionados en el cuerpo de la solicitud. Este endpoint requiere autenticación mediante API key.

## Autenticación

Requiere los encabezados `Api-Key` y `Api-Secret`.

## Cuerpo de la Solicitud

```json
{
  "features": {
    "new-checkout": true
  }
}
```

Se puede proporcionar cualquier subconjunto de la configuración. La actualización se fusiona con la configuración existente.

## Solicitud

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "features": {
      "new-checkout": true
    }
  }'
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
    "new-checkout": true
  }
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | Bad Request | JSON inválido en el cuerpo de la solicitud |
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 500 | Internal Server Error | Ocurrió un error inesperado |
