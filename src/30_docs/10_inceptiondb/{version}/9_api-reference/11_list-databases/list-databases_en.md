# List databases

`GET /v1/databases` · signed-in session

These account endpoints need the session of a signed-in hola.cloud user, not an API key. The [console](https://inceptiondb.hola.cloud/) does it for you in the database selector.

Returns the databases you own.

```json
[{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","owners_length":1}]
```
