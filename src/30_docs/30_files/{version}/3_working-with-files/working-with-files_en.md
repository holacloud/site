# Working With Files

This guide covers file operations in HolaCloud Files: uploading, downloading, listing, retrieving metadata, and deleting files.

## File Paths

File paths are specified in the URL after `/files/` or `/list/`. Paths can include directories using forward slashes:

```
/v1/buckets/{id}/files/images/logo.png
/v1/buckets/{id}/files/backups/db-2026-06-21.sql.gz
```

There is no need to pre-create directories -- the file path is treated as a flat key with slashes.

## Uploading a File

Use PUT to upload a file. Set the `Content-Type` header to the appropriate MIME type.

### Upload a Text File

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/notes/readme.txt" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: text/plain" \
  --data-binary @readme.txt
```

### Upload a Binary File (Image)

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/photo.jpg" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: image/jpeg" \
  --data-binary @photo.jpg
```

### Upload with Metadata

You can attach custom metadata to files using the `File-Metadata` header:

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/report.pdf" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/pdf" \
  -H "File-Metadata: {\"author\": \"Jane Doe\", \"department\": \"engineering\"}" \
  --data-binary @report.pdf
```

Expected response:

```json
{
  "path": "report.pdf",
  "size": 245760,
  "contentType": "application/pdf",
  "uploadedAt": "2026-06-21T12:00:00Z",
  "etag": "\"abc123def456\"",
  "metadata": {
    "author": "Jane Doe",
    "department": "engineering"
  }
}
```

## Downloading a File

Use GET to download a file by its path:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/photo.jpg" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -o photo.jpg
```

HTTP request:

```http
GET /v1/buckets/bkt_abc123/files/images/photo.jpg HTTP/1.1
Host: api.hola.cloud
Api-Key: YOUR_API_KEY
Api-Secret: YOUR_API_SECRET
```

The response includes the file content with appropriate `Content-Type` and `Content-Length` headers.

### Partial Downloads (Range Requests)

Files supports HTTP range requests for partial downloads:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/video.mp4" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Range: bytes=0-1023" \
  -o video_chunk.mp4
```

## Listing Files

List files in a bucket with an optional prefix filter. The `*` after `/list/` is required and indicates all files. Use a specific path to filter by prefix.

### List All Files

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

### List Files in a Directory

List files under the `images/` prefix:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/images/*" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Response:

```json
{
  "files": [
    {
      "path": "images/photo.jpg",
      "size": 245760,
      "contentType": "image/jpeg",
      "modifiedAt": "2026-06-21T12:00:00Z"
    },
    {
      "path": "images/banner.png",
      "size": 512000,
      "contentType": "image/png",
      "modifiedAt": "2026-06-21T11:30:00Z"
    }
  ],
  "prefix": "images/",
  "total": 2
}
```

### List with Pagination

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*?limit=10&offset=0" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

## Getting File Metadata

Use HEAD to retrieve file metadata without downloading the content:

```bash
curl -I "https://api.hola.cloud/v1/buckets/bkt_abc123/files/report.pdf" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Response headers:

```http
HTTP/1.1 200 OK
Content-Type: application/pdf
Content-Length: 245760
ETag: "abc123def456"
Last-Modified: Sun, 21 Jun 2026 12:00:00 GMT
File-Metadata: {"author": "Jane Doe", "department": "engineering"}
```

## Deleting a File

Delete a file by its path:

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123/files/notes/readme.txt" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Expected response: HTTP 204 No Content.

### Error: File Not Found

```json
{
  "error": {
    "code": "FILE_NOT_FOUND",
    "message": "File 'notes/readme.txt' not found in bucket 'bkt_abc123'."
  }
}
```

## Content-Type Handling

When uploading a file, you should set the `Content-Type` header to the correct MIME type for your file. If omitted, the service defaults to `application/octet-stream`.

Common MIME types:

| Extension | Content-Type |
|-----------|-------------|
| .txt | text/plain |
| .html | text/html |
| .css | text/css |
| .js | application/javascript |
| .json | application/json |
| .jpg / .jpeg | image/jpeg |
| .png | image/png |
| .gif | image/gif |
| .pdf | application/pdf |
| .zip | application/zip |
| .gz | application/gzip |
| .mp4 | video/mp4 |

## File Path Conventions

- File paths are case-sensitive
- Leading slashes are not allowed: `notes/readme.txt` (correct), `/notes/readme.txt` (incorrect)
- Paths can be up to 1024 characters
- Use forward slashes (`/`) as directory separators
- Special characters in path segments should be URL-encoded

## HTTP Status Code Summary

| Code | Description |
|------|-------------|
| 200 | Success (GET, HEAD, PUT) |
| 204 | No Content (DELETE) |
| 400 | Bad Request |
| 404 | Not Found -- file or bucket does not exist |
| 416 | Range Not Satisfiable |
| 500 | Internal Server Error |
