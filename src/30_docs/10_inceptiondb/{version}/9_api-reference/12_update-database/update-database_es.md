# Renombrar una base de datos

`PATCH /v1/databases/{databaseId}` · sesión iniciada

Estos endpoints de cuenta necesitan la sesión de un usuario identificado en hola.cloud, no una API key. La [consola](https://inceptiondb.hola.cloud/) lo hace por ti en **API keys & settings → General**.

| Campo del cuerpo | Tipo | Significado |
|---|---|---|
| `name` | string | Nombre nuevo. El id no cambia. |

Cuerpo de la petición:

```json
{"name": "my-app-prod"}
```

Respuesta `200 OK`:

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app-prod","creation_date":"2026-10-28T05:27:03.539038505Z","owners":["owner-1"],"api_keys":[]}
```
