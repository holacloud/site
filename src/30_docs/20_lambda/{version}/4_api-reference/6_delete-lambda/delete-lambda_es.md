# Eliminar Lambda

Elimina una lambda permanentemente.

## Autenticación

Requiere `X-Glue-Authentication`.

## Parámetros de Path

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `lambda_id` | string | Identificador de la lambda |

## Solicitud HTTP

```http
DELETE /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: TU_TOKEN
```

## Ejemplo

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

## Respuesta

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Autenticación faltante o inválida |
| 404 | Lambda no encontrada |
