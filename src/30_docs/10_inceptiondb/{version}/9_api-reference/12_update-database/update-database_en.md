# Rename a database

`PATCH /v1/databases/{databaseId}` · Console only

Console only: this endpoint needs the session of a user signed in to the console. `Api-Key` and `Api-Secret` do not work on it, and there is no supported way to call it from your own code. It is documented to explain what the console does. The [console](https://inceptiondb.hola.cloud/) does it for you in **API keys & settings → General**.

| Body field | Type | Meaning |
|---|---|---|
| `name` | string | New name. The id does not change. |

Request body:

```json
{"name": "my-app-prod"}
```

Response `200 OK`:

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app-prod","creation_date":"2026-10-28T05:27:03.539038505Z","owners":["owner-1"],"api_keys":[]}
```
