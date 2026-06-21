
# POST /api/v0/run/{id}

Invoca una función lambda usando credenciales de administrador. Requiere autenticación.

## Autenticación

Requiere las cabeceras `Api-Key` y `Api-Secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | uuid | Identificador único de la lambda a ejecutar |

## Cuerpo de la Solicitud

Cualquier carga útil JSON que desee pasar a la función lambda. La lambda lo recibe como el parámetro `req`.

## Solicitud HTTP

```http
POST /api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "name": "Alice"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice"
  }'
```

## Respuesta

```json
{
  "status": 200,
  "body": {
    "message": "¡Hola, Alice!"
  }
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Cuerpo de solicitud inválido |
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 404 | Lambda no encontrada o inactiva |
| 500 | Error de ejecución de la lambda |
