# Ejecutar Lambda (Admin)

Invoca una lambda mediante la ruta administrativa autenticada. El endpoint acepta cualquier método HTTP.

## Autenticación

Requiere `X-Glue-Authentication`.

## Parámetros de Path

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `lambda_id` | string | Identificador de la lambda |

## Cuerpo de la Solicitud

Envía cualquier carga que quieras que la lambda reciba como `req.body`.

## Solicitud HTTP

```http
POST /api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: TU_TOKEN
Content-Type: application/json

{
  "name": "Alice"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: TU_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice"
  }'
```

## Respuesta

La respuesta es lo que devuelva el manejador de la lambda.

```json
{
  "body": {
    "message": "Hello, Alice!"
  }
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Cuerpo de solicitud inválido |
| 401 | Autenticación faltante o inválida |
| 404 | Lambda no encontrada |
| 500 | Error de ejecución de Lambda |
