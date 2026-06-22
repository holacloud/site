# Obtener Config de Usuario

Recupera el mapa de entradas de configuración del usuario autenticado.

## Autenticación

Requiere `X-Glue-Authentication`.

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "X-Glue-Authentication: TU_GLUE_TOKEN"
```

```json
{
  "entries": {
    "feature.newCheckout": true
  }
}
```
