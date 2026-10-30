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
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":2}]
```

A wrong key or secret answers `403 Forbidden`.

## Keep it safe

- InceptionDB is a server-to-server API. Do not call it from code that runs in a browser or a mobile app: anyone can read the secret there. If a web page needs data, have it call your own backend, which keeps the secret and calls InceptionDB.
- Keep the secret on your server, in environment variables or in a secret manager. Never in the code, in a repository or in a web page or mobile app: anyone who sees it can read and change your data.
- Use one API key per application or environment. You can then revoke one without affecting the others.
- If a secret leaks, create a new key, deploy it, and **revoke** the old one in **API keys & settings**. Revoked keys stop working at once.

## Owners

Owners are the hola.cloud users who manage a database from the console: its collections, API keys and other owners. Add or remove them in **API keys & settings**. A database always keeps at least one owner, and you cannot remove yourself.

## Account operations: console only

Creating, renaming and deleting databases, and managing API keys and owners, are done in the console. The [API reference](/docs/inceptiondb/v2/api-reference/) lists their endpoints to explain what the console does, but they are not meant for automation:

- `Api-Key` and `Api-Secret` do not work on them. An API key works only with the collections of its database.
- They need the session of a user signed in to the console, and there is no supported way to get one outside the console.

## See the code of a console action

Every action in the console is an HTTP request, and the console shows how to make it from your code. Click **API calls**, with the `{ }` icon, at the bottom of the side menu. It lists the latest requests; choose one to see it in curl, Go, Python and the other languages, ready to copy. Each new request also appears as a coloured dot in the narrow column at the right edge of the console; clicking a dot opens the same view.
