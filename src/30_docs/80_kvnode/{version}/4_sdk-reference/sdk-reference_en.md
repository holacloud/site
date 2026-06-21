# SDK Quick Reference

KVNode provides official SDKs for **Go**, **Python**, **Java**, **JavaScript**, **Kotlin**, **PHP**, and **Node.js**. All SDKs target the same REST API at `https://api.hola.cloud`.

---

## Installation

| Language   | Package                                          | Command                                  |
|-----------|--------------------------------------------------|------------------------------------------|
| Go        | `github.com/hola-cloud/kvnode-go`                | `go get github.com/hola-cloud/kvnode-go` |
| Python    | `hola-kvnode`                                    | `pip install hola-kvnode`                |
| Java      | `cloud.hola:kvnode-client`                       | Maven: add to `pom.xml`                  |
| JavaScript| `@hola-cloud/kvnode-client`                      | `npm install @hola-cloud/kvnode-client`  |
| Kotlin    | `cloud.hola:kvnode-client`                       | Gradle: add to `build.gradle.kts`        |
| PHP       | `hola-cloud/kvnode-php`                          | `composer require hola-cloud/kvnode-php` |
| Node.js   | `@hola-cloud/kvnode-client`                      | `npm install @hola-cloud/kvnode-client`  |

---

## Connecting

All SDKs connect using a **host**, **API key**, and **secret**. Connections are typically pooled internally.

```go
import "github.com/hola-cloud/kvnode-go"

client, err := kvnode.NewClient(kvnode.Config{
    Host:   "https://api.hola.cloud",
    APIKey: "your-api-key",
    Secret: "your-api-secret",
})
```

```python
from hola_kvnode import KVNodeClient

client = KVNodeClient(
    host="https://api.hola.cloud",
    api_key="your-api-key",
    secret="your-api-secret",
)
```

```java
import cloud.hola.kvnode.KVNodeClient;

KVNodeClient client = KVNodeClient.builder()
    .host("https://api.hola.cloud")
    .apiKey("your-api-key")
    .secret("your-api-secret")
    .build();
```

```javascript
import { KVNodeClient } from '@hola-cloud/kvnode-client';

const client = new KVNodeClient({
    host: 'https://api.hola.cloud',
    apiKey: 'your-api-key',
    secret: 'your-api-secret',
});
```

```kotlin
import cloud.hola.kvnode.KVNodeClient

val client = KVNodeClient.Builder()
    .host("https://api.hola.cloud")
    .apiKey("your-api-key")
    .secret("your-api-secret")
    .build()
```

```php
use HolaCloud\KVNode\KVNodeClient;

$client = new KVNodeClient([
    'host'   => 'https://api.hola.cloud',
    'apiKey' => 'your-api-key',
    'secret' => 'your-api-secret',
]);
```

// Node.js
```javascript
const { KVNodeClient } = require('@hola-cloud/kvnode-client');

const client = new KVNodeClient({
    host: 'https://api.hola.cloud',
    apiKey: 'your-api-key',
    secret: 'your-api-secret',
});
```

---

## Creating a Collection

Collections group related keys. Create one before storing data.

```go
err := client.CreateCollection(ctx, "my-collection")
```

```python
client.create_collection("my-collection")
```

```java
client.createCollection("my-collection");
```

```javascript
await client.createCollection('my-collection');
```

```kotlin
client.createCollection("my-collection")
```

```php
$client->createCollection('my-collection');
```

---

## Setting a Key-Value Pair

Values are arbitrary JSON. If the key already exists the version is incremented.

```go
version, err := client.Set(ctx, "my-collection", "user:alice", map[string]any{
    "user": "alice",
    "role": "admin",
})
```

```python
version = client.set("my-collection", "user:alice",
    {"user": "alice", "role": "admin"})
```

```java
long version = client.set("my-collection", "user:alice",
    Map.of("user", "alice", "role", "admin"));
```

```javascript
const { version } = await client.set('my-collection', 'user:alice', {
    user: 'alice',
    role: 'admin',
});
```

```kotlin
val version = client.set("my-collection", "user:alice",
    mapOf("user" to "alice", "role" to "admin"))
```

```php
$version = $client->set('my-collection', 'user:alice', [
    'user' => 'alice',
    'role' => 'admin',
]);
```

---

## Getting a Value by Key

Retrieve the full record including value, version, and timestamp.

```go
entry, err := client.Get(ctx, "my-collection", "user:alice")
// entry.Key, entry.Value, entry.Version, entry.UpdatedAt
```

```python
entry = client.get("my-collection", "user:alice")
# entry.key, entry.value, entry.version, entry.updated_at
```

```java
Entry entry = client.get("my-collection", "user:alice");
// entry.getKey(), entry.getValue(), entry.getVersion(), entry.getUpdatedAt()
```

```javascript
const entry = await client.get('my-collection', 'user:alice');
// entry.key, entry.value, entry.version, entry.updatedAt
```

```kotlin
val entry = client.get("my-collection", "user:alice")
// entry.key, entry.value, entry.version, entry.updatedAt
```

```php
$entry = $client->get('my-collection', 'user:alice');
// $entry['key'], $entry['value'], $entry['version'], $entry['updatedAt']
```

---

## Listing Keys with Prefix

List keys optionally filtered by prefix and limited in count.

```go
entries, err := client.List(ctx, "my-collection", kvnode.ListOpts{
    Prefix: "user:",
    Limit:  10,
})
```

