# Ejecutar Lambda (Pública)

Invoca una lambda públicamente. No requiere autenticación. El endpoint acepta cualquier método HTTP.

Usa esta ruta para webhooks y llamadas de API públicas. El ID de la lambda proviene de los endpoints de administración.

## Parámetros de Path

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `lambda_id` | string | Identificador de la lambda |

## Cuerpo de la Solicitud

Envía cualquier carga que quieras que la lambda reciba como `req.body`.

## Solicitud HTTP

```http
POST /run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "event": "payment_received",
  "amount": 49.99,
  "currency": "USD"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Content-Type: application/json" \
  -d '{
    "event": "payment_received",
    "amount": 49.99,
    "currency": "USD"
  }'
```

## Respuesta

La respuesta es lo que devuelva el manejador de la lambda.

```json
{
  "body": {
    "processed": true,
    "event": "payment_received"
  }
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Cuerpo de solicitud inválido |
| 404 | Lambda no encontrada |
| 500 | Error de ejecución de Lambda |
