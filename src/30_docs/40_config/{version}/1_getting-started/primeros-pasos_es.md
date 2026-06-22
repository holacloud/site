# Primeros Pasos

Esta guía te muestra cómo recuperar y actualizar tu configuración usando la API de Config. Antes de comenzar, necesitarás una clave API con los permisos adecuados.

## Paso 1: Obtener tu Configuración

Usa el endpoint `GET /v1/config` para obtener tu configuración actual:

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret"
```

Respuesta esperada:

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

## Paso 2: Actualizar un Valor Específico

Usa `PATCH /v1/config` con un payload JSON parcial para actualizar solo los campos que necesites. Por ejemplo, para activar la nueva función de pago:

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "features": {
      "new-checkout": true
    }
  }'
```

Respuesta esperada:

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

## Paso 3: Verificar la Actualización

Vuelve a llamar a `GET /v1/config` para confirmar que los cambios se aplicaron:

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret"
```

La respuesta debería mostrar ahora `"new-checkout": true`.

## Resumen

En pocos pasos has obtenido una configuración, aplicado una actualización parcial y verificado el cambio. La API de Config facilita la gestión de ajustes en tiempo de ejecución sin necesidad de redesplegar tu aplicación.
