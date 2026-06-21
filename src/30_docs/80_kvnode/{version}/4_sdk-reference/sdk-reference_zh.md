# SDK快速参考

KVNode 为 **Go**、**Python**、**Java**、**JavaScript**、**Kotlin**、**PHP** 和 **Node.js** 提供官方 SDK。所有 SDK 均指向 `https://api.hola.cloud` 上的同一 REST API。

---

## 安装

| 语言       | 包                                               | 命令                                   |
|-----------|--------------------------------------------------|----------------------------------------|
| Go        | `github.com/hola-cloud/kvnode-go`                | `go get github.com/hola-cloud/kvnode-go` |
| Python    | `hola-kvnode`                                    | `pip install hola-kvnode`              |
| Java      | `cloud.hola:kvnode-client`                       | Maven: 添加到 `pom.xml`                 |
| JavaScript| `@hola-cloud/kvnode-client`                      | `npm install @hola-cloud/kvnode-client` |
| Kotlin    | `cloud.hola:kvnode-client`                       | Gradle: 添加到 `build.gradle.kts`       |
| PHP       | `hola-cloud/kvnode-php`                          | `composer require hola-cloud/kvnode-php` |
| Node.js   | `@hola-cloud/kvnode-client`                      | `npm install @hola-cloud/kvnode-client` |

---

## 连接

所有 SDK 都使用 **host**、**API 密钥** 和 **secret** 进行连接。连接内部采用连接池管理。

```go
import "github.com/hola-cloud/kvnode-go"

client, err := kvnode.NewClient(kvnode.Config{
    Host:   "https://api.hola.cloud",
    APIKey: "你的API密钥",
    Secret: "你的API密钥secret",
})
```

```python
from hola_kvnode import KVNodeClient

client = KVNodeClient(
    host="https://api.hola.cloud",
    api_key="你的API密钥",
    secret="你的API密钥secret",
)
```

```java
import cloud.hola.kvnode.KVNodeClient;

KVNodeClient client = KVNodeClient.builder()
    .host("https://api.hola.cloud")
    .apiKey("你的API密钥")
    .secret("你的API密钥secret")
    .build();
```

```javascript
import { KVNodeClient } from '@hola-cloud/kvnode-client';

const client = new KVNodeClient({
    host: 'https://api.hola.cloud',
    apiKey: '你的API密钥',
    secret: '你的API密钥secret',
});
```

```kotlin
import cloud.hola.kvnode.KVNodeClient

val client = KVNodeClient.Builder()
    .host("https://api.hola.cloud")
    .apiKey("你的API密钥")
    .secret("你的API密钥secret")
    .build()
```

```php
use HolaCloud\KVNode\KVNodeClient;

$client = new KVNodeClient([
    'host'   => 'https://api.hola.cloud',
    'apiKey' => '你的API密钥',
    'secret' => '你的API密钥secret',
]);
```

```javascript
// Node.js
const { KVNodeClient } = require('@hola-cloud/kvnode-client');

const client = new KVNodeClient({
    host: 'https://api.hola.cloud',
    apiKey: '你的API密钥',
    secret: '你的API密钥secret',
});
```

---

## 创建集合

集合用于对相关键进行分组。在存储数据之前先创建集合。

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

## 设置键值对

值为任意的 JSON 数据。如果键已存在，版本号会增加。

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

## 根据键获取值

获取完整的记录，包括值、版本号和时间戳。

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

## 按前缀列出键

列出键，可按前缀过滤并限制数量。

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

## 删除键

从集合中删除一个键。

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

## 错误处理

SDK 会在网络故障、身份验证问题和无效请求时抛出或返回错误。

```go
entry, err := client.Get(ctx, "my-collection", "missing-key")
if errors.Is(err, kvnode.ErrNotFound) {
    // 键不存在
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
    // 键未找到
}
```

```javascript
try {
    const entry = await client.get('my-collection', 'missing-key');
} catch (err) {
    if (err.status === 404) {
        // 键未找到
    }
}
```

```kotlin
try {
    val entry = client.get("my-collection", "missing-key")
} catch (e: NotFoundException) {
    // 键未找到
}
```

```php
use HolaCloud\KVNode\Exception\NotFoundException;

try {
    $entry = $client->get('my-collection', 'missing-key');
} catch (NotFoundException $e) {
    // 键未找到
}
```

| 错误             | HTTP 状态码 | SDK 异常                        |
|-----------------|-------------|---------------------------------|
| 未找到           | 404         | `ErrNotFound` / `NotFoundException` |
| 未授权           | 401         | `ErrUnauthorized` / `AuthException` |
| 版本冲突         | 409         | `ErrConflict` / `ConflictException` |
| 请求频率限制     | 429         | `ErrRateLimit` / `RateLimitException` |
| 内部错误         | 500         | 包装 HTTP 错误                   |

---

## 身份验证

KVNode 支持两种身份验证方法：

### API 密钥 + Secret

通过 `apikey` 和 `secret` 请求头传递。这是所有 SDK 的默认方式。

```
apikey: 你的API密钥
secret: 你的API密钥secret
```

### X-Glue-Authentication

对于 HolaCloud Glue 服务账号，使用带有签名 JWT 的 `X-Glue-Authentication` 头。

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

## 最佳实践

### 连接池

所有 SDK 内部都维护了一个连接池。在 Go 中，使用共享的 `http.Client`。在 Java/Kotlin 中，`OkHttpClient` 的连接池会自动配置。对于高吞吐量工作负载，可通过 SDK 配置或客户端构造函数调整池大小。

### 重试机制

KVNode SDK 对 429（请求频率限制）和 5xx 错误实现了指数退避。默认值：

| SDK      | 最大重试次数 | 初始等待时间 |
|----------|-------------|-------------|
| Go       | 3           | 100ms       |
| Python   | 3           | 200ms       |
| Java     | 3           | 100ms       |
| JavaScript| 3          | 150ms       |
| Kotlin   | 3           | 100ms       |
| PHP      | 2           | 200ms       |

### 键命名规范

- 使用**冒号**作为命名空间分隔符：`org:project:env:key`
- 键的长度保持在 **512 字节**以内（UTF-8 编码）
- 使用带显式版本号的 `Set` 实现**乐观锁**
- 键区分大小写；建议使用小写
- 避免键以 `_` 开头（保留给内部元数据）

### 集合管理

- 首次使用前先创建集合；对不存在的集合执行 `Set` 会失败
- 集合数量没有限制
- 单个集合可容纳数百万个键

---

## 端点参考

| 方法     | 路径                                             | 描述                   |
|---------|--------------------------------------------------|-----------------------|
| `GET`   | `/healthz`                                       | 健康检查               |
| `GET`   | `/readyz`                                        | 就绪检查               |
| `POST`  | `/v1/collections`                                | 创建集合               |
| `DELETE`| `/v1/collections/{name}`                         | 删除集合               |
| `POST`  | `/v1/collections/{name}/keys/{key}`              | 设置键值对             |
| `GET`   | `/v1/collections/{name}/keys/{key}`              | 根据键获取值           |
| `GET`   | `/v1/collections/{name}/keys`                    | 列出键 (参数: `prefix`, `limit`) |
| `DELETE`| `/v1/collections/{name}/keys/{key}`              | 删除键                 |
| `POST`  | `/v1/collections/{name}/keys/{key}?version={v}`  | 条件设置（乐观锁）     |

所有端点均使用 `https://api.hola.cloud` 作为基础 URL。每个请求都需要提供身份验证头（`apikey` + `secret` 或 `X-Glue-Authentication`）。
