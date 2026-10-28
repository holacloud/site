# InceptionDB

InceptionDB is a JSON document database that you use over HTTP. There is no driver to install and no schema to define: you send JSON, you get JSON back, from any language that can make an HTTP request.

## How it is organized

| Concept | What it is |
|---|---|
| **Database** | Your space. It holds collections and has its own API keys. You create it in the [console](https://inceptiondb.hola.cloud/). |
| **Collection** | A named group of documents, like `customers` or `orders`. Documents in a collection can have different fields. |
| **Document** | A JSON object. Every document has an `id`; if you do not send one, InceptionDB generates it. |
| **Index** | Optional. Makes lookups and range queries on a field fast, and can enforce unique values. |
| **API key** | A key and a secret that your application sends in the `Api-Key` and `Api-Secret` headers to reach the data of one database. |

Every request goes to `https://inceptiondb.hola.cloud`, and the examples read your credentials from three environment variables, so you can copy and paste them as they are:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

## Where to start

- **[Getting started](/docs/inceptiondb/v2/getting-started/)**: from zero to your first query in five minutes.
- **[Collections](/docs/inceptiondb/v2/managing-collections/)**: create collections, set default values, check their size.
- **[Working with data](/docs/inceptiondb/v2/working-with-data/)**: insert, read, update and delete documents.
- **[Querying](/docs/inceptiondb/v2/querying/)**: filters, operators, nested fields, arrays and pagination.
- **[Indexes](/docs/inceptiondb/v2/indexes/)**: make queries fast and keep values unique.
- **[API keys](/docs/inceptiondb/v2/api-keys/)**: create databases and keys, and keep them safe.
- **[API reference](/docs/inceptiondb/v2/api-reference/)**: every endpoint with its request and response.

## Good fits

- **Application data**: users, settings, carts and orders: any JSON whose shape changes over time.
- **Content**: articles, pages and product catalogs with fields that vary per item.
- **Events and telemetry**: device readings and product analytics written at high speed.
- **Prototypes**: start storing data in minutes and add indexes once the queries are clear.
