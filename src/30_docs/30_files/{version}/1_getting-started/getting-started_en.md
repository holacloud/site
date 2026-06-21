# Getting Started with Files

This guide walks you through the basic operations of HolaCloud Files: creating a bucket, uploading a file, listing contents, downloading, and deleting.

## Prerequisites

Before you begin, you need:

- A HolaCloud account with an API key and API secret
- `curl` installed on your machine

Your API credentials are passed as headers in every request:

```
Api-Key: your-api-key
Api-Secret: your-api-secret
```

All requests use the base URL `https://api.hola.cloud`.

## Step 1: Create a Bucket

Buckets are containers for your files. Each bucket must have a globally unique name.

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-first-bucket"
  }'
```

Expected response:

```json
{
  "id": "bkt_abc123",
  "name": "my-first-bucket",
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## Step 2: Upload a File

Upload a text file to the bucket using PUT. The file path is specified in the URL after `/files/`.

```bash
echo "Hello, HolaCloud Files!" > hello.txt

curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/hello.txt" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: text/plain" \
  --data-binary @hello.txt
```

Expected response:

```json
{
  "path": "hello.txt",
  "size": 24,
  "contentType": "text/plain",
  "uploadedAt": "2026-06-21T10:01:00Z",
  "etag": "\"d41d8cd98f00b204e9800998ecf8427e\""
}
```

## Step 3: List Bucket Contents

List all files in the bucket with an optional prefix filter:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Expected response:

```json
{
  "files": [
    {
      "path": "hello.txt",
      "size": 24,
      "contentType": "text/plain",
      "modifiedAt": "2026-06-21T10:01:00Z"
    }
  ],
  "prefix": "",
  "total": 1
}
```

## Step 4: Download the File

Download the file using GET:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/hello.txt" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -o downloaded_hello.txt

cat downloaded_hello.txt
```

Output:

```
Hello, HolaCloud Files!
```

## Step 5: Delete the File

Delete the file using DELETE:

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123/files/hello.txt" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Expected response: HTTP 204 No Content.

## Step 6: Delete the Bucket

The bucket must be empty before you can delete it. Since we deleted the file, this will succeed:

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Expected response: HTTP 204 No Content.

## Summary

You have successfully created a bucket, uploaded a file, listed bucket contents, downloaded the file, and cleaned up. You are now ready to integrate HolaCloud Files into your applications.
