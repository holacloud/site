# InceptionDB API Reference

Base URL: `https://api.hola.cloud`

## Authentication

Management endpoints for `GET /v1/databases` and `POST /v1/databases` use `X-Glue-Authentication`.

Database access endpoints use `Api-Key` and `Api-Secret`. A Glue owner token can also be used where the database owner is allowed.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/v1/databases` | List databases |
| POST | `/v1/databases` | Create a database |
| GET | `/v1/databases/{databaseId}` | Get a database by ID |
| DELETE | `/v1/databases/{databaseId}` | Drop a database |
| PATCH | `/v1/databases/{databaseId}` | Update a database |
| POST | `/v1/databases/{databaseId}:createApiKey` | Create a database API key |
| POST | `/v1/databases/{databaseId}:deleteApiKey` | Delete a database API key |
| POST | `/v1/databases/{databaseId}:addOwner` | Add a database owner |
| POST | `/v1/databases/{databaseId}:deleteOwner` | Delete a database owner |
| GET | `/v1/databases/{databaseId}/collections` | List collections in a database |
| POST | `/v1/databases/{databaseId}/collections` | Create a collection |
| GET | `/v1/databases/{databaseId}/collections/{collection}` | Get a collection |
| POST | `/v1/databases/{databaseId}/collections/{collection}:dropCollection` | Drop a collection |
| POST | `/v1/databases/{databaseId}/collections/{collection}:setDefaults` | Set collection defaults |
| POST | `/v1/databases/{databaseId}/collections/{collection}:size` | Get collection size metrics |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insert` | Insert documents |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insertStream` | Insert documents from a stream |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insertFullduplex` | Insert documents over a full-duplex stream |
| POST | `/v1/databases/{databaseId}/collections/{collection}:find` | Find documents |
| POST | `/v1/databases/{databaseId}/collections/{collection}:patch` | Patch documents |
| POST | `/v1/databases/{databaseId}/collections/{collection}:remove` | Remove documents |
| GET | `/v1/databases/{databaseId}/collections/{collection}/documents/{documentId}` | Get a document by ID |
| POST | `/v1/databases/{databaseId}/collections/{collection}:listIndexes` | List collection indexes |
| POST | `/v1/databases/{databaseId}/collections/{collection}:createIndex` | Create a collection index |
| POST | `/v1/databases/{databaseId}/collections/{collection}:getIndex` | Get a collection index |
| POST | `/v1/databases/{databaseId}/collections/{collection}:dropIndex` | Drop a collection index |
| GET | `/doc` | Get a route tree for the service |
| GET | `/v1` | Get the v1 welcome response |
| GET | `/release` | Get service release metadata |
| GET | `/openapi.json` | Get the OpenAPI document |
| GET | `/v1/debug/dbs` | Debug database list |

## Common Error Codes

| Code | Description |
|------|-------------|
| 400 | Bad request: malformed syntax or invalid parameters |
| 401 | Unauthorized: missing or invalid credentials |
| 403 | Forbidden: credentials do not have access to the resource |
| 404 | Not found: the requested resource does not exist |
| 409 | Conflict: resource already exists |
| 500 | Internal server error |
