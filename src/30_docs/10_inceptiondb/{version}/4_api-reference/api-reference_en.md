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
| GET | `/v1/databases/{databaseId}/collections` | List collections in a database |
| POST | `/v1/databases/{databaseId}/collections` | Create a collection |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insert` | Insert documents |
| POST | `/v1/databases/{databaseId}/collections/{collection}:find` | Find documents |
| POST | `/v1/databases/{databaseId}/collections/{collection}:patch` | Patch documents |
| POST | `/v1/databases/{databaseId}/collections/{collection}:remove` | Remove documents |
| GET | `/v1/databases/{databaseId}/collections/{collection}/documents/{documentId}` | Get a document by ID |

## Common Error Codes

| Code | Description |
|------|-------------|
| 400 | Bad request: malformed syntax or invalid parameters |
| 401 | Unauthorized: missing or invalid credentials |
| 403 | Forbidden: credentials do not have access to the resource |
| 404 | Not found: the requested resource does not exist |
| 409 | Conflict: resource already exists |
| 500 | Internal server error |
