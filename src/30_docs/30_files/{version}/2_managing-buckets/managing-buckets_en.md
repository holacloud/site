# Managing Buckets

Buckets are the fundamental containers in HolaCloud Files. This guide covers all bucket operations: creating, listing, inspecting, modifying, and deleting.

## Bucket Naming Rules

When creating a bucket, the name must follow these rules:

- Must be between 3 and 63 characters long
- Can only contain lowercase letters, numbers, and hyphens (-)
- Must start and end with a letter or number
- Must not be formatted like an IP address
- Must be globally unique across all HolaCloud customers

## Creating a Bucket

Create a new bucket with a POST request:

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-assets-bucket",
    "metadata": {
      "environment": "production",
      "project": "my-app"
    }
  }'
```

HTTP request equivalent:

```http
POST /v1/buckets HTTP/1.1
Host: api.hola.cloud
Api-Key: YOUR_API_KEY
Api-Secret: YOUR_API_SECRET
Content-Type: application/json

{
  "name": "my-assets-bucket",
  "metadata": {
    "environment": "production",
    "project": "my-app"
  }
}
```

Expected response (HTTP 201):

```json
{
  "id": "bkt_xyz789",
  "name": "my-assets-bucket",
  "metadata": {
    "environment": "production",
    "project": "my-app"
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## Listing Buckets

List all buckets in your account:

```bash
curl "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Expected response:

```json
[
  {
    "id": "bkt_abc123",
    "name": "my-first-bucket",
    "createdAt": "2026-06-21T09:00:00Z",
    "size": 1024,
    "fileCount": 3
  },
  {
    "id": "bkt_xyz789",
    "name": "my-assets-bucket",
    "createdAt": "2026-06-21T10:00:00Z",
    "size": 0,
    "fileCount": 0
  }
]
```

## Getting Bucket Details

Retrieve detailed information about a specific bucket by its ID:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_xyz789" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Expected response:

```json
{
  "id": "bkt_xyz789",
  "name": "my-assets-bucket",
  "metadata": {
    "environment": "production",
    "project": "my-app"
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "modifiedAt": "2026-06-21T10:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## Modifying a Bucket

Update bucket metadata using PATCH:

```bash
curl -X PATCH "https://api.hola.cloud/v1/buckets/bkt_xyz789" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "metadata": {
      "environment": "staging"
    }
  }'
```

Expected response:

```json
{
  "id": "bkt_xyz789",
  "name": "my-assets-bucket",
  "metadata": {
    "environment": "staging",
    "project": "my-app"
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "modifiedAt": "2026-06-21T11:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## Deleting a Bucket

A bucket must be **empty** (contain no files) before it can be deleted.

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_xyz789" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Expected response: HTTP 204 No Content.

### Error: Bucket Not Empty

If you attempt to delete a bucket that still contains files, you will receive:

```json
{
  "error": {
    "code": "BUCKET_NOT_EMPTY",
    "message": "Cannot delete bucket 'my-assets-bucket': the bucket contains 3 file(s). Please delete all files before deleting the bucket."
  }
}
```

## Error Handling

### Duplicate Bucket Name

```json
{
  "error": {
    "code": "BUCKET_ALREADY_EXISTS",
    "message": "A bucket with the name 'my-assets-bucket' already exists."
  }
}
```

### Invalid Bucket Name

```json
{
  "error": {
    "code": "INVALID_BUCKET_NAME",
    "message": "Bucket name must be between 3 and 63 characters, contain only lowercase letters, numbers, and hyphens, and start/end with a letter or number."
  }
}
```

### Bucket Not Found

```json
{
  "error": {
    "code": "BUCKET_NOT_FOUND",
    "message": "Bucket 'bkt_nonexistent' not found."
  }
}
```

## HTTP Status Code Summary

| Code | Description |
|------|-------------|
| 200 | Success (GET, PATCH) |
| 201 | Created (POST) |
| 204 | No Content (DELETE) |
| 400 | Bad Request -- invalid input |
| 404 | Not Found -- bucket does not exist |
| 409 | Conflict -- bucket already exists or not empty |
| 500 | Internal Server Error |
