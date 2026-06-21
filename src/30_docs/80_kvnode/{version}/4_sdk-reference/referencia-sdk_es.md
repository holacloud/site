# Referencia Rápida de SDK

KVNode ofrece SDKs oficiales para **Go**, **Python**, **Java**, **JavaScript**, **Kotlin**, **PHP** y **Node.js**. Todos los SDKs están orientados a la misma API REST en `https://api.hola.cloud`.

---

## Instalación

| Lenguaje  | Paquete                                         | Comando                                  |
|-----------|-------------------------------------------------|------------------------------------------|
| Go        | `github.com/hola-cloud/kvnode-go`               | `go get github.com/hola-cloud/kvnode-go` |
| Python    | `hola-kvnode`                                   | `pip install hola-kvnode`                |
| Java      | `cloud.hola:kvnode-client`                      | Maven: agregar al `pom.xml`              |
| JavaScript| `@hola-cloud/kvnode-client`                     | `npm install @hola-cloud/kvnode-client`  |
| Kotlin    | `cloud.hola:kvnode-client`                      | Gradle: agregar al `build.gradle.kts`    |
| PHP       | `hola-cloud/kvnode-php`                         | `composer require hola-cloud/kvnode-php` |
| Node.js   | `@hola-cloud/kvnode-client`                     | `npm install @hola-cloud/kvnode-client`  |

---

## Conexión

Todos los SDKs se conectan usando un **host**, una **clave API** y un **secreto**. Las conexiones se agrupan internamente en un pool.

```go
import "github.com/hola-cloud/kvnode-go"

client, err := kvnode.NewClient(kvnode.Config{
    Host:   "https://api.hola.cloud",
    APIKey: "tu-api-key",
    Secret: "tu-api-secret",
})
```

```python
from hola_kvnode import KVNodeClient

client = KVNodeClient(
    host="https://api.hola.cloud",
    api_key="tu-api-key",
    secret="tu-api-secret",
)
```

```java
import cloud.hola.kvnode.KVNodeClient;

KVNodeClient client = KVNodeClient.builder()
    .host("https://api.hola.cloud")
    .apiKey("tu-api-key")
    .secret("tu-api-secret")
    .build();
```

```javascript
import { KVNodeClient } from '@hola-cloud/kvnode-client';

const client = new KVNodeClient({
    host: 'https://api.hola.cloud',
    apiKey: 'tu-api-key',
    secret: 'tu-api-secret',
});
```

```kotlin
import cloud.hola.kvnode.KVNodeClient

val client = KVNodeClient.Builder()
    .host("https://api.hola.cloud")
    .apiKey("tu-api-key")
    .secret("tu-api-secret")
    .build()
```

```php
use HolaCloud\KVNode\KVNodeClient;

$client = new KVNodeClient([
    'host'   => 'https://api.hola.cloud',
    'apiKey' => 'tu-api-key',
    'secret' => 'tu-api-secret',
]);
```

```javascript
// Node.js
const { KVNodeClient } = require('@hola-cloud/kvnode-client');

const client = new KVNodeClient({
    host: 'https://api.hola.cloud',
    apiKey: 'tu-api-key',
    secret: 'tu-api-secret',
});
```

---

## Crear una Colección

Las colecciones agrupan claves relacionadas. Créala antes de almacenar datos.

```go
err := client.CreateCollection(ctx, "mi-coleccion")
```

```python
client.create_collection("mi-coleccion")
```

```java
client.createCollection("mi-coleccion");
```

```javascript
await client.createCollection('mi-coleccion');
```

```kotlin
client.createCollection("mi-coleccion")
```

```php
$client->createCollection('mi-coleccion');
```

---

## Establecer un Par Clave-Valor

Los valores son JSON arbitrarios. Si la clave ya existe, la versión se incrementa.

```go
version, err := client.Set(ctx, "mi-coleccion", "usuario:alice", map[string]any{
    "usuario": "alice",
    "rol":     "admin",
})
```

```python
version = client.set("mi-coleccion", "usuario:alice",
    {"usuario": "alice", "rol": "admin"})
```

