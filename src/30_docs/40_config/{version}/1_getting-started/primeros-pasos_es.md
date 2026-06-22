# Primeros Pasos

Esta guía muestra la forma actual de la API Config. El endpoint de usuario v1 requiere autenticación Glue.

## Leer Configuración de Usuario

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "X-Glue-Authentication: TU_GLUE_TOKEN"
```

Respuesta:

```json
{
  "entries": {
    "feature.newCheckout": true,
    "database.host": "db.example.com"
  }
}
```

## Actualizar Configuración de Usuario

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "X-Glue-Authentication: TU_GLUE_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "entries": {
      "feature.newCheckout": false
    }
  }'
```

Respuesta:

```json
{
  "entries": {
    "feature.newCheckout": false
  }
}
```

La API v0 gestiona Things de configuración directamente. Cada Thing tiene `id` y `entries`.
