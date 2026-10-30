# Getting started

In five minutes you will create a collection, store some customers, query them, change one and delete another. Every command can be copied and pasted as it is.

**Before you start:** you need a hola.cloud account. This guide creates the `customers` collection and leaves two customers in it, Ada and Grace. [Working with data](/docs/inceptiondb/v2/working-with-data/), [Querying](/docs/inceptiondb/v2/querying/) and [Indexes](/docs/inceptiondb/v2/indexes/) start from there.

## 1. Get your credentials

1. Open the [InceptionDB console](https://inceptiondb.hola.cloud/) and sign in.
2. Create a database (for example `my-app`), or pick one you already have.
3. Open **API keys & settings** and create an API key. The secret is shown **only once**: copy it now.

The console shows the three values ready to paste in a terminal. Export them as environment variables, so they never end up in your code:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

Check them by listing the collections of the database:

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

- `200 OK` means the credentials work. In a new database the list is empty: `[]`.
- `403 Forbidden` means that the key, the secret or the database ID is wrong. Check the three variables.

## Run the examples in other languages

Every example has a tab for curl, raw HTTP, Go, PHP, Python, Node.js and Java. All of them read the same three environment variables and print the status code and the body of the response. Save the code in a new file and run it:

| Language | Needs | Run |
|---|---|---|
| Go | Go 1.21 or newer. Only the standard library. | `go run main.go` |
| PHP | PHP 8 with the `curl` extension. | `php example.php` |
| Python | Python 3.8 or newer and `pip install requests`. | `python example.py` |
| Node.js | Node.js 18 or newer. Save the file as `.mjs`. | `node example.mjs` |
| Java | Java 17 or newer. Only the standard library. | `java Main.java` |

There is no version for the browser: the secret gives full access to your data, so it must stay on a server. See [API keys](/docs/inceptiondb/v2/api-keys/).

## 2. Create a collection

A collection groups documents. Create one called `customers`:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "customers"}'
```

Response (`201 Created`):

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

`defaults` means that a document without `id` gets a random one.

If the collection already exists (you ran this guide before), the answer is a `500` with the message ``collection `customers` already exists``. Delete it and start again from this step. This removes its documents and cannot be undone, so check the name first:

<!-- docs-check: skip -->
```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

## 3. Insert documents

Inserts use **JSON Lines** (JSONL): one JSON object per line. It is not a JSON array, so do not wrap the documents in `[...]` nor separate them with commas.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary '{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}'
```

The response (`201 Created`) has the stored documents, one per line:

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## 4. Query

Find the customers from the United Kingdom. `limit` is how many documents you want back: **without it you get only one**.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"country": "GB"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

Filters can also compare numbers, look inside arrays and nested objects, and combine conditions. The customers that pay 120 or more a month:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"mrr": {"$ge": 120}}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

> **Operator names:** InceptionDB uses `$ge` (greater or equal) and `$le` (less or equal). `$gte` and `$lte` are not valid: a filter with them matches nothing and the answer is `204 No Content`, without an error.

[Querying](/docs/inceptiondb/v2/querying/) has every operator.

## 5. Read one document

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/ada" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

The document comes in `document`:

```json
{"id":"ada","document":{"address":{"city":"London","zip":"N1"},"country":"GB","email":"ada@example.com","id":"ada","mrr":120,"name":"Ada Lovelace","plan":"pro","tags":["vip","beta"]},"source":{"type":"fullscan"}}
```

## 6. Update a document

Alan moves to the `pro` plan. A patch changes only the fields you send:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}, "patch": {"plan": "pro", "mrr": 120}}'
```

```jsonl
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"pro","country":"GB","mrr":120,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

## 7. Delete a document

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}}'
```

The response has the deleted document:

```jsonl
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"pro","country":"GB","mrr":120,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

`customers` now has two documents, Ada and Grace. Keep it: the next guides use it.

## Next steps

- [Working with data](/docs/inceptiondb/v2/working-with-data/): bulk inserts, updating and deleting many documents.
- [Querying](/docs/inceptiondb/v2/querying/): every filter operator, nested fields, arrays and pagination.
- [Indexes](/docs/inceptiondb/v2/indexes/): instant lookups by email, ranges by amount, word search.
- [Collections](/docs/inceptiondb/v2/managing-collections/): default values, size and deleting a collection.
- The [console](https://inceptiondb.hola.cloud/) shows the code of every action you do in it. [API keys](/docs/inceptiondb/v2/api-keys/) explains where.