```python
entries = client.list("my-collection", prefix="user:", limit=10)
```

```java
List<Entry> entries = client.list("my-collection",
    ListOpts.builder().prefix("user:").limit(10).build());
```

```javascript
const entries = await client.list('my-collection', {
    prefix: 'user:',
    limit: 10,
});
```

```kotlin
val entries = client.list("my-collection",
    ListOpts(prefix = "user:", limit = 10))
```

```php
$entries = $client->list('my-collection', [
    'prefix' => 'user:',
    'limit'  => 10,
]);
```

---

## Deleting a Key

Remove a key from the collection.

```go
version, err := client.Delete(ctx, "my-collection", "user:alice")
```

```python
version = client.delete("my-collection", "user:alice")
```

```java
long version = client.delete("my-collection", "user:alice");
```

```javascript
const { version } = await client.delete('my-collection', 'user:alice');
```

```kotlin
val version = client.delete("my-collection", "user:alice")
```

```php
$version = $client->delete('my-collection', 'user:alice');
```

---

## Error Handling

SDKs throw or return errors on network failures, authentication issues, and invalid requests.

```go
entry, err := client.Get(ctx, "my-collection", "missing-key")
if errors.Is(err, kvnode.ErrNotFound) {
    // key does not exist
}
```

```python
from hola_kvnode.exceptions import NotFoundError

try:
    entry = client.get("my-collection", "missing-key")
except NotFoundError:
    pass
```

```java
import cloud.hola.kvnode.exceptions.NotFoundException;

try {
    Entry entry = client.get("my-collection", "missing-key");
} catch (NotFoundException e) {
    // key not found
}
```

```javascript
try {
    const entry = await client.get('my-collection', 'missing-key');
} catch (err) {
    if (err.status === 404) {
        // key not found
    }
}
```

```kotlin
try {
    val entry = client.get("my-collection", "missing-key")
} catch (e: NotFoundException) {
    // key not found
}
```

```php
use HolaCloud\KVNode\Exception\NotFoundException;

try {
    $entry = $client->get('my-collection', 'missing-key');
} catch (NotFoundException $e) {
    // key not found
}
```

| Error              | HTTP Status | SDK Exception              |
|--------------------|-------------|----------------------------|
| Not Found          | 404         | `ErrNotFound` / `NotFoundException` |
| Unauthorized       | 401         | `ErrUnauthorized` / `AuthException` |
| Conflict (version) | 409         | `ErrConflict` / `ConflictException` |
| Rate Limited       | 429         | `ErrRateLimit` / `RateLimitException` |
| Internal Error     | 500         | wrapping the HTTP error    |

---

## Authentication

KVNode supports two authentication methods:

### API Key + Secret

Passed as `apikey` and `secret` headers. This is the default for all SDKs.

```
apikey: your-api-key
secret: your-api-secret
```

### X-Glue-Authentication

For HolaCloud Glue service accounts, set `X-Glue-Authentication` with a signed JWT instead.

```javascript
const client = new KVNodeClient({
    host: 'https://api.hola.cloud',
    glueAuth: 'glue-jwt-token',
});
```

```python
client = KVNodeClient(
    host="https://api.hola.cloud",
    glue_auth="glue-jwt-token",
)
```

---

## Best Practices

### Connection Pooling

All SDKs maintain a connection pool internally. In Go, use a shared `http.Client`. In Java/Kotlin, the `OkHttpClient` connection pool is configured automatically. For high-throughput workloads, tune pool size via SDK config or client constructor.

### Retry Logic

KVNode SDKs implement exponential backoff on 429 (rate limit) and 5xx errors. Defaults:

| SDK      | Max Retries | Initial Backoff |
|----------|-------------|-----------------|
| Go       | 3           | 100ms           |
| Python   | 3           | 200ms           |
| Java     | 3           | 100ms           |
| JavaScript| 3          | 150ms           |
| Kotlin   | 3           | 100ms           |
| PHP      | 2           | 200ms           |

### Key Naming Conventions

- Use **colons** as namespace separators: `org:project:env:key`
- Keep keys under **512 bytes** (UTF-8 encoded)
- Use `Set` with explicit version for **optimistic locking**
- Keys are case-sensitive; prefer lowercase
- Avoid keys starting with `_` (reserved for internal metadata)

### Collection Management

- Create collections before first use; `Set` on a missing collection fails
- There is no limit on collection count
- A single collection can hold millions of keys

---

## Endpoint Reference

| Method   | Path                                               | Description                        |
|----------|----------------------------------------------------|------------------------------------|
| `GET`    | `/healthz`                                         | Health check                       |
| `GET`    | `/readyz`                                          | Readiness check                    |
| `POST`   | `/v1/collections`                                  | Create a collection                |
| `DELETE` | `/v1/collections/{name}`                           | Delete a collection                |
| `POST`   | `/v1/collections/{name}/keys/{key}`                | Set a key-value pair               |
| `GET`    | `/v1/collections/{name}/keys/{key}`                | Get value by key                   |
| `GET`    | `/v1/collections/{name}/keys`                      | List keys (query: `prefix`, `limit`) |
| `DELETE` | `/v1/collections/{name}/keys/{key}`                | Delete a key                       |
| `POST`   | `/v1/collections/{name}/keys/{key}?version={v}`    | Conditional set (optimistic lock)  |

All endpoints use `https://api.hola.cloud` as base URL. Authentication headers (`apikey` + `secret` or `X-Glue-Authentication`) are required on every request.
