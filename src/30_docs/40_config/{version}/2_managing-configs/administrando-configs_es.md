# Administrando Configs

La API v0 gestiona Things de configuración directamente. Un Thing de configuración tiene `id` y un objeto `entries`.

## Listar

```bash
curl "https://api.hola.cloud/v0/configs"
```

```json
[
  {
    "id": "cfg_abc123",
    "entries": {
      "feature.newCheckout": true
    }
  }
]
```

## Crear

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":true}}'
```

## Obtener

```bash
curl "https://api.hola.cloud/v0/configs/cfg_abc123"
```

## Parchear

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":false}}'
```

## Eliminar

```bash
curl -X DELETE "https://api.hola.cloud/v0/configs/cfg_abc123"
```
