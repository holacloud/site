# Delete a database

`DELETE /v1/databases/{databaseId}` · signed-in session

These account endpoints need the session of a signed-in hola.cloud user, not an API key. The [console](https://inceptiondb.hola.cloud/) does it for you in **API keys & settings → Delete database**.

Deletes the database with every collection, document and API key. It cannot be undone. Response `204 No Content`.
