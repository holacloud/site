# Working With Files

Files are addressed by the path after `/files/`. Listing uses the path after `/list/` as a prefix filter.

## Upload a File

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: text/plain" \
  --data-binary @readme.txt
```

Response:

```json
{
  "id": "file-9f0b7b3c-1d2e-4a5f-8b9c-0123456789ab",
  "uuid": "9f0b7b3c-1d2e-4a5f-8b9c-0123456789ab",
  "created_timestamp": 1782045660000000000,
  "updated_timestamp": 1782045660000000000,
  "owners": ["user-123"],
  "status": "available",
  "size": 24576,
  "name": "notes/readme.txt",
  "bucket": "bucket-550e8400-e29b-41d4-a716-446655440000",
  "hash_md5": "example-md5",
  "hash_sha256": "example-sha256",
  "mime_type": "text/plain; charset=utf-8"
}
```

## Download a File

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -o readme.txt
```

The response body is the stored file content. `Content-Type` is set from the stored `mime_type`.

## Get File JSON Metadata

Use `?metadata` on the download endpoint to return the file object as JSON.

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt?metadata" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## List Files

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/list/notes/" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

Response:

```json
[
  {
    "id": "file-9f0b7b3c-1d2e-4a5f-8b9c-0123456789ab",
    "name": "notes/readme.txt",
    "bucket": "bucket-550e8400-e29b-41d4-a716-446655440000",
    "created_timestamp": 1782045660000000000,
    "updated_timestamp": 1782045660000000000,
    "size": 24576,
    "mime_type": "text/plain; charset=utf-8",
    "hash_md5": "example-md5",
    "hash_sha256": "example-sha256",
    "status": "available",
    "owners": ["user-123"]
  }
]
```

## HEAD File

```bash
curl -I "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

The handler sets `Last-Modified` and `Content-Length`.

## Delete a File

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notes/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```
