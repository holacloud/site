# Getting Started

This guide walks you through the basic operations of KVNode: health checks, collection management, and key-value operations.

## Health Check

Verify the node is running:

```bash
curl "https://api.hola.cloud/healthz"
```

Expected response:

```json
{"ok":true,"node":"node-1","role":"primary"}
```

## Readiness Check

Check if the node is ready to serve traffic (verifies parent connection for replicas):

```bash
curl "https://api.hola.cloud/readyz"
```

Expected response when ready:

```json
{"ok":true,"node":"node-1","role":"primary","ready":true,"checks":{"wal_replayed":true,"parent_connected":true}}
```

## Create a Collection

Collections are containers for key-value pairs. Create one with a POST request:

```bash
curl -X POST "https://api.hola.cloud/v1/collections" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret" \
  -d '{"name": "my-collection"}'
```

Expected response:

```json
{"ok":true,"collection":"my-collection"}
```

## Set a Key-Value Pair

Store a JSON value under a key inside a collection:

```bash
curl -X POST "https://api.hola.cloud/v1/collections/my-collection/keys/*" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret" \
  -d '{"value": {"user": "alice", "role": "admin"}}'
```

The `*` path segment captures the key name. To set the key `user:alice`, use:

```bash
curl -X POST "https://api.hola.cloud/v1/collections/my-collection/keys/user:alice" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret" \
  -d '{"value": {"user": "alice", "role": "admin"}}'
```

Expected response:

```json
{"ok":true,"seq":1,"version":1}
```

## Get a Key

Retrieve the value for a specific key:

```bash
curl "https://api.hola.cloud/v1/collections/my-collection/keys/user:alice" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Expected response:

```json
{"key":"user:alice","value":{"user":"alice","role":"admin"},"version":1,"updatedAt":"2025-01-01T00:00:00Z"}
```

## List Keys with Prefix

List all keys in a collection, optionally filtered by prefix and limited in count:

```bash
curl "https://api.hola.cloud/v1/collections/my-collection/keys?prefix=user:&limit=10" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Expected response:

```json
[{"key":"user:alice","value":{"user":"alice","role":"admin"},"version":1}]
```

## Delete a Key

Remove a key from a collection:

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/my-collection/keys/user:alice" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Expected response:

```json
{"ok":true,"seq":2,"version":2}
```

## Summary

You have now performed the basic KVNode operations: health checks, collection creation, and full CRUD on key-value pairs. From here, explore replication configuration and multi-language SDK integration.
