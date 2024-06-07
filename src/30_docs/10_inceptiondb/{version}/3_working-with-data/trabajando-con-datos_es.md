
# Trabajando con datos

Este documento explica cómo se insertan, listan, modifican y eliminan documentos en una colección. También se incluye cómo se pueden filtrar los datos al listar.

## Insertar Documentos

### Insertar un documento

Para insertar un solo documento en una colección, se utiliza la siguiente petición HTTP:

```http
POST /databases/{database}/collections/{collection}/documents
Content-Type: application/json
X-API-Key: your-api-key

{
  "document": {
    "key1": "value1",
    "key2": "value2"
  }
}
```

### Insertar múltiples documentos

Para insertar múltiples documentos en una colección, se utiliza la siguiente petición HTTP:

```http
POST /databases/{database}/collections/{collection}/documents/bulk
Content-Type: application/json
X-API-Key: your-api-key

{
  "documents": [
    {
      "key1": "value1",
      "key2": "value2"
    },
    {
      "key1": "value3",
      "key2": "value4"
    }
  ]
}
```

## Listar Documentos

### Listar todos los documentos (Full Scan)

Para listar todos los documentos en una colección, se utiliza la siguiente petición HTTP:

```http
GET /databases/{database}/collections/{collection}/documents
X-API-Key: your-api-key
```

### Listar documentos con un índice único

Para listar documentos utilizando un índice único, se utiliza la siguiente petición HTTP:

```http
GET /databases/{database}/collections/{collection}/documents?filter={"key":"value"}
X-API-Key: your-api-key
```

### Listar documentos con un límite

Para listar documentos con un límite específico, se utiliza la siguiente petición HTTP:

```http
GET /databases/{database}/collections/{collection}/documents?limit=10
X-API-Key: your-api-key
```

## Modificar Documentos

### Modificar documentos por un escaneo completo

Para modificar documentos mediante un escaneo completo, se utiliza la siguiente petición HTTP:

```http
PATCH /databases/{database}/collections/{collection}/documents
Content-Type: application/json
X-API-Key: your-api-key

{
  "filter": {"key": "value"},
  "update": {"$set": {"key": "new_value"}}
}
```

### Modificar documentos por índice

Para modificar documentos utilizando un índice, se utiliza la siguiente petición HTTP:

```http
PATCH /databases/{database}/collections/{collection}/documents?filter={"key":"value"}
Content-Type: application/json
X-API-Key: your-api-key

{
  "update": {"$set": {"key": "new_value"}}
}
```

## Eliminar Documentos

### Eliminar documentos por índice

Para eliminar documentos utilizando un índice, se utiliza la siguiente petición HTTP:

```http
DELETE /databases/{database}/collections/{collection}/documents?filter={"key":"value"}
X-API-Key: your-api-key
```

### Eliminar documentos por escaneo completo

Para eliminar documentos mediante un escaneo completo, se utiliza la siguiente petición HTTP:

```http
DELETE /databases/{database}/collections/{collection}/documents
Content-Type: application/json
X-API-Key: your-api-key

{
  "filter": {"key": "value"}
}
```
