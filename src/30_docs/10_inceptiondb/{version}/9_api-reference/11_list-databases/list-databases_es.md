# Listar bases de datos

`GET /v1/databases` · Solo consola

Solo consola: este endpoint necesita la sesión de un usuario identificado en la consola. `Api-Key` y `Api-Secret` no funcionan con él, y no hay una forma admitida de llamarlo desde tu código. Se documenta para explicar lo que hace la consola. La [consola](https://inceptiondb.hola.cloud/) lo hace por ti en el selector de base de datos.

Devuelve las bases de datos de las que eres owner.

```json
[{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","owners_length":1}]
```
