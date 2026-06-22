# SDK Reference

KVNode includes a local Python configservice SDK at `kvnode/sdk/python/configservice.py`. It targets the same `/v1/collections...` REST API and replication stream used by the service.

## Package

| Language | Package | Notes |
|----------|---------|-------|
| Python | `ConfigServiceClient` from `kvnode/sdk/python/configservice.py` | Local repository package, not a published external package |

## Connecting

```python
from configservice import ConfigServiceClient

client = ConfigServiceClient(
    base_url="https://api.hola.cloud",
    mode="http",
    node="sdk-python",
)
```

## Set

`set` sends `POST /v1/collections/{collection}/keys/{key}` with a JSON `value` and returns the write response.

```python
out = client.set("my-collection", "user:alice", {"role": "admin"})
# {"ok": true, "seq": 1, "version": 1}
```

## Get

```python
entry = client.get("my-collection", "user:alice")
# {"key": "user:alice", "value": {"role": "admin"}, "version": 1, "updatedAt": "..."}
```

## List Keys

```python
keys = client.list_keys("my-collection", prefix="user:", limit=100)
# {"items": [...], "next": "..."}
```

## Delete

```python
out = client.delete("my-collection", "user:alice")
# {"ok": true, "seq": 2, "version": 2}
```

## Replication Sync

`sync` reads `POST /v1/replicate` NDJSON and applies uppercase commands such as `SET`, `DELETE`, `BASELINE_BEGIN`, and `BASELINE_END` to the local in-memory view.

```python
client.sync()
```

## Authentication

The service accepts requests when either `X-Glue-Authentication` is present or both `apikey` and `secret` are present. Missing authentication headers return `403 forbidden`.
