# API keys

Your applications reach a database with an **API key**: a key and a secret sent in two headers. You manage databases, API keys and owners in the [InceptionDB console](https://inceptiondb.hola.cloud/), signed in with your hola.cloud account.

## Create a database

In the console, open the database selector and choose **New database**. It is ready in about a second. Its ID is on the overview page: that is the value of `INCEPTIONDB_DATABASE`.

## Create an API key

1. Choose the database and open **API keys & settings**.
2. Click **New API key** and name it after the application that will use it, like `backend-production`.
3. Copy the secret. It is shown **only once**: InceptionDB stores only a hash of it.

The console gives you the credentials ready to use, for example as environment variables:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

An API key reaches every collection of its database, and only that database.

## Use it

Send the key and the secret in the `Api-Key` and `Api-Secret` headers of every request:

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":3}]
```

A wrong key or secret answers `403 Forbidden`.

## Keep it safe

- Keep the secret on your server, in environment variables or in a secret manager. Never in the code, in a repository or in a web page or mobile app: anyone who sees it can read and change your data.
- Use one API key per application or environment. You can then revoke one without affecting the others.
- If a secret leaks, create a new key, deploy it, and **revoke** the old one in **API keys & settings**. Revoked keys stop working at once.

## Owners

Owners are the hola.cloud users who manage a database from the console: its collections, API keys and other owners. Add or remove them in **API keys & settings**. A database always keeps at least one owner, and you cannot remove yourself.

## Account operations and the API

Creating, renaming and deleting databases, managing API keys and owners are done in the console. They are HTTP endpoints too (see the [API reference](/docs/inceptiondb/v2/api-reference/)), but they need the session of a signed-in user: an API key cannot call them. Every action of the console shows its request: click the balls on the right of the screen.
