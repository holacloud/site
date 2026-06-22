# Guía rápida de InceptionDB

Esta guía crea una colección, inserta tres documentos JSON y consulta documentos filtrados.

Los endpoints de acceso a la base de datos aceptan `Api-Key` y `Api-Secret`. También se puede usar un token Glue de owner cuando el propietario de la base de datos está permitido.

## Paso 1: Crear una colección

```bash
curl -X POST "https://api.hola.cloud/v1/databases/tu-id-de-base-de-datos/collections" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-coleccion"
  }'
```

Respuesta esperada:

```json
{
  "name": "mi-coleccion",
  "total": 0,
  "indexes": 0,
  "defaults": {
    "id": "uuid()"
  }
}
```

## Paso 2: Insertar documentos

Usa el endpoint de acción de colección y envía JSONL: un documento JSON por línea.

```bash
curl -X POST "https://api.hola.cloud/v1/databases/tu-id-de-base-de-datos/collections/mi-coleccion:insert" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret" \
  -H "Content-Type: application/jsonl" \
  --data-binary '{"name":"Elemento 1","value":100}
{"name":"Elemento 2","value":200}
{"name":"Elemento 3","value":300}'
```

Respuesta esperada:

```jsonl
{"id":"document-id-1","name":"Elemento 1","value":100}
{"id":"document-id-2","name":"Elemento 2","value":200}
{"id":"document-id-3","name":"Elemento 3","value":300}
```

## Paso 3: Buscar documentos filtrados

```bash
curl -X POST "https://api.hola.cloud/v1/databases/tu-id-de-base-de-datos/collections/mi-coleccion:find" \
  -H "Api-Key: tu-api-key" \
  -H "Api-Secret: tu-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "value": { "$gte": 200 }
    }
  }'
```

Respuesta esperada:

```jsonl
{"id":"document-id-2","name":"Elemento 2","value":200}
{"id":"document-id-3","name":"Elemento 3","value":300}
```
