# Crear una base de datos

`POST /v1/databases` · Solo consola

Solo consola: este endpoint necesita la sesión de un usuario identificado en la consola. `Api-Key` y `Api-Secret` no funcionan con él, y no hay una forma admitida de llamarlo desde tu código. Se documenta para explicar lo que hace la consola. La [consola](https://inceptiondb.hola.cloud/) lo hace por ti en **selector de base de datos → New database**.

| Campo del cuerpo | Tipo | Significado |
|---|---|---|
| `name` | string | Nombre de la base de datos. |

Cuerpo de la petición:

```json
{"name": "my-app"}
```

Respuesta `201 Created`. El `id` es el valor de `INCEPTIONDB_DATABASE`:

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","creation_date":"2026-10-28T05:27:03.539038505Z"}
```
