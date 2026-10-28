# Listar bases de datos

`GET /v1/databases` · sesión iniciada

Estos endpoints de cuenta necesitan la sesión de un usuario identificado en hola.cloud, no una API key. La [consola](https://inceptiondb.hola.cloud/) lo hace por ti en el selector de base de datos.

Devuelve las bases de datos de las que eres owner.

```json
[{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","owners_length":1}]
```
