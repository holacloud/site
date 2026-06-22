# Establecer Clave

Establece el valor de una clave en una colección. El comodín `*` en la ruta se reemplaza con el nombre de la clave.

## Autenticación

Requiere autenticación interna. Pasa las credenciales mediante el encabezado `X-Glue-Authentication`, o los encabezados `apikey` y `secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| col | string | Nombre de la colección |

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| value | any | El valor JSON a almacenar |

```json
{
  "value": {
    "name": "Alicia",
    "email": "alicia@example.com",
    "role": "admin"
  }
}
```

## Ejemplo de Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/collections/usuarios/keys/usuario:1001" \
  -H "X-Glue-Authentication: TU_TOKEN_AUTH" \
  -H "Content-Type: application/json" \
  -d '{
    "value": {
      "name": "Alicia",
      "email": "alicia@example.com",
      "role": "admin"
    }
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
  "seq": 1,
  "version": 1
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | invalid_json | Valor faltante o inválido |
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | Colección no encontrada |
| 500 | internal_error | Error interno del servidor |
