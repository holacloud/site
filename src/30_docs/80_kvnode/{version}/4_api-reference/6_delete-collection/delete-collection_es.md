# Eliminar Colección

Elimina una colección y todas sus claves.

## Autenticación

Requiere autenticación interna. Pasa las credenciales mediante el encabezado `X-Glue-Authentication`, o los encabezados `apikey` y `secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| col | string | Nombre de la colección |

## Ejemplo de Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/usuarios" \
  -H "X-Glue-Authentication: TU_TOKEN_AUTH"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "ok": true,
  "collection": "usuarios"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | Colección no encontrada |
| 500 | internal_error | Error interno del servidor |
