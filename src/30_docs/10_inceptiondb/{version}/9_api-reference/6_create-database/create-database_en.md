# Create a database

`POST /v1/databases` · signed-in session

These account endpoints need the session of a signed-in hola.cloud user, not an API key. The [console](https://inceptiondb.hola.cloud/) does it for you in **database selector → New database**.

| Body field | Type | Meaning |
|---|---|---|
| `name` | string | Name of the database. |

Request body:

```json
{"name": "my-app"}
```

Response `201 Created`. The `id` is the value of `INCEPTIONDB_DATABASE`:

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","creation_date":"2026-10-28T05:27:03.539038505Z"}
```
