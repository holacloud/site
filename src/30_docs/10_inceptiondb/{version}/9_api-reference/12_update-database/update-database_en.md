# Rename a database

`PATCH /v1/databases/{databaseId}` · signed-in session

These account endpoints need the session of a signed-in hola.cloud user, not an API key. The [console](https://inceptiondb.hola.cloud/) does it for you in **API keys & settings → General**.

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
