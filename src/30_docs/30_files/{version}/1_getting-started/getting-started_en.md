# Getting Started

This guide walks through creating a bucket, uploading a file, listing files, downloading, and deleting.

## Prerequisites

- `curl` installed locally
- A valid `X-Glue-Authentication` header for the authenticated user

```http
X-Glue-Authentication: {"user":{"id":"user-123"}}
```

All requests use `https://api.hola.cloud`.

## Step 1: Create a Bucket

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-first-bucket",
    "description": "First test bucket"
  }'
```

Expected response:

```json
{
  "id": "bucket-550e8400-e29b-41d4-a716-446655440000",
  "project_id": "",
  "created_timestamp": 1782045600000000000,
  "owners": ["user-123"],
  "name": "my-first-bucket",
  "description": "First test bucket"
}
```

## Step 2: Upload a File

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/hello.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: text/plain" \
  --data-binary "Hello, HolaCloud Files!"
```

Expected response:

```json
{
  "id": "file-9f0b7b3c-1d2e-4a5f-8b9c-0123456789ab",
  "uuid": "9f0b7b3c-1d2e-4a5f-8b9c-0123456789ab",
  "created_timestamp": 1782045660000000000,
  "updated_timestamp": 1782045660000000000,
  "owners": ["user-123"],
  "status": "available",
  "size": 23,
  "name": "hello.txt",
  "bucket": "bucket-550e8400-e29b-41d4-a716-446655440000",
  "hash_md5": "example-md5",
  "hash_sha256": "example-sha256",
  "mime_type": "text/plain; charset=utf-8"
}
```

## Step 3: List Files

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/list/*" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

Expected response:

```json
[
  {
    "id": "file-9f0b7b3c-1d2e-4a5f-8b9c-0123456789ab",
    "name": "hello.txt",
    "bucket": "bucket-550e8400-e29b-41d4-a716-446655440000",
    "created_timestamp": 1782045660000000000,
    "updated_timestamp": 1782045660000000000,
    "size": 23,
    "mime_type": "text/plain; charset=utf-8",
    "hash_md5": "example-md5",
    "hash_sha256": "example-sha256",
    "status": "available",
    "owners": ["user-123"]
  }
]
```

## Step 4: Download the File

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/hello.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Step 5: Delete the File

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/hello.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Step 6: Delete the Bucket

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```
