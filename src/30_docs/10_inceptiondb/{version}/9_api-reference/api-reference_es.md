# Referencia del API

URL base: `https://inceptiondb.hola.cloud`

## Autenticación

| Endpoints | Credenciales |
|---|---|
| Datos: todo lo que cuelga de `/v1/databases/{databaseId}/collections` | Cabeceras `Api-Key` y `Api-Secret` de una [API key](/es/docs/inceptiondb/v2/claves-api/) de esa base de datos. |
| Cuenta: bases de datos, API keys y owners | La sesión de un usuario identificado en hola.cloud. Úsalos desde la [consola](https://inceptiondb.hola.cloud/). |

Los ejemplos leen las credenciales de variables de entorno:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

## Formatos

- Las peticiones envían JSON. Las inserciones envían **JSON Lines**: un objeto JSON por línea.
- `:find`, `:insert`, `:patch` y `:remove` responden **JSON Lines**, un documento por línea. Los demás endpoints responden JSON.
- Una consulta sin coincidencias responde `204 No Content` con el cuerpo vacío.
- Los errores responden un objeto JSON: `{"error": {"message": "...", "description": "..."}}`.

## Endpoints

### Datos

| Método | Ruta | Hace |
|---|---|---|
| GET | `/v1/databases/{databaseId}/collections` | [Listar colecciones](/es/docs/inceptiondb/v2/api-reference/list-collections/) |
| POST | `/v1/databases/{databaseId}/collections` | [Crear una colección](/es/docs/inceptiondb/v2/api-reference/create-collection/) |
| GET | `/v1/databases/{databaseId}/collections/{collection}` | [Obtener una colección](/es/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:setDefaults` | [Valores por defecto](/es/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:size` | [Tamaño en memoria y disco](/es/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:dropCollection` | [Borrar una colección](/es/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:insert` | [Insertar documentos](/es/docs/inceptiondb/v2/api-reference/insert-documents/) |
| POST | `…/collections/{collection}:find` | [Buscar documentos](/es/docs/inceptiondb/v2/api-reference/find-documents/) |
| GET | `…/collections/{collection}/documents/{id}` | [Obtener un documento por id](/es/docs/inceptiondb/v2/api-reference/get-document/) |
| POST | `…/collections/{collection}:patch` | [Actualizar documentos](/es/docs/inceptiondb/v2/api-reference/patch-documents/) |
| POST | `…/collections/{collection}:remove` | [Borrar documentos](/es/docs/inceptiondb/v2/api-reference/remove-documents/) |
| POST | `…/collections/{collection}:createIndex` | [Crear un índice](/es/docs/inceptiondb/v2/api-reference/index-management/) |
| POST | `…/collections/{collection}:listIndexes` | [Listar índices](/es/docs/inceptiondb/v2/api-reference/index-management/) |
| POST | `…/collections/{collection}:getIndex` | [Obtener un índice](/es/docs/inceptiondb/v2/api-reference/index-management/) |
| POST | `…/collections/{collection}:dropIndex` | [Borrar un índice](/es/docs/inceptiondb/v2/api-reference/index-management/) |

`…` equivale a `/v1/databases/{databaseId}`.

### Cuenta

| Método | Ruta | Hace |
|---|---|---|
| GET | `/v1/databases` | [Listar tus bases de datos](/es/docs/inceptiondb/v2/api-reference/list-databases/) |
| POST | `/v1/databases` | [Crear una base de datos](/es/docs/inceptiondb/v2/api-reference/create-database/) |
| GET | `/v1/databases/{databaseId}` | [Obtener una base de datos](/es/docs/inceptiondb/v2/api-reference/get-database/) |
| PATCH | `/v1/databases/{databaseId}` | [Renombrar una base de datos](/es/docs/inceptiondb/v2/api-reference/update-database/) |
| DELETE | `/v1/databases/{databaseId}` | [Borrar una base de datos](/es/docs/inceptiondb/v2/api-reference/drop-database/) |
| POST | `/v1/databases/{databaseId}:createApiKey` | [Crear una API key](/es/docs/inceptiondb/v2/api-reference/access-management/) |
| POST | `/v1/databases/{databaseId}:deleteApiKey` | [Revocar una API key](/es/docs/inceptiondb/v2/api-reference/access-management/) |
| POST | `/v1/databases/{databaseId}:addOwner` | [Añadir un owner](/es/docs/inceptiondb/v2/api-reference/access-management/) |
| POST | `/v1/databases/{databaseId}:deleteOwner` | [Quitar un owner](/es/docs/inceptiondb/v2/api-reference/access-management/) |

La descripción OpenAPI de todo el API está en [`/openapi.json`](https://inceptiondb.hola.cloud/openapi.json).

## Códigos de estado

| Código | Cuándo |
|---|---|
| 200 OK | La petición ha funcionado y hay cuerpo. |
| 201 Created | Se ha creado una colección, un documento o un índice. |
| 204 No Content | Ha funcionado y no hay nada que devolver: ningún documento coincide, o es un borrado. |
| 400 Bad Request | JSON no válido, un documento que no es un objeto JSON o nombres de campo repetidos. |
| 401 Unauthorized | Un endpoint de cuenta llamado sin sesión iniciada. |
| 403 Forbidden | API key o secreto incorrectos, o sin acceso a esa base de datos. |
| 404 Not Found | El documento no existe. |
| 409 Conflict | El valor ya existe en un índice único. |
| 500 Internal Server Error | También se devuelve cuando la colección no existe, o ya existe al crearla. |
