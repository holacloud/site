# Collections

A collection is a named group of JSON documents inside a database, like `customers` or `orders`. Documents in the same collection can have different fields: there is no schema to declare.

The examples use the environment variables of [Getting started](/docs/inceptiondb/v2/getting-started/).

## Create a collection

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "customers"}'
```

Response (`201 Created`):

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

| Field | Meaning |
|---|---|
| `name` | Name of the collection. You use it in every URL. |
| `total` | Number of documents. |
| `indexes` | Number of indexes. |
| `defaults` | Values added to the documents that do not have them (see below). |

## List the collections

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}]
```

## Get one collection

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

## Default values

Defaults fill in the fields that an inserted document does not have. A field sent as `null` also takes its default. Besides fixed values, there are three generated ones:

| Value | Generates |
|---|---|
| `"uuid()"` | A random unique ID, like `"35ec7d8f-96f4-4cdf-a09e-463c720e7c8c"`. |
| `"auto()"` | An increasing number: 1, 2, 3... It never repeats, though it can skip numbers. |
| `"unixnano()"` | The insert time, in nanoseconds since 1970. |

Give every new customer an ID, the `free` plan and its creation time:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:setDefaults" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"id": "uuid()", "plan": "free", "created": "unixnano()"}'
```

```json
{"created":"unixnano()","id":"uuid()","plan":"free"}
```

Now insert a customer with only a name and an email:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "Margaret Hamilton", "email": "margaret@example.com"}'
```

```jsonl
{"created":1793164963789350717,"email":"margaret@example.com","id":"35ec7d8f-96f4-4cdf-a09e-463c720e7c8c","name":"Margaret Hamilton","plan":"free"}
```

Defaults apply to the documents inserted from now on; existing documents do not change. `setDefaults` adds the fields you send to the current defaults and replaces the ones that already exist. To remove a default, send it as `null`: `{"plan": null}`.

## Size

How much memory the documents use, the size on disk and, for each index, its memory as `index.<name>`, in bytes:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:size" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"disk":265,"memory":203}
```

## Delete a collection

This deletes the collection with all its documents and indexes. It cannot be undone.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

The response is `204 No Content`.
