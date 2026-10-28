# Working with data

Insert, read, update and delete documents. The examples use the `customers` collection and the environment variables of [Getting started](/docs/inceptiondb/v2/getting-started/).

## Insert documents

Send one JSON object per line (JSON Lines). One request can carry one document or millions.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  --data-binary '{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300}'
```

The response (`201 Created`) has the stored documents, with the defaults of the collection already applied:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300}
```

Good to know:

- Without `id`, the document gets one from the [defaults](/docs/inceptiondb/v2/managing-collections/) of the collection (a random UUID unless you change it).
- Documents are stored in order and the insert **stops at the first rejected document**. The ones before it stay stored.
- To load a file, use `--data-binary @customers.jsonl`.
- Each document must be a JSON object with no repeated field names. Otherwise the answer is `400 Bad Request`.
- If a [unique index](/docs/inceptiondb/v2/indexes/) already has the value, the answer is `409 Conflict`.

## Read a document by id

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/linus" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"id":"linus","document":{"country":"FI","id":"linus","mrr":0,"name":"Linus Torvalds","plan":"free"},"source":{"type":"fullscan"}}
```

`source` tells how the document was found: `index` when an index on `id` answered, `fullscan` when the collection was read. A missing document answers `404 Not Found`.

To read many documents at once, use a [query](/docs/inceptiondb/v2/querying/).

## Update documents

A patch changes only the fields you send (a [JSON merge patch](https://www.rfc-editor.org/rfc/rfc7386)): new fields are added, existing ones are replaced, and `null` removes a field.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"id": "linus"}, "patch": {"plan": "pro", "mrr": 49, "country": null}}'
```

The response has the documents as they are after the change:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49}
```

`filter` chooses the documents, with the same syntax as [queries](/docs/inceptiondb/v2/querying/). Like queries, a patch changes **one document unless you set `limit`**. Give every `team` customer the `priority` flag:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"plan": "team"}, "patch": {"priority": true}, "limit": -1}'
```

`"limit": -1` means all the matching documents. The `id` of a document cannot change.

## Delete documents

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"id": "linus"}}'
```

The response has the deleted documents:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49}
```

It deletes one document unless you set `limit`. Delete every customer of the `free` plan:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"plan": "free"}, "limit": -1}'
```

Deleted documents cannot be recovered. Run the same filter with `:find` first to check what it matches.

## Export a collection

`:find` with `"limit": -1` streams every document as JSON Lines, ready to save to a file or to insert somewhere else:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"limit": -1}' \
  -o customers.jsonl
```
