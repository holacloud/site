
# Get Lambda

Obtiene los detalles de una función lambda específica por su ID.

## Autenticación

Requiere las cabeceras `Api-Key` y `Api-Secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | uuid | Identificador único de la lambda |

## Solicitud HTTP

```http
GET /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf"
```

## Respuesta

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "name": "hello-world",
  "runtime": "javascript",
  "active": true,
  "code": "export default async (req) => { return { status: 200, body: { message: 'Hello, World!' } }; }",
  "created_at": "2025-07-01T14:00:00Z",
  "updated_at": "2025-07-01T14:00:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 404 | Lambda no encontrada |