```java
long version = client.set("mi-coleccion", "usuario:alice",
    Map.of("usuario", "alice", "rol", "admin"));
```

```javascript
const { version } = await client.set('mi-coleccion', 'usuario:alice', {
    usuario: 'alice',
    rol: 'admin',
});
```

```kotlin
val version = client.set("mi-coleccion", "usuario:alice",
    mapOf("usuario" to "alice", "rol" to "admin"))
```

```php
$version = $client->set('mi-coleccion', 'usuario:alice', [
    'usuario' => 'alice',
    'rol'     => 'admin',
]);
```

---

## Obtener un Valor por Clave

Recupera el registro completo incluyendo valor, versión y marca de tiempo.

```go
entry, err := client.Get(ctx, "mi-coleccion", "usuario:alice")
// entry.Key, entry.Value, entry.Version, entry.UpdatedAt
```

```python
entry = client.get("mi-coleccion", "usuario:alice")
# entry.key, entry.value, entry.version, entry.updated_at
```

```java
Entry entry = client.get("mi-coleccion", "usuario:alice");
// entry.getKey(), entry.getValue(), entry.getVersion(), entry.getUpdatedAt()
```

```javascript
const entry = await client.get('mi-coleccion', 'usuario:alice');
// entry.key, entry.value, entry.version, entry.updatedAt
```

```kotlin
val entry = client.get("mi-coleccion", "usuario:alice")
// entry.key, entry.value, entry.version, entry.updatedAt
```

```php
$entry = $client->get('mi-coleccion', 'usuario:alice');
// $entry['key'], $entry['value'], $entry['version'], $entry['updatedAt']
```

---

## Listar Claves con Prefijo

Lista las claves filtradas opcionalmente por prefijo y limitadas en cantidad.

```go
entries, err := client.List(ctx, "mi-coleccion", kvnode.ListOpts{
    Prefix: "usuario:",
    Limit:  10,
})
```

```python
entries = client.list("mi-coleccion", prefix="usuario:", limit=10)
```

```java
List<Entry> entries = client.list("mi-coleccion",
    ListOpts.builder().prefix("usuario:").limit(10).build());
```

```javascript
const entries = await client.list('mi-coleccion', {
    prefix: 'usuario:',
    limit: 10,
});
```

```kotlin
val entries = client.list("mi-coleccion",
    ListOpts(prefix = "usuario:", limit = 10))
```

```php
$entries = $client->list('mi-coleccion', [
    'prefix' => 'usuario:',
    'limit'  => 10,
]);
```

---

## Eliminar una Clave

Elimina una clave de la colección.

```go
version, err := client.Delete(ctx, "mi-coleccion", "usuario:alice")
```

```python
version = client.delete("mi-coleccion", "usuario:alice")
```

```java
long version = client.delete("mi-coleccion", "usuario:alice");
```

```javascript
const { version } = await client.delete('mi-coleccion', 'usuario:alice');
```

```kotlin
val version = client.delete("mi-coleccion", "usuario:alice")
```

```php
$version = $client->delete('mi-coleccion', 'usuario:alice');
```

---

## Manejo de Errores

Los SDKs lanzan o retornan errores ante fallos de red, problemas de autenticación y solicitudes inválidas.

```go
entry, err := client.Get(ctx, "mi-coleccion", "clave-inexistente")
if errors.Is(err, kvnode.ErrNotFound) {
    // la clave no existe
}
```

```python
from hola_kvnode.exceptions import NotFoundError

try:
    entry = client.get("mi-coleccion", "clave-inexistente")
except NotFoundError:
    pass
```

```java
import cloud.hola.kvnode.exceptions.NotFoundException;

try {
    Entry entry = client.get("mi-coleccion", "clave-inexistente");
} catch (NotFoundException e) {
    // clave no encontrada
}
```

```javascript
try {
    const entry = await client.get('mi-coleccion', 'clave-inexistente');
} catch (err) {
    if (err.status === 404) {
        // clave no encontrada
    }
}
```

