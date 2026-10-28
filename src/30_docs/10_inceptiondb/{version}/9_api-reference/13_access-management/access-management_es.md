# API keys y owners

Estos endpoints de cuenta necesitan la sesión de un usuario identificado en hola.cloud, no una API key. La [consola](https://inceptiondb.hola.cloud/) lo hace por ti en **API keys & settings**. En [API keys](/es/docs/inceptiondb/v2/claves-api/) se explica cómo usarlas y protegerlas.

## Crear una API key

`POST /v1/databases/{databaseId}:createApiKey`

```json
{"name": "backend"}
```

Respuesta `201 Created`. El `secret` solo se devuelve ahora; guárdalo en un lugar seguro:

```json
{"name":"backend","key":"4fdd0267-fded-4d51-b1a2-e7674e02126e","secret":"02863644-fbf9-4e57-8b64-e11ead5878ec","creation_date":"2026-10-28T06:27:03.555714157+01:00"}
```

## Revocar una API key

`POST /v1/databases/{databaseId}:deleteApiKey`

```json
{"key": "4fdd0267-fded-4d51-b1a2-e7674e02126e"}
```

Respuesta `204 No Content`. Desde ese momento se rechazan las peticiones con esa clave.

## Añadir un owner

`POST /v1/databases/{databaseId}:addOwner`

```json
{"owner_id": "user-2"}
```

Respuesta `200 OK` con los owners:

```json
["owner-1","user-2"]
```

## Quitar un owner

`POST /v1/databases/{databaseId}:deleteOwner`

```json
{"owner_id": "user-2"}
```

Respuesta `200 OK` con los owners que quedan. No puedes quitarte a ti mismo.

```json
["owner-1"]
```
