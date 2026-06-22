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
  "collections": [
    {
      "name": "usuarios",
      "key_count": 1200,
      "created_at": "2025-06-01T10:00:00Z"
    },
    {
      "name": "sesiones",
      "key_count": 450,
      "created_at": "2025-06-10T14:30:00Z"
    }
  ]
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | unauthorized | Autenticación faltante o inválida |
| 500 | internal_error | Error interno del servidor |
