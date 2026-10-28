# Get a database

`GET /v1/databases/{databaseId}` · signed-in session

These account endpoints need the session of a signed-in hola.cloud user, not an API key. The [console](https://inceptiondb.hola.cloud/) does it for you in **API keys & settings**.

Returns the database with its owners and API keys (never their secrets).

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","creation_date":"2026-10-28T05:27:03.539038505Z","owners":["owner-1"],"api_keys":[{"name":"backend","key":"4fdd0267-fded-4d51-b1a2-e7674e02126e","creation_date":"2026-10-28T05:30:12Z"}]}
```
