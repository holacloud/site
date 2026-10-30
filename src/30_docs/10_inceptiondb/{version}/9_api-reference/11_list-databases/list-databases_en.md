# List databases

`GET /v1/databases` · Console only

Console only: this endpoint needs the session of a user signed in to the console. `Api-Key` and `Api-Secret` do not work on it, and there is no supported way to call it from your own code. It is documented to explain what the console does. The [console](https://inceptiondb.hola.cloud/) does it for you in the database selector.

Returns the databases you own.

```json
[{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","owners_length":1}]
```
