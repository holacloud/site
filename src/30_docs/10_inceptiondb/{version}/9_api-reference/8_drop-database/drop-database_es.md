# Borrar una base de datos

`DELETE /v1/databases/{databaseId}` · sesión iniciada

Estos endpoints de cuenta necesitan la sesión de un usuario identificado en hola.cloud, no una API key. La [consola](https://inceptiondb.hola.cloud/) lo hace por ti en **API keys & settings → Delete database**.

Borra la base de datos con todas sus colecciones, documentos y API keys. No se puede deshacer. Respuesta `204 No Content`.
