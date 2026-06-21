
# PATCH /api/v0/lambdas/{id}

Actualiza el código, el entorno de ejecución o el estado activo de una función lambda existente.

## Autenticación

Requiere las cabeceras `Api-Key` y `Api-Secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | uuid | Identificador único de la lambda |

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| name | string | Nuevo nombre (opcional) |
| runtime | string | Nuevo entorno de ejecución (opcional) |
| code | string | Nuevo código fuente (opcional) |
| active | boolean | Si la lambda está activa (opcional) |

## Solicitud HTTP

```http
PATCH /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "code": "export default async (req) => { return { status: 200, body: { message: 'Updated lambda' } }; }",
  "active": false
}
```

## Ejemplo

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "export default async (req) => { return { status: 200, body: { message: \"Updated lambda\" } }; }",
    "active": false
  }'
```

## Respuesta

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "name": "hello-world",
  "runtime": "javascript",
  "active": false,
  "updated_at": "2025-07-02T09:15:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Cuerpo de solicitud inválido |
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 404 | Lambda no encontrada |
