# Working with data

Insert, read, update and delete documents.

**Before you start:** this guide uses the `customers` collection as [Getting started](/docs/inceptiondb/v2/getting-started/) leaves it, with Ada and Grace, and its environment variables. It adds two temporary customers marked with `"trial": true` and deletes them at the end, so `customers` is left as it was.

## Insert documents

Inserts use **JSON Lines** (JSONL): one JSON object per line. One request can carry one document or millions.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary '{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0,"trial":true}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true}'
```

The response (`201 Created`) has the stored documents, with the defaults of the collection already applied:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0,"trial":true}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true}
```

A JSON array is **not** JSONL. This body is rejected with `400 Bad Request` and the message `document must be a JSON object`:

```text
[{"id":"linus"},{"id":"barbara"}]
```

The same two documents in JSONL:

```text
{"id":"linus"}
{"id":"barbara"}
```

Good to know:

- Without `id`, the document gets one from the [defaults](/docs/inceptiondb/v2/managing-collections/) of the collection (a random UUID unless you change it).
- InceptionDB does not check that ids are unique: inserting the same `id` twice stores two documents. To prevent it, create a [`map` index](/docs/inceptiondb/v2/indexes/) on `id`. Then a repeated id answers `409 Conflict`, and reads by id are instant.
- Documents are stored in order and the insert **stops at the first rejected document**. The ones before it stay stored.
- To load a file, use `--data-binary @customers.jsonl`. `-d` would remove the line breaks.
- Each document must be a JSON object with no repeated field names. Otherwise the answer is `400 Bad Request`.
- If a [unique index](/docs/inceptiondb/v2/indexes/) already has the value, the answer is `409 Conflict`.
- Inserting into a collection that does not exist creates it. Check the name of the collection in the URL.

## Read a document by id

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/linus" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"id":"linus","document":{"country":"FI","id":"linus","mrr":0,"name":"Linus Torvalds","plan":"free","trial":true},"source":{"type":"fullscan"}}
```

`source` tells how the document was found: `index` when an index on `id` answered, `fullscan` when the collection was read. A missing document answers `404 Not Found`.

To read many documents at once, use a [query](/docs/inceptiondb/v2/querying/).

## Update documents

A patch changes only the fields you send (a [JSON merge patch](https://www.rfc-editor.org/rfc/rfc7386)): new fields are added, existing ones are replaced, and `null` removes a field.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "linus"}, "patch": {"plan": "pro", "mrr": 49, "country": null}}'
```

The response has the documents as they are after the change:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49,"trial":true}
```

`filter` chooses the documents, with the same syntax as [queries](/docs/inceptiondb/v2/querying/). Like queries, a patch changes **one document unless you set `limit`**. Give every trial customer the `priority` flag:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"trial": true}, "patch": {"priority": true}, "limit": -1}'
```

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49,"trial":true,"priority":true}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true,"priority":true}
```

`"limit": -1` means all the matching documents. When nothing matches, the answer is `204 No Content`.

A patch can also change the `id` of a document. If the collection has a `map` index on `id`, a patch to an id that already exists is not applied.

## Delete documents

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "linus"}}'
```

The response has the deleted documents:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49,"trial":true,"priority":true}
```

It deletes one document unless you set `limit`. Delete every trial customer, which also cleans up this guide:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"trial": true}, "limit": -1}'
```

```jsonl
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true,"priority":true}
```

Deleted documents cannot be recovered. Run the same filter with `:find` first to check what it matches.

## Export a collection

`:find` with `"limit": -1` returns every document as JSONL, ready to save to a file or to insert somewhere else:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"limit": -1}' \
  -o customers.jsonl
```

`customers.jsonl` has Ada and Grace, one per line: the collection is back as it was before this guide.
