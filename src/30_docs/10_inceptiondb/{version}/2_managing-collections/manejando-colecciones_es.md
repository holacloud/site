# Manejando colecciones

Los endpoints de colección usan credenciales de acceso a la base de datos: `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Crear una colección

```bash
curl -X POST "https://api.hola.cloud/v1/databases/{databaseId}/collections" \
  -H "Api-Key: {api_key}" \
  -H "Api-Secret: {api_secret}" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-coleccion"
  }'
```

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

## Listar colecciones

```bash
curl "https://api.hola.cloud/v1/databases/{databaseId}/collections" \
  -H "Api-Key: {api_key}" \
  -H "Api-Secret: {api_secret}"
```

```json
[
  {
    "name": "mi-coleccion",
    "total": 3,
    "indexes": 0,
    "defaults": {
      "id": "uuid()"
    }
  }
]
```

## Insertar documentos

```bash
curl -X POST "https://api.hola.cloud/v1/databases/{databaseId}/collections/mi-coleccion:insert" \
  -H "Api-Key: {api_key}" \
  -H "Api-Secret: {api_secret}" \
  -H "Content-Type: application/jsonl" \
  --data-binary '{"id":"1","name":"Alfonso"}
{"id":"2","name":"Gerardo"}'
```

La respuesta es JSONL con los documentos insertados.

```jsonl
{"id":"1","name":"Alfonso"}
{"id":"2","name":"Gerardo"}
```
