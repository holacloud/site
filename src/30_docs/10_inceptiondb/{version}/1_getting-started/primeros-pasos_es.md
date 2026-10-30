# Primeros pasos

En cinco minutos vas a crear una colección, guardar unos clientes, consultarlos, cambiar uno y borrar otro. Todos los comandos se pueden copiar y pegar tal cual.

**Antes de empezar:** necesitas una cuenta de hola.cloud. Esta guía crea la colección `customers` y deja en ella dos clientes, Ada y Grace. [Trabajando con datos](/es/docs/inceptiondb/v2/trabajando-con-datos/), [Consultas](/es/docs/inceptiondb/v2/consultas/) e [Índices](/es/docs/inceptiondb/v2/indices/) parten de ahí.

## 1. Consigue tus credenciales

1. Abre la [consola de InceptionDB](https://inceptiondb.hola.cloud/) e inicia sesión.
2. Crea una base de datos (por ejemplo `my-app`), o elige una que ya tengas.
3. Abre **API keys & settings** y crea una API key. El secreto se muestra **una sola vez**: cópialo ahora.

La consola muestra los tres valores listos para pegar en un terminal. Expórtalos como variables de entorno, para que nunca acaben en tu código:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

Compruébalas listando las colecciones de la base de datos:

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

- `200 OK` significa que las credenciales funcionan. En una base de datos nueva la lista está vacía: `[]`.
- `403 Forbidden` significa que la clave, el secreto o el ID de la base de datos no son correctos. Revisa las tres variables.

## Ejecuta los ejemplos en otros lenguajes

Cada ejemplo tiene una pestaña para curl, HTTP, Go, PHP, Python, Node.js y Java. Todas leen las mismas tres variables de entorno e imprimen el código de estado y el cuerpo de la respuesta. Guarda el código en un archivo nuevo y ejecútalo:

| Lenguaje | Necesita | Ejecuta |
|---|---|---|
| Go | Go 1.21 o posterior. Solo la biblioteca estándar. | `go run main.go` |
| PHP | PHP 8 con la extensión `curl`. | `php example.php` |
| Python | Python 3.8 o posterior y `pip install requests`. | `python example.py` |
| Node.js | Node.js 18 o posterior. Guarda el archivo como `.mjs`. | `node example.mjs` |
| Java | Java 17 o posterior. Solo la biblioteca estándar. | `java Main.java` |

No hay versión para el navegador: el secreto da acceso completo a tus datos, así que debe quedarse en un servidor. Consulta [API keys](/es/docs/inceptiondb/v2/claves-api/).

## 2. Crea una colección

Una colección agrupa documentos. Crea una llamada `customers`:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "customers"}'
```

Respuesta (`201 Created`):

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

`defaults` significa que un documento sin `id` recibe uno aleatorio.

Si la colección ya existe (ya seguiste esta guía antes), la respuesta es un `500` con el mensaje ``collection `customers` already exists``. Bórrala y empieza otra vez desde este paso. Esto elimina sus documentos y no se puede deshacer, así que comprueba antes el nombre:

<!-- docs-check: skip -->
```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

## 3. Inserta documentos

Las inserciones usan **JSON Lines** (JSONL): un objeto JSON por línea. No es un array JSON, así que no envuelvas los documentos en `[...]` ni los separes con comas.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary '{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}'
```

La respuesta (`201 Created`) contiene los documentos guardados, uno por línea:

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## 4. Consulta

Busca los clientes del Reino Unido. `limit` es cuántos documentos quieres recibir: **sin él recibes solo uno**.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"country": "GB"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

Los filtros también comparan números, miran dentro de arrays y objetos anidados, y combinan condiciones. Los clientes que pagan 120 o más al mes:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"mrr": {"$ge": 120}}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

> **Nombres de operadores:** InceptionDB usa `$ge` (mayor o igual) y `$le` (menor o igual). `$gte` y `$lte` no son válidos: un filtro con ellos no coincide con nada y la respuesta es `204 No Content`, sin error.

En [Consultas](/es/docs/inceptiondb/v2/consultas/) están todos los operadores.

## 5. Lee un documento

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/ada" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

El documento viene en `document`:

```json
{"id":"ada","document":{"address":{"city":"London","zip":"N1"},"country":"GB","email":"ada@example.com","id":"ada","mrr":120,"name":"Ada Lovelace","plan":"pro","tags":["vip","beta"]},"source":{"type":"fullscan"}}
```

## 6. Actualiza un documento

Alan se pasa al plan `pro`. Un patch cambia solo los campos que envías:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}, "patch": {"plan": "pro", "mrr": 120}}'
```

```jsonl
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"pro","country":"GB","mrr":120,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

## 7. Borra un documento

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}}'
```

La respuesta contiene el documento borrado:

```jsonl
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"pro","country":"GB","mrr":120,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

`customers` tiene ahora dos documentos, Ada y Grace. Consérvala: las siguientes guías la usan.

## Siguientes pasos

- [Trabajando con datos](/es/docs/inceptiondb/v2/trabajando-con-datos/): inserciones masivas, actualizar y borrar muchos documentos.
- [Consultas](/es/docs/inceptiondb/v2/consultas/): todos los operadores de filtro, campos anidados, arrays y paginación.
- [Índices](/es/docs/inceptiondb/v2/indices/): búsquedas instantáneas por email, rangos por importe, búsqueda por palabras.
- [Colecciones](/es/docs/inceptiondb/v2/manejando-colecciones/): valores por defecto, tamaño y borrar una colección.
- La [consola](https://inceptiondb.hola.cloud/) muestra el código de cada acción que haces en ella. En [API keys](/es/docs/inceptiondb/v2/claves-api/) se explica dónde.
