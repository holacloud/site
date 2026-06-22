# Nodo Métricas

Devuelve métricas operativas del nodo KVNode, incluyendo conteos de escritura y lectura.

## Autenticación

Requiere autenticación interna. Pasa las credenciales mediante el encabezado `X-Glue-Authentication`, o los encabezados `apikey` y `secret`.

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/v1/metrics" \
  -H "X-Glue-Authentication: TU_TOKEN_AUTH"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "writes_total": 15000,
  "reads_total": 98000,
  "collections": 3,
  "keys_total": 4500,
  "uptime_seconds": 86400
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | unauthorized | Autenticación faltante o inválida |
| 500 | internal_error | Error interno del servidor |
