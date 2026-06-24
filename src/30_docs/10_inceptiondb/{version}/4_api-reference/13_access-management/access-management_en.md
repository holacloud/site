# Access Management

Database access-management actions create API keys and manage owner user IDs.

## Authentication

Requires `X-Glue-Authentication` for a user that owns the database.

## Create API Key

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:createApiKey HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"name":"production"}
```

Response status: `201 Created`

```json
{
  "name": "production",
  "key": "1abbe476-6ad6-4b97-9cca-6deb6ab2901d",
  "secret": "4bda6d52-762b-4e5d-bed7-85614c13b8bf",
  "creation_date": "2026-06-24T12:00:00Z"
}
```

The secret is returned only when the key is created.

## Delete API Key

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:deleteApiKey HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"key":"1abbe476-6ad6-4b97-9cca-6deb6ab2901d"}
```

Successful responses have an empty body.

## Add Owner

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:addOwner HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"owner_id":"user-123"}
```

Response:

```json
["current-owner","user-123"]
```

## Delete Owner

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:deleteOwner HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"owner_id":"user-123"}
```

Response:

```json
["current-owner"]
```

The authenticated user cannot delete themselves from the owner list.