```kotlin
try {
    val entry = client.get("mi-coleccion", "clave-inexistente")
} catch (e: NotFoundException) {
    // clave no encontrada
}
```

```php
use HolaCloud\KVNode\Exception\NotFoundException;

try {
    $entry = $client->get('mi-coleccion', 'clave-inexistente');
} catch (NotFoundException $e) {
    // clave no encontrada
}
```

| Error            | Estado HTTP | Excepción del SDK             |
|------------------|-------------|-------------------------------|
| No Encontrado    | 404         | `ErrNotFound` / `NotFoundException` |
| No Autorizado    | 401         | `ErrUnauthorized` / `AuthException` |
| Conflicto (versión)| 409       | `ErrConflict` / `ConflictException` |
| Límite de Tasa   | 429         | `ErrRateLimit` / `RateLimitException` |
| Error Interno    | 500         | envolviendo el error HTTP     |

---

## Autenticación

KVNode soporta dos métodos de autenticación:

### Clave API + Secreto

Se envían como cabeceras `apikey` y `secret`. Es el método predeterminado en todos los SDKs.

```
apikey: tu-api-key
secret: tu-api-secret
```

### X-Glue-Authentication

Para cuentas de servicio de HolaCloud Glue, usa `X-Glue-Authentication` con un JWT firmado.

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

## Buenas Prácticas

### Pool de Conexiones

Todos los SDKs mantienen un pool de conexiones interno. En Go, usa un `http.Client` compartido. En Java/Kotlin, el pool de `OkHttpClient` se configura automáticamente. Para cargas de trabajo altas, ajusta el tamaño del pool mediante la configuración del SDK o el constructor del cliente.

### Reintentos

Los SDKs de KVNode implementan retroceso exponencial ante errores 429 (límite de tasa) y 5xx. Valores predeterminados:

| SDK      | Reintentos Máx. | Espera Inicial |
|----------|-----------------|----------------|
| Go       | 3               | 100ms          |
| Python   | 3               | 200ms          |
| Java     | 3               | 100ms          |
| JavaScript| 3              | 150ms          |
| Kotlin   | 3               | 100ms          |
| PHP      | 2               | 200ms          |

### Convenciones de Nombres de Claves

- Usa **dos puntos** como separadores de espacio de nombres: `org:proyecto:entorno:clave`
- Mantén las claves por debajo de **512 bytes** (codificación UTF-8)
- Usa `Set` con versión explícita para **bloqueo optimista**
- Las claves distinguen mayúsculas; prefiere minúsculas
- Evita claves que empiecen con `_` (reservado para metadatos internos)

### Gestión de Colecciones

- Crea las colecciones antes de usarlas; `Set` sobre una colección inexistente falla
- No hay límite en la cantidad de colecciones
- Una sola colección puede almacenar millones de claves

---

## Referencia de Endpoints

| Método  | Ruta                                              | Descripción                           |
|---------|---------------------------------------------------|---------------------------------------|
| `GET`   | `/healthz`                                        | Verificación de salud                 |
| `GET`   | `/readyz`                                         | Verificación de disponibilidad        |
| `POST`  | `/v1/collections`                                 | Crear una colección                   |
| `DELETE`| `/v1/collections/{name}`                          | Eliminar una colección                |
| `POST`  | `/v1/collections/{name}/keys/{key}`               | Establecer un par clave-valor         |
| `GET`   | `/v1/collections/{name}/keys/{key}`               | Obtener valor por clave               |
| `GET`   | `/v1/collections/{name}/keys`                     | Listar claves (query: `prefix`, `limit`) |
| `DELETE`| `/v1/collections/{name}/keys/{key}`               | Eliminar una clave                    |
| `POST`  | `/v1/collections/{name}/keys/{key}?version={v}`   | Set condicional (bloqueo optimista)   |

Todos los endpoints usan `https://api.hola.cloud` como URL base. Las cabeceras de autenticación (`apikey` + `secret` o `X-Glue-Authentication`) son obligatorias en cada solicitud.
