
# Working with Data

This document explains how to insert, list, modify, and delete documents in a collection. It also includes how to filter data when listing.

## Inserting Documents

### Insert a Single Document

To insert a single document into a collection, use the following HTTP request:

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

### Insert Multiple Documents

To insert multiple documents into a collection, use the following HTTP request:

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

## Listing Documents

### List All Documents (Full Scan)

To list all documents in a collection, use the following HTTP request:

```http
GET /databases/{database}/collections/{collection}/documents
X-API-Key: your-api-key
```

### List Documents by Unique Index

To list documents using a unique index, use the following HTTP request:

```http
GET /databases/{database}/collections/{collection}/documents?filter={"key":"value"}
X-API-Key: your-api-key
```

### List Documents with a Limit

To list documents with a specific limit, use the following HTTP request:

```http
GET /databases/{database}/collections/{collection}/documents?limit=10
X-API-Key: your-api-key
```

## Modifying Documents

### Modify Documents by Full Scan

To modify documents by a full scan, use the following HTTP request:

```http
PATCH /databases/{database}/collections/{collection}/documents
Content-Type: application/json
X-API-Key: your-api-key

{
  "filter": {"key": "value"},
  "update": {"$set": {"key": "new_value"}}
}
```

### Modify Documents by Index

To modify documents using an index, use the following HTTP request:

```http
PATCH /databases/{database}/collections/{collection}/documents?filter={"key":"value"}
Content-Type: application/json
X-API-Key: your-api-key

{
  "update": {"$set": {"key": "new_value"}}
}
```

## Deleting Documents

### Delete Documents by Index

To delete documents using an index, use the following HTTP request:

```http
DELETE /databases/{database}/collections/{collection}/documents?filter={"key":"value"}
X-API-Key: your-api-key
```

### Delete Documents by Full Scan

To delete documents by a full scan, use the following HTTP request:

```http
DELETE /databases/{database}/collections/{collection}/documents
Content-Type: application/json
X-API-Key: your-api-key

{
  "filter": {"key": "value"}
}
```
