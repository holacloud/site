# Referencia SDK

KVNode incluye un SDK local Python configservice en `kvnode/sdk/python/configservice.py`. Usa la misma API REST `/v1/collections...` y el stream de replicacion del servicio.

## Paquete

| Lenguaje | Paquete | Notas |
|----------|---------|-------|
| Python | `ConfigServiceClient` de `kvnode/sdk/python/configservice.py` | Paquete local del repositorio, no paquete externo publicado |

## Conexion

```python
from configservice import ConfigServiceClient

client = ConfigServiceClient(
    base_url="https://api.hola.cloud",
    mode="http",
    node="sdk-python",
)
```

## Operaciones

```python
out = client.set("my-collection", "user:alice", {"role": "admin"})
# {"ok": true, "seq": 1, "version": 1}

entry = client.get("my-collection", "user:alice")
keys = client.list_keys("my-collection", prefix="user:", limit=100)
# {"items": [...], "next": "..."}

out = client.delete("my-collection", "user:alice")
# {"ok": true, "seq": 2, "version": 2}
```

## Sincronizacion

`sync` lee `POST /v1/replicate` como NDJSON y aplica comandos en mayusculas como `SET`, `DELETE`, `BASELINE_BEGIN` y `BASELINE_END`.

## Autenticacion

El servicio acepta solicitudes si esta presente `X-Glue-Authentication` o si estan presentes `apikey` y `secret`. Si faltan los headers de autenticacion devuelve `403 forbidden`.
