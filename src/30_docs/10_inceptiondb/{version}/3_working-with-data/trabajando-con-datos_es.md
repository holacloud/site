# Trabajando con datos

Las operaciones de documentos son acciones sobre una colección. Usa `Api-Key` y `Api-Secret` para acceso a la base de datos, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Insertar documentos

```http
POST /v1/databases/{databaseId}/collections/{collection}:insert
Content-Type: application/jsonl
Api-Key: tu-api-key
Api-Secret: tu-api-secret

{"id":"1","name":"Alice"}
{"id":"2","name":"Bob"}
```

La respuesta son los documentos insertados en JSONL.

```jsonl
{"id":"1","name":"Alice"}
{"id":"2","name":"Bob"}
```

## Buscar documentos

```http
POST /v1/databases/{databaseId}/collections/{collection}:find
Content-Type: application/json
Api-Key: tu-api-key
Api-Secret: tu-api-secret

{
  "filter": {
    "name": "Alice"
  }
}
```

La respuesta son los documentos encontrados en JSONL.

```jsonl
{"id":"1","name":"Alice"}
```

## Obtener un documento por ID

```http
GET /v1/databases/{databaseId}/collections/{collection}/documents/{documentId}
Api-Key: tu-api-key
Api-Secret: tu-api-secret
```

La respuesta es el documento JSON.

```json
{
  "id": "1",
  "name": "Alice"
}
```

## Modificar documentos

```http
POST /v1/databases/{databaseId}/collections/{collection}:patch
Content-Type: application/json
Api-Key: tu-api-key
Api-Secret: tu-api-secret

{
  "filter": {
    "id": "1"
  },
  "patch": {
    "name": "Alice Updated"
  }
}
```

La respuesta son los documentos modificados en JSONL.

```jsonl
{"id":"1","name":"Alice Updated"}
```

## Eliminar documentos

```http
POST /v1/databases/{databaseId}/collections/{collection}:remove
Content-Type: application/json
Api-Key: tu-api-key
Api-Secret: tu-api-secret

{
  "filter": {
    "id": "1"
  }
}
```

La respuesta son los documentos eliminados en JSONL.

```jsonl
{"id":"1","name":"Alice Updated"}
```
