# Listar Keys

Lista las claves de una colección con paginación opcional y filtrado por prefijo.

## Autenticación

Requiere autenticación interna. Pasa las credenciales mediante el encabezado `X-Glue-Authentication`, o los encabezados `apikey` y `secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| col | string | Nombre de la colección |

## Parámetros de Consulta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| limit | integer | Número máximo de claves a devolver (predeterminado: 100) |
| prefix | string | Filtrar claves por prefijo |

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/v1/collections/usuarios/keys?prefix=usuario:&limit=10" \
  -H "X-Glue-Authentication: TU_TOKEN_AUTH"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "items": [
    {
      "key": "usuario:1001",
      "value": { "name": "Alice" },
      "version": 1,
      "updatedAt": "2025-06-20T15:30:00Z"
    },
    {
      "key": "usuario:1002",
      "value": { "name": "Bob" },
      "version": 1,
      "updatedAt": "2025-06-21T08:00:00Z"
    }
  ],
  "next": "user:1002"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | Colección no encontrada |
| 500 | internal_error | Error interno del servidor |
