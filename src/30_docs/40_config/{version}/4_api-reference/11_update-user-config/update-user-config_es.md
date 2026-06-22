# Actualizar Config de Usuario

Actualiza el mapa de entradas de configuración del usuario autenticado.

## Autenticación

Requiere `X-Glue-Authentication`.

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "X-Glue-Authentication: TU_GLUE_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":false}}'
```

```json
{
  "entries": {
    "feature.newCheckout": false
  }
}
```
