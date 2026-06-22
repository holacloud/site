# Crear Colección

Crea una nueva colección en el nodo KVNode.

## Autenticación

Requiere autenticación interna. Pasa las credenciales mediante el encabezado `X-Glue-Authentication`, o los encabezados `apikey` y `secret`.

## Cuerpo de la Solicitud

```json
{
  "name": "usuarios"
}
```

## Ejemplo de Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/collections" \
  -H "X-Glue-Authentication: TU_TOKEN_AUTH" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "usuarios"
  }'
```

## Ejemplo de Respuesta

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "name": "usuarios",
  "key_count": 0,
  "created_at": "2025-06-21T10:00:00Z"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | invalid_request | Nombre de colección faltante o inválido |
| 401 | unauthorized | Autenticación faltante o inválida |
| 409 | conflict | La colección ya existe |
| 500 | internal_error | Error interno del servidor |
