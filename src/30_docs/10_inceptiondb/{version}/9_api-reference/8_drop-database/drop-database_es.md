# Borrar una base de datos

`DELETE /v1/databases/{databaseId}` · Solo consola

Solo consola: este endpoint necesita la sesión de un usuario identificado en la consola. `Api-Key` y `Api-Secret` no funcionan con él, y no hay una forma admitida de llamarlo desde tu código. Se documenta para explicar lo que hace la consola. La [consola](https://inceptiondb.hola.cloud/) lo hace por ti en **API keys & settings → Delete database**.

Borra la base de datos con todas sus colecciones, documentos y API keys. No se puede deshacer. Respuesta `204 No Content`.
