# Create a database

`POST /v1/databases` · Console only

Console only: this endpoint needs the session of a user signed in to the console. `Api-Key` and `Api-Secret` do not work on it, and there is no supported way to call it from your own code. It is documented to explain what the console does. The [console](https://inceptiondb.hola.cloud/) does it for you in **database selector → New database**.

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
