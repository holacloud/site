# Listar Collections

Devuelve una lista de todas las colecciones en el nodo KVNode.

## Autenticación

Requiere autenticación interna. Pasa las credenciales mediante el encabezado `X-Glue-Authentication`, o los encabezados `apikey` y `secret`.

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/v1/collections" \
  -H "X-Glue-Authentication: TU_TOKEN_AUTH"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "collections": ["usuarios", "sesiones"]
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 403 | forbidden | Missing authentication headers |
| 500 | internal_error | Error interno del servidor |
