# POST /v1/buckets

Crear un nuevo bucket. El nombre del bucket debe ser único globalmente en todas las cuentas.

## Autenticación

Requiere los encabezados `Api-Key` y `Api-Secret`.

## Cuerpo de la Solicitud

```json
{
  "name": "mi-nuevo-bucket",
  "public": false
}
```

| Campo | Tipo | Requerido | Descripción |
|-------|------|-----------|-------------|
| `name` | string | sí | Nombre único global (3-63 caracteres, minúsculas, números y guiones) |
| `public` | boolean | no | Si los archivos pueden accederse sin autenticación (por defecto: `false`) |

## Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-nuevo-bucket",
    "public": false
  }'
```

## Respuesta

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "bkt_xyz789",
  "name": "mi-nuevo-bucket",
  "createdAt": "2026-06-21T12:00:00Z",
  "size": 0,
  "fileCount": 0,
  "public": false
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | Bad Request | Nombre de bucket inválido o campos requeridos faltantes |
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 409 | Conflict | Ya existe un bucket con este nombre |
| 500 | Internal Server Error | Ocurrió un error inesperado |
