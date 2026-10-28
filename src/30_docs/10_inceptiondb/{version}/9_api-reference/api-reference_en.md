# API reference

Base URL: `https://inceptiondb.hola.cloud`

## Authentication

| Endpoints | Credentials |
|---|---|
| Data: everything under `/v1/databases/{databaseId}/collections` | `Api-Key` and `Api-Secret` headers of an [API key](/docs/inceptiondb/v2/api-keys/) of that database. |
| Account: databases, API keys and owners | The session of a signed-in hola.cloud user. Use them from the [console](https://inceptiondb.hola.cloud/). |

The examples read the credentials from environment variables:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

## Formats

- Requests send JSON. Inserts send **JSON Lines**: one JSON object per line.
- `:find`, `:insert`, `:patch` and `:remove` answer **JSON Lines**, one document per line. The other endpoints answer JSON.
- A query that matches nothing answers `204 No Content` with an empty body.
- Errors answer a JSON object: `{"error": {"message": "...", "description": "..."}}`.

## Endpoints

### Data

| Method | Path | Does |
|---|---|---|
| GET | `/v1/databases/{databaseId}/collections` | [List collections](/docs/inceptiondb/v2/api-reference/list-collections/) |
| POST | `/v1/databases/{databaseId}/collections` | [Create a collection](/docs/inceptiondb/v2/api-reference/create-collection/) |
| GET | `/v1/databases/{databaseId}/collections/{collection}` | [Get a collection](/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:setDefaults` | [Set default values](/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:size` | [Memory and disk size](/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:dropCollection` | [Delete a collection](/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:insert` | [Insert documents](/docs/inceptiondb/v2/api-reference/insert-documents/) |
| POST | `…/collections/{collection}:find` | [Find documents](/docs/inceptiondb/v2/api-reference/find-documents/) |
| GET | `…/collections/{collection}/documents/{id}` | [Get a document by id](/docs/inceptiondb/v2/api-reference/get-document/) |
| POST | `…/collections/{collection}:patch` | [Update documents](/docs/inceptiondb/v2/api-reference/patch-documents/) |
| POST | `…/collections/{collection}:remove` | [Delete documents](/docs/inceptiondb/v2/api-reference/remove-documents/) |
| POST | `…/collections/{collection}:createIndex` | [Create an index](/docs/inceptiondb/v2/api-reference/index-management/) |
| POST | `…/collections/{collection}:listIndexes` | [List indexes](/docs/inceptiondb/v2/api-reference/index-management/) |
| POST | `…/collections/{collection}:getIndex` | [Get an index](/docs/inceptiondb/v2/api-reference/index-management/) |
| POST | `…/collections/{collection}:dropIndex` | [Delete an index](/docs/inceptiondb/v2/api-reference/index-management/) |

`…` stands for `/v1/databases/{databaseId}`.

### Account

| Method | Path | Does |
|---|---|---|
| GET | `/v1/databases` | [List your databases](/docs/inceptiondb/v2/api-reference/list-databases/) |
| POST | `/v1/databases` | [Create a database](/docs/inceptiondb/v2/api-reference/create-database/) |
| GET | `/v1/databases/{databaseId}` | [Get a database](/docs/inceptiondb/v2/api-reference/get-database/) |
| PATCH | `/v1/databases/{databaseId}` | [Rename a database](/docs/inceptiondb/v2/api-reference/update-database/) |
| DELETE | `/v1/databases/{databaseId}` | [Delete a database](/docs/inceptiondb/v2/api-reference/drop-database/) |
| POST | `/v1/databases/{databaseId}:createApiKey` | [Create an API key](/docs/inceptiondb/v2/api-reference/access-management/) |
| POST | `/v1/databases/{databaseId}:deleteApiKey` | [Revoke an API key](/docs/inceptiondb/v2/api-reference/access-management/) |
| POST | `/v1/databases/{databaseId}:addOwner` | [Add an owner](/docs/inceptiondb/v2/api-reference/access-management/) |
| POST | `/v1/databases/{databaseId}:deleteOwner` | [Remove an owner](/docs/inceptiondb/v2/api-reference/access-management/) |

The OpenAPI description of the whole API is at [`/openapi.json`](https://inceptiondb.hola.cloud/openapi.json).

## Status codes

| Code | When |
|---|---|
| 200 OK | The request worked and there is a body. |
| 201 Created | A collection, document or index was created. |
| 204 No Content | It worked and there is nothing to return: no matching documents, or a deletion. |
| 400 Bad Request | Invalid JSON, a document that is not a JSON object, or repeated field names. |
| 401 Unauthorized | An account endpoint called without a signed-in session. |
| 403 Forbidden | Wrong API key or secret, or no access to that database. |
| 404 Not Found | The document does not exist. |
| 409 Conflict | The value already exists in a unique index. |
| 500 Internal Server Error | Also returned when the collection does not exist, or already exists when you create it. |
