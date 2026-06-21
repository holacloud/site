# Managing Configurations

The admin API (`/v0/configs`) provides full CRUD operations for managing configurations. These endpoints are publicly accessible and do not require an API key.

## Listing All Configurations

Retrieve a list of all stored configurations:

```bash
curl "https://api.hola.cloud/v0/configs"
```

Expected response:

```json
[
  {
    "id": "cfg_abc123",
    "project": "my-app",
    "environment": "production",
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-06-20T14:22:00Z"
  },
  {
    "id": "cfg_def456",
    "project": "my-app",
    "environment": "staging",
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-06-19T09:00:00Z"
  }
]
```

## Creating a Configuration

Create a new configuration entry:

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{
    "project": "my-app",
    "environment": "production",
    "config": {
      "database": {
        "host": "db.example.com",
        "port": 5432
      }
    }
  }'
```

Expected response:

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "config": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    }
  },
  "created_at": "2025-06-20T14:22:00Z",
  "updated_at": "2025-06-20T14:22:00Z"
}
```

## Retrieving a Configuration

Fetch a single configuration by its ID:

```bash
curl "https://api.hola.cloud/v0/configs/cfg_abc123"
```

Expected response:

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "config": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    }
  },
  "created_at": "2025-01-15T10:30:00Z",
  "updated_at": "2025-06-20T14:22:00Z"
}
```

## Patching a Configuration

Apply a partial update to an existing configuration:

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{
    "config": {
      "database": {
        "host": "db-new.example.com"
      }
    }
  }'
```

Expected response:

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "config": {
    "database": {
      "host": "db-new.example.com",
      "port": 5432
    }
  },
  "created_at": "2025-01-15T10:30:00Z",
  "updated_at": "2025-06-20T15:00:00Z"
}
```

## Deleting a Configuration

Remove a configuration permanently:

```bash
curl -X DELETE "https://api.hola.cloud/v0/configs/cfg_abc123"
```

Expected response: `HTTP 204 No Content`

## HTTP Request Reference

```http
GET /v0/configs HTTP/1.1
Host: api.hola.cloud

POST /v0/configs HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "project": "my-app",
  "environment": "production",
  "config": { ... }
}

GET /v0/configs/{id} HTTP/1.1
Host: api.hola.cloud

PATCH /v0/configs/{id} HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "config": { ... }
}

DELETE /v0/configs/{id} HTTP/1.1
Host: api.hola.cloud
```
