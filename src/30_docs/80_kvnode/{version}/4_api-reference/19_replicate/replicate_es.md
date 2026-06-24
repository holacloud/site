# Replicate

Replica datos desde un nodo padre a este KVNode.

## Autenticación

Requiere autenticación interna. Pasa las credenciales mediante el encabezado `X-Glue-Authentication`, o los encabezados `apikey` y `secret`.

## Ejemplo de Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/replicate" \
  -H "X-Glue-Authentication: TU_TOKEN_AUTH" \
  -H "Content-Type: application/json" \
  -d '{
    "entries": [
      {
        "collection": "users",
        "key": "user:1001",
        "value": {"name": "Alice"},
        "seq": 42
      }
    ]
  }'
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "ok": true,
  "applied": 1
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | invalid_json | Carga JSON inválida |
| 403 | forbidden | Missing authentication headers |
| 502 | parent_unavailable | El nodo padre no está accesible |
| 500 | internal_error | Error interno del servidor |
