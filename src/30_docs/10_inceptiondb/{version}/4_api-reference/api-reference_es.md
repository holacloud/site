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
| POST | `/v1/databases/{databaseId}:createApiKey` | Crear una API key de base de datos |
| POST | `/v1/databases/{databaseId}:deleteApiKey` | Eliminar una API key de base de datos |
| POST | `/v1/databases/{databaseId}:addOwner` | Agregar un owner a la base de datos |
| POST | `/v1/databases/{databaseId}:deleteOwner` | Eliminar un owner de la base de datos |
| GET | `/v1/databases/{databaseId}/collections` | Listar colecciones de una base de datos |
| POST | `/v1/databases/{databaseId}/collections` | Crear una colección |
| GET | `/v1/databases/{databaseId}/collections/{collection}` | Obtener una colección |
| POST | `/v1/databases/{databaseId}/collections/{collection}:dropCollection` | Eliminar una colección |
| POST | `/v1/databases/{databaseId}/collections/{collection}:setDefaults` | Definir defaults de una colección |
| POST | `/v1/databases/{databaseId}/collections/{collection}:size` | Obtener métricas de tamaño de una colección |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insert` | Insertar documentos |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insertStream` | Insertar documentos desde un stream |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insertFullduplex` | Insertar documentos con stream full-duplex |
| POST | `/v1/databases/{databaseId}/collections/{collection}:find` | Buscar documentos |
| POST | `/v1/databases/{databaseId}/collections/{collection}:patch` | Modificar documentos |
| POST | `/v1/databases/{databaseId}/collections/{collection}:remove` | Eliminar documentos |
| GET | `/v1/databases/{databaseId}/collections/{collection}/documents/{documentId}` | Obtener un documento por ID |
| POST | `/v1/databases/{databaseId}/collections/{collection}:listIndexes` | Listar índices de una colección |
| POST | `/v1/databases/{databaseId}/collections/{collection}:createIndex` | Crear un índice de colección |
| POST | `/v1/databases/{databaseId}/collections/{collection}:getIndex` | Obtener un índice de colección |
| POST | `/v1/databases/{databaseId}/collections/{collection}:dropIndex` | Eliminar un índice de colección |
| GET | `/doc` | Obtener el árbol de rutas del servicio |
| GET | `/v1` | Obtener la respuesta de bienvenida de v1 |
| GET | `/release` | Obtener metadatos de release del servicio |
| GET | `/openapi.json` | Obtener el documento OpenAPI |
| GET | `/v1/debug/dbs` | Listado debug de bases de datos |

## Códigos de error comunes

| Código | Descripción |
|--------|-------------|
| 400 | Solicitud incorrecta: sintaxis o parámetros inválidos |
| 401 | No autorizado: credenciales faltantes o inválidas |
| 403 | Prohibido: las credenciales no tienen acceso al recurso |
| 404 | No encontrado: el recurso solicitado no existe |
| 409 | Conflicto: el recurso ya existe |
| 500 | Error interno del servidor |
