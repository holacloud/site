# Collection management

Get, configure, measure and delete a collection. All of them need an API key.

## Get a collection

`GET /v1/databases/{databaseId}/collections/{collection}`

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":3}
```

## Set default values

`POST …/collections/{collection}:setDefaults`

The fields you send are added to the current defaults, replacing the ones with the same name; a field sent as `null` removes that default. Besides fixed values, `"uuid()"` generates a random id, `"auto()"` an increasing number and `"unixnano()"` the insert time in nanoseconds. A default applies when the document does not have the field or has it as `null`.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:setDefaults" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"id": "uuid()", "plan": "free", "created": "unixnano()"}'
```

Response `200 OK` with all the defaults now in use:

```json
{"created":"unixnano()","id":"uuid()","plan":"free"}
```

## Size

`POST …/collections/{collection}:size`

Bytes of memory used by the documents (`memory`) and by each index (`index.<name>`), and size on disk (`disk`).

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:size" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"disk":630,"index.by_email":260,"memory":646}
```

## Delete a collection

`POST …/collections/{collection}:dropCollection`

Deletes the collection, its documents and its indexes. It cannot be undone. Response `204 No Content`.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```
