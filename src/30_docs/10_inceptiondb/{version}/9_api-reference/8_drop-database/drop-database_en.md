# Delete a database

`DELETE /v1/databases/{databaseId}` · Console only

Console only: this endpoint needs the session of a user signed in to the console. `Api-Key` and `Api-Secret` do not work on it, and there is no supported way to call it from your own code. It is documented to explain what the console does. The [console](https://inceptiondb.hola.cloud/) does it for you in **API keys & settings → Delete database**.

Deletes the database with every collection, document and API key. It cannot be undone. Response `204 No Content`.
