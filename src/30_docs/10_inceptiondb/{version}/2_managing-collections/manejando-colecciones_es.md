# Colecciones

Una colección es un grupo de documentos JSON con nombre dentro de una base de datos, como `customers` u `orders`. Los documentos de una misma colección pueden tener campos distintos: no hay esquema que declarar.

Los ejemplos usan las variables de entorno de [Primeros pasos](/es/docs/inceptiondb/v2/primeros-pasos/).

## Crear una colección

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

| Campo | Significado |
|---|---|
| `name` | Nombre de la colección. Se usa en todas sus URLs. |
| `total` | Número de documentos. |
| `indexes` | Número de índices. |
| `defaults` | Valores que se añaden a los documentos que no los tienen (ver más abajo). |

## Listar las colecciones

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}]
```

## Obtener una colección

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

## Valores por defecto

Los valores por defecto rellenan los campos que un documento insertado no tiene. Un campo enviado como `null` también recibe su valor por defecto. Además de valores fijos, hay tres generados:

| Valor | Genera |
|---|---|
| `"uuid()"` | Un ID único aleatorio, como `"35ec7d8f-96f4-4cdf-a09e-463c720e7c8c"`. |
| `"auto()"` | Un número creciente: 1, 2, 3... Nunca se repite, aunque puede saltarse números. |
| `"unixnano()"` | El momento de la inserción, en nanosegundos desde 1970. |

Dale a cada cliente nuevo un ID, el plan `free` y su fecha de creación:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:setDefaults" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"id": "uuid()", "plan": "free", "created": "unixnano()"}'
```

```json
{"created":"unixnano()","id":"uuid()","plan":"free"}
```

Ahora inserta un cliente solo con nombre y email:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "Margaret Hamilton", "email": "margaret@example.com"}'
```

```jsonl
{"created":1793164963789350717,"email":"margaret@example.com","id":"35ec7d8f-96f4-4cdf-a09e-463c720e7c8c","name":"Margaret Hamilton","plan":"free"}
```

Los valores por defecto se aplican a los documentos que se insertan a partir de ahora; los que ya existen no cambian. `setDefaults` añade los campos que envías a los valores actuales y sustituye los que ya existen. Para quitar un valor por defecto, envíalo como `null`: `{"plan": null}`.

## Tamaño

Cuánta memoria ocupan los documentos, el tamaño en disco y, para cada índice, su memoria como `index.<nombre>`, en bytes:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:size" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"disk":265,"memory":203}
```

## Borrar una colección

Esto borra la colección con todos sus documentos e índices. No se puede deshacer.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

La respuesta es `204 No Content`.
