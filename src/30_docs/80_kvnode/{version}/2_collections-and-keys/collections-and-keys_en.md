# Collections and Keys

KVNode organizes data into **collections**, each containing **key-value** pairs. This document covers collection management and key operations in detail.

## Collection Management

### Create a Collection

```bash
curl -X POST "https://api.hola.cloud/v1/collections" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret" \
  -d '{"name": "my-collection"}'
```

Response:

```json
{"ok":true,"collection":"my-collection"}
```

Collections are created instantly. If the collection already exists, the operation is idempotent.

### List Collections

```bash
curl "https://api.hola.cloud/v1/collections" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Response:

```json
{"collections":["my-collection","config","sessions"]}
```

### Delete a Collection

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/my-collection" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Response:

```json
{"ok":true,"collection":"my-collection"}
```

Deleting a collection removes all keys stored in it. This operation cannot be undone.

## Key Operations

KVNode keys are strings that can contain any UTF-8 characters. Values must be valid JSON.

### Set a Key

```bash
curl -X POST "https://api.hola.cloud/v1/collections/my-collection/keys/settings:theme" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret" \
  -d '{"value": {"mode": "dark", "fontSize": 14}}'
```

Response includes the sequence number and version:

```json
{"ok":true,"seq":3,"version":1}
```

Each successful write returns a monotonically increasing `seq` (global sequence) and `version` (per-key version), useful for optimistic concurrency control.

### Get a Key

```bash
curl "https://api.hola.cloud/v1/collections/my-collection/keys/settings:theme" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Response:

```json
{"key":"settings:theme","value":{"mode":"dark","fontSize":14},"version":1,"updatedAt":"2025-01-01T00:00:00Z"}
```

Attempting to get a non-existent key returns a 404 error.

### Delete a Key

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/my-collection/keys/settings:theme" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Response:

```json
{"ok":true,"seq":4,"version":2}
```

### List Keys with Prefix and Limit

```bash
curl "https://api.hola.cloud/v1/collections/my-collection/keys?prefix=settings:&limit=20" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Both `prefix` and `limit` are optional query parameters:
- `prefix` — filters keys that start with the given string
- `limit` — caps the number of returned records (default: no limit)

Response:

```json
[
  {"key":"settings:theme","value":{"mode":"dark","fontSize":14},"version":1,"updatedAt":"2025-01-01T00:00:00Z"},
  {"key":"settings:locale","value":"en-US","version":1,"updatedAt":"2025-01-01T00:00:01Z"}
]
```

## Key Naming Conventions

While KVNode accepts any UTF-8 string as a key, we recommend following these conventions:

- **Use namespacing with colons**: `app:module:key` (e.g. `config:database:host`)
- **Avoid special characters** that may conflict with URL encoding (`/`, `?`, `#`)
- **Keep keys readable** — meaningful names simplify debugging
- **Be consistent** — choose a naming pattern and apply it across all collections

## Value Types

All values must be valid JSON. KVNode supports:

- **Objects**: `{"user": "alice", "role": "admin"}`
- **Arrays**: `["item1", "item2"]`
- **Strings**: `"hello world"`
- **Numbers**: `42` or `3.14`
- **Booleans**: `true` or `false`
- **Null**: `null`

The maximum value size is determined by the WAL backend configuration.
