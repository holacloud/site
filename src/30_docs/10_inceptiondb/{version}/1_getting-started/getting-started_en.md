
# Quick Start Guide with InceptionDB

This guide will help you get started with InceptionDB by creating a collection, inserting three elements, and querying filtered elements.

## Step 1: Create a Collection

To create a collection in InceptionDB, you need to define the name of the collection. Here is an example of how to do it using curl in bash:

```bash
curl -X POST "https://example.com/v1/databases/your-database-id/collections" \
-H "Api-Key: your-api-key" \
-H "Api-Secret: your-api-secret" \
-d '{
      "name": "my-collection"
    }'
```

Expected Response

```json
{
    "defaults": {
        "id": "uuid()"
    },
    "indexes": 0,
    "name": "my-collection",
    "total": 0
}
```

## Step 2: Insert Elements

Once you have created the collection, you can insert elements into it. Here is how to insert three elements into the newly created collection:

```bash
curl -X POST "https://example.com/v1/databases/your-database-id/collections/my-collection/documents" \
-H "Api-Key: your-api-key" \
-H "Api-Secret: your-api-secret" \
-d '{
      "name": "Element 1",
      "value": 100
    }
    {
      "name": "Element 2",
      "value": 200
    }
    {
      "name": "Element 3",
      "value": 300
    }'
```

## Step 3: List Filtered Elements

To list elements in the collection with a specific filter, you can use the following curl request:

```bash
curl -X POST "https://example.com/v1/databases/your-database-id/collections/my-collection/find" \
-H "Api-Key: your-api-key" \
-H "Api-Secret: your-api-secret" \
-d '{
      "filter": {
        "value": { "$gte": 200 }
      }
    }'
```

Expected Response

```json
{
    "documents": [
        {
            "id": "document-id-2",
            "name": "Element 2",
            "value": 200
        },
        {
            "id": "document-id-3",
            "name": "Element 3",
            "value": 300
        }
    ]
}
```

## Summary

By following these steps, you have created a collection in InceptionDB, inserted three elements, and queried elements with a filter. This provides you with a solid foundation to start working with InceptionDB and take advantage of its capabilities to efficiently manage your data.
