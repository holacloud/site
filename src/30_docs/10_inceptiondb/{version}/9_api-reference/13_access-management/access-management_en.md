# API keys and owners

These account endpoints need the session of a signed-in hola.cloud user, not an API key. The [console](https://inceptiondb.hola.cloud/) does it for you in **API keys & settings**. See [API keys](/docs/inceptiondb/v2/api-keys/) for how to use and protect them.

## Create an API key

`POST /v1/databases/{databaseId}:createApiKey`

```json
{"name": "backend"}
```

Response `201 Created`. The `secret` is returned only now; store it safely:

```json
{"name":"backend","key":"4fdd0267-fded-4d51-b1a2-e7674e02126e","secret":"02863644-fbf9-4e57-8b64-e11ead5878ec","creation_date":"2026-10-28T06:27:03.555714157+01:00"}
```

## Revoke an API key

`POST /v1/databases/{databaseId}:deleteApiKey`

```json
{"key": "4fdd0267-fded-4d51-b1a2-e7674e02126e"}
```

Response `204 No Content`. Requests with that key are rejected from then on.

## Add an owner

`POST /v1/databases/{databaseId}:addOwner`

```json
{"owner_id": "user-2"}
```

Response `200 OK` with the owners:

```json
["owner-1","user-2"]
```

## Remove an owner

`POST /v1/databases/{databaseId}:deleteOwner`

```json
{"owner_id": "user-2"}
```

Response `200 OK` with the owners that remain. You cannot remove yourself.

```json
["owner-1"]
```
