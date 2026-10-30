# API reference

Base URL: `https://inceptiondb.hola.cloud`

## Authentication

| Endpoints | Credentials |
|---|---|
| Data: everything under `/v1/databases/{databaseId}/collections` | `Api-Key` and `Api-Secret` headers of an [API key](/docs/inceptiondb/v2/api-keys/) of that database. |
| Account: databases, API keys and owners | **Console only.** The session of a user signed in to the [console](https://inceptiondb.hola.cloud/). `Api-Key` and `Api-Secret` do not work on them. |

The examples read the credentials from environment variables:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

## Formats

| Endpoint | Request body | Response body |
|---|---|---|
| `:insert` | JSON Lines (JSONL): one JSON object per line, not a JSON array. Send `Content-Type: application/x-ndjson`. | JSONL: the stored documents. |
| `:find`, `:patch`, `:remove` | JSON. Send `Content-Type: application/json`. | JSONL: the matching documents. |
| Every other endpoint | JSON, when it has a body. | JSON. |

- InceptionDB reads the body by the endpoint, not by `Content-Type`, but the examples always send it so that every client and proxy knows the format.
- `204 No Content` has an empty body: no document matched, or something was deleted. Check the status code before decoding the body.
- Errors answer a JSON object: `{"error": {"description": "...", "message": "..."}}`. `description` is often the generic `Unexpected error`; the detail is in `message`.

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

### Account (console only)

These endpoints need the session of a user signed in to the console. They are listed to explain what the console does, not for automation: an API key cannot call them. See [API keys](/docs/inceptiondb/v2/api-keys/).

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
| 204 No Content | It worked and there is nothing to return: no matching documents, or a deletion. The body is empty. |
| 400 Bad Request | The body is not valid: see the errors below. |
| 401 Unauthorized | An account endpoint called without a signed-in session. |
| 403 Forbidden | Wrong API key or secret, or no access to that database. |
| 404 Not Found | The collection, the document or the index does not exist. |
| 409 Conflict | The value already exists in a unique index. |
| 500 Internal Server Error | Some known mistakes answer 500 today: see the errors below. |

## Errors and what to do

The `message` of the error tells the cause. These are the ones you can meet, with their real message:

| Status | `message` | Cause | What to do |
|---|---|---|---|
| 403 | ``can not access database '…'`` | The key, the secret or the database ID is wrong, or the key was revoked. | Check `INCEPTIONDB_DATABASE`, `INCEPTIONDB_API_KEY` and `INCEPTIONDB_API_SECRET`. |
| 401 | `unauthorized` | An account endpoint without a console session. | Do it in the [console](https://inceptiondb.hola.cloud/). |
| 400 | `document must be a JSON object` | An `:insert` with a JSON array or a value that is not an object. | Send one object per line (JSONL). |
| 400 | ``jsontext: unexpected EOF …``, ``jsontext: duplicate object member name …`` | A document of an `:insert` is not valid JSON, or repeats a field name. | Fix that line. The documents before it were stored. |
| 400 | `unexpected end of JSON input` (description `Malformed JSON`) | The body of `:find`, `:patch` or `:remove` is not valid JSON. | Check quotes, commas and braces. |
| 404 | `collection not found` | `GET` of a collection that does not exist. | Check the name: it is case-sensitive. |
| 404 | ``document '…' not found`` | `GET` of a document by an id that does not exist. | |
| 404 | ``index '…' not found in collection '…'`` | `:getIndex` of an index that does not exist. | List them with `:listIndexes`. |
| 409 | ``index add '…': index conflict: field '…' with value '…'`` | The value is already in a `map` index or a `btree` index with `unique`. | Use another value, or update the existing document. |
| 500 | ``collection `…` already exists`` | `POST /collections` with the name of an existing collection. | Use that collection, or another name. |
| 500 | `collection not found` | `:find`, `:patch` or `:remove` on a collection that does not exist. | Check the name. Note that `:insert` creates the collection instead. |
| 500 | ``collection `…` not found in database `…` `` | `:dropCollection` of a collection that does not exist. | Nothing to delete. |
| 500 | ``index '…' already exists`` | `:createIndex` with the name of an existing index. | Use another name, or delete the old index first. |
| 500 | ``error indexing existing data: …`` | A `map` or `btree` index on a field that some documents do not have, or a unique index on values that repeat. | Add `"sparse": true`, or fix the documents. |
| 500 | ``index '…' not found, available indexes […]`` | `:find` with an `index` that does not exist. | Use one of the indexes the message lists. |
| 204 | none | A filter with an operator that does not exist, like `$gte`. | Use `$ge` and `$le`. See [Querying](/docs/inceptiondb/v2/querying/). |

The `500` answers above are known mistakes of the request, not failures of the service. Any other `500` is a problem on our side.
