# Eliminar Clave

Elimina una clave de una colección. El comodín `*` en la ruta se reemplaza con el nombre de la clave.

## Autenticación

Requiere autenticación interna. Pasa las credenciales mediante el encabezado `X-Glue-Authentication`, o los encabezados `apikey` y `secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| col | string | Nombre de la colección |

## Ejemplo de Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/usuarios/keys/usuario:1001" \
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
  "seq": 2,
  "version": 2
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | Colección o clave no encontrada |
| 500 | internal_error | Error interno del servidor |
