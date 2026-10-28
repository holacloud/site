# Primeros pasos

En cinco minutos vas a crear una colección, guardar algunos clientes, consultarlos, cambiar uno y borrar otro. Todos los comandos se pueden copiar y pegar tal cual.

## 1. Consigue tus credenciales

1. Abre la [consola de InceptionDB](https://inceptiondb.hola.cloud/) e inicia sesión.
2. Crea una base de datos (por ejemplo `my-app`) o elige una que ya tengas.
3. Abre **API keys & settings** y crea una API key. El secreto se muestra **una sola vez**: cópialo ahora.

La consola te da los tres valores listos para pegar en un terminal. Expórtalos como variables de entorno, así nunca acaban en tu código:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

Todos los ejemplos de esta documentación leen estas variables, también las versiones en Go, Python, Node.js, PHP y Java.

## 2. Crea una colección

Una colección agrupa documentos. Crea una llamada `customers`:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "customers"}'
```

Respuesta (`201 Created`):

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

`defaults` indica que un documento sin `id` recibe uno aleatorio.

## 3. Inserta documentos

Envía un documento JSON por línea (un formato llamado JSON Lines):

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  --data-binary '{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}'
```

La respuesta (`201 Created`) trae los documentos guardados, uno por línea:

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
  -d '{"filter": {"country": "GB"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

Los filtros también comparan números, miran dentro de arrays y objetos anidados, y combinan condiciones. Los clientes que pagan más de 100 al mes:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"mrr": {"$gt": 100}}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

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
  -d '{"filter": {"id": "grace"}}'
```

La respuesta trae el documento borrado:

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## Siguientes pasos

- [Consultas](/es/docs/inceptiondb/v2/consultas/): todos los operadores de filtro, paginación y exportaciones.
- [Índices](/es/docs/inceptiondb/v2/indices/): búsquedas instantáneas por email, rangos por importe y búsqueda por palabras.
- [Trabajando con datos](/es/docs/inceptiondb/v2/trabajando-con-datos/): inserciones masivas, actualizar muchos documentos y errores.
- La [consola](https://inceptiondb.hola.cloud/) muestra el código de cada acción que haces en ella: pulsa las bolas de la derecha de la pantalla.
