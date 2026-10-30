# Obtener una base de datos

`GET /v1/databases/{databaseId}` · Solo consola

Solo consola: este endpoint necesita la sesión de un usuario identificado en la consola. `Api-Key` y `Api-Secret` no funcionan con él, y no hay una forma admitida de llamarlo desde tu código. Se documenta para explicar lo que hace la consola. La [consola](https://inceptiondb.hola.cloud/) lo hace por ti en **API keys & settings**.

Devuelve la base de datos con sus owners y API keys (nunca sus secretos).

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","creation_date":"2026-10-28T05:27:03.539038505Z","owners":["owner-1"],"api_keys":[{"name":"backend","key":"4fdd0267-fded-4d51-b1a2-e7674e02126e","creation_date":"2026-10-28T05:30:12Z"}]}
```
