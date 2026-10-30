# Referencia del API

URL base: `https://inceptiondb.hola.cloud`

## Autenticación

| Endpoints | Credenciales |
|---|---|
| Datos: todo lo que cuelga de `/v1/databases/{databaseId}/collections` | Cabeceras `Api-Key` y `Api-Secret` de una [API key](/es/docs/inceptiondb/v2/claves-api/) de esa base de datos. |
| Cuenta: bases de datos, API keys y owners | **Solo consola.** La sesión de un usuario identificado en la [consola](https://inceptiondb.hola.cloud/). `Api-Key` y `Api-Secret` no funcionan con ellos. |

Los ejemplos leen las credenciales de variables de entorno:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

## Formatos

| Endpoint | Cuerpo de la petición | Cuerpo de la respuesta |
|---|---|---|
| `:insert` | JSON Lines (JSONL): un objeto JSON por línea, no un array JSON. Envía `Content-Type: application/x-ndjson`. | JSONL: los documentos guardados. |
| `:find`, `:patch`, `:remove` | JSON. Envía `Content-Type: application/json`. | JSONL: los documentos que coinciden. |
| Todos los demás | JSON, cuando tiene cuerpo. | JSON. |

- InceptionDB interpreta el cuerpo según el endpoint, no según `Content-Type`, pero los ejemplos lo envían siempre para que cualquier cliente o proxy sepa el formato.
- `204 No Content` tiene el cuerpo vacío: ningún documento coincide, o se ha borrado algo. Comprueba el código de estado antes de decodificar el cuerpo.
- Los errores responden un objeto JSON: `{"error": {"description": "...", "message": "..."}}`. `description` suele ser el genérico `Unexpected error`; el detalle está en `message`.

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

### Cuenta (solo consola)

Estos endpoints necesitan la sesión de un usuario identificado en la consola. Se listan para explicar lo que hace la consola, no para automatizar: una API key no puede llamarlos. Ver [API keys](/es/docs/inceptiondb/v2/claves-api/).

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
| 204 No Content | Ha funcionado y no hay nada que devolver: ningún documento coincide, o es un borrado. El cuerpo está vacío. |
| 400 Bad Request | El cuerpo no es válido: ver los errores más abajo. |
| 401 Unauthorized | Un endpoint de cuenta llamado sin sesión iniciada. |
| 403 Forbidden | API key o secreto incorrectos, o sin acceso a esa base de datos. |
| 404 Not Found | La colección, el documento o el índice no existe. |
| 409 Conflict | El valor ya existe en un índice único. |
| 500 Internal Server Error | Algunos errores conocidos responden 500 hoy: ver los errores más abajo. |

## Errores y qué hacer

El `message` del error indica la causa. Estos son los que te puedes encontrar, con su mensaje real:

| Estado | `message` | Causa | Qué hacer |
|---|---|---|---|
| 403 | ``can not access database '…'`` | La clave, el secreto o el ID de la base de datos no son correctos, o la clave fue revocada. | Revisa `INCEPTIONDB_DATABASE`, `INCEPTIONDB_API_KEY` e `INCEPTIONDB_API_SECRET`. |
| 401 | `unauthorized` | Un endpoint de cuenta sin sesión de la consola. | Hazlo en la [consola](https://inceptiondb.hola.cloud/). |
| 400 | `document must be a JSON object` | Un `:insert` con un array JSON o un valor que no es un objeto. | Envía un objeto por línea (JSONL). |
| 400 | ``jsontext: unexpected EOF …``, ``jsontext: duplicate object member name …`` | Un documento de un `:insert` no es JSON válido, o repite un nombre de campo. | Corrige esa línea. Los documentos anteriores se guardaron. |
| 400 | `unexpected end of JSON input` (description `Malformed JSON`) | El cuerpo de `:find`, `:patch` o `:remove` no es JSON válido. | Revisa comillas, comas y llaves. |
| 404 | `collection not found` | `GET` de una colección que no existe. | Revisa el nombre: distingue mayúsculas y minúsculas. |
| 404 | ``document '…' not found`` | `GET` de un documento por un id que no existe. | |
| 404 | ``index '…' not found in collection '…'`` | `:getIndex` de un índice que no existe. | Lístalos con `:listIndexes`. |
| 409 | ``index add '…': index conflict: field '…' with value '…'`` | El valor ya está en un índice `map` o en un `btree` con `unique`. | Usa otro valor, o actualiza el documento existente. |
| 500 | ``collection `…` already exists`` | `POST /collections` con el nombre de una colección que existe. | Usa esa colección, u otro nombre. |
| 500 | `collection not found` | `:find`, `:patch` o `:remove` sobre una colección que no existe. | Revisa el nombre. Ten en cuenta que `:insert` crea la colección. |
| 500 | ``collection `…` not found in database `…` `` | `:dropCollection` de una colección que no existe. | No hay nada que borrar. |
| 500 | ``index '…' already exists`` | `:createIndex` con el nombre de un índice que existe. | Usa otro nombre, o borra antes el índice antiguo. |
| 500 | ``error indexing existing data: …`` | Un índice `map` o `btree` sobre un campo que algunos documentos no tienen, o un índice único sobre valores que se repiten. | Añade `"sparse": true`, o corrige los documentos. |
| 500 | ``index '…' not found, available indexes […]`` | `:find` con un `index` que no existe. | Usa uno de los índices que lista el mensaje. |
| 204 | ninguno | Un filtro con un operador que no existe, como `$gte`. | Usa `$ge` y `$le`. Ver [Consultas](/es/docs/inceptiondb/v2/consultas/). |

Las respuestas `500` de arriba son errores conocidos de la petición, no fallos del servicio. Cualquier otro `500` es un problema de nuestro lado.
