# Referencia de la API de InceptionDB

URL base: `https://api.hola.cloud`

## Autenticación

Los endpoints de management `GET /v1/databases` y `POST /v1/databases` usan `X-Glue-Authentication`.

Los endpoints de acceso a la base de datos usan `Api-Key` y `Api-Secret`. También se puede usar un token Glue de owner cuando el propietario de la base de datos está permitido.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/v1/databases` | Listar bases de datos |
| POST | `/v1/databases` | Crear una base de datos |
| GET | `/v1/databases/{databaseId}` | Obtener una base de datos por ID |
| DELETE | `/v1/databases/{databaseId}` | Eliminar una base de datos |
| PATCH | `/v1/databases/{databaseId}` | Actualizar una base de datos |
| GET | `/v1/databases/{databaseId}/collections` | Listar colecciones de una base de datos |
| POST | `/v1/databases/{databaseId}/collections` | Crear una colección |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insert` | Insertar documentos |
| POST | `/v1/databases/{databaseId}/collections/{collection}:find` | Buscar documentos |
| POST | `/v1/databases/{databaseId}/collections/{collection}:patch` | Modificar documentos |
| POST | `/v1/databases/{databaseId}/collections/{collection}:remove` | Eliminar documentos |
| GET | `/v1/databases/{databaseId}/collections/{collection}/documents/{documentId}` | Obtener un documento por ID |

## Códigos de error comunes

| Código | Descripción |
|--------|-------------|
| 400 | Solicitud incorrecta: sintaxis o parámetros inválidos |
| 401 | No autorizado: credenciales faltantes o inválidas |
| 403 | Prohibido: las credenciales no tienen acceso al recurso |
| 404 | No encontrado: el recurso solicitado no existe |
| 409 | Conflicto: el recurso ya existe |
| 500 | Error interno del servidor |
