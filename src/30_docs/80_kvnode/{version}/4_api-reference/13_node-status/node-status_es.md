# GET /v1/status

Devuelve el estado actual del nodo KVNode, incluyendo colecciones, estado de replicación y tiempo de actividad.

## Autenticación

Requiere autenticación interna. Pasa las credenciales mediante el encabezado `X-Glue-Authentication`, o los encabezados `apikey` y `secret`.

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/v1/status" \
  -H "X-Glue-Authentication: TU_TOKEN_AUTH"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "node_id": "node-abc123",
  "collections": 3,
  "uptime_seconds": 86400,
  "replication": {
    "role": "primary",
    "connected_replicas": 2
  }
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | unauthorized | Autenticación faltante o inválida |
| 500 | internal_error | Error interno del servidor |
