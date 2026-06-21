
# POST /v1/queues

Crea una nueva cola.

## Autenticación

Pública — no requiere autenticación.

## Cuerpo de la Solicitud

| Campo | Tipo | Requerido | Descripción |
|-------|------|-----------|-------------|
| `name` | string | sí | Un nombre para la cola |

```json
{
  "name": "mi-cola"
}
```

## Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-cola"
  }'
```

```http
POST /v1/queues HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "name": "mi-cola"
}
```

## Respuesta

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "mi-cola",
  "length": 0,
  "reads": 0,
  "writes": 0
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Falta el campo `name` o es inválido |
| 409 | Ya existe una cola con este nombre |
