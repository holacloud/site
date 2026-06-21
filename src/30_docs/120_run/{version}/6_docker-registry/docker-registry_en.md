# Docker Registry API

Run implements the Docker Registry v2 API specification for pushing and pulling container images.

## Authentication

Docker Registry endpoints require Basic Authentication using the DockerAuth mechanism. Credentials are typically configured via `docker login`.

```bash
docker login run.hola.cloud
```

## Endpoints

### GET /v2/

Check if the registry is running and supports API v2.

```bash
curl -X GET "https://api.hola.cloud/v2/" \
  -H "Authorization: Basic <base64-credentials>"
```

```http
HTTP/1.1 200 OK
Docker-Distribution-API-Version: registry/2.0
```

```json
{}
```

### HEAD /v2/:repository/blobs/:digest

Check if a blob exists.

```bash
curl -X HEAD "https://api.hola.cloud/v2/my-project/my-app/blobs/sha256:a1b2c3d4..." \
  -H "Authorization: Basic <base64-credentials>"
```

```http
HTTP/1.1 200 OK
Content-Length: 72450000
Docker-Content-Digest: sha256:a1b2c3d4...
```

### POST /v2/:repository/blobs/uploads/

Start a blob upload. Returns an upload UUID for subsequent operations.

```bash
curl -X POST "https://api.hola.cloud/v2/my-project/my-app/blobs/uploads/" \
  -H "Authorization: Basic <base64-credentials>"
```

```http
HTTP/1.1 202 Accepted
Location: /v2/my-project/my-app/blobs/uploads/abc-123-xyz
Range: 0-0
```

```json
{}
```

### PATCH /v2/:repository/blobs/uploads/:uuid

Upload a chunk of the blob.

```bash
curl -X PATCH "https://api.hola.cloud/v2/my-project/my-app/blobs/uploads/abc-123-xyz" \
  -H "Authorization: Basic <base64-credentials>" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @blob-chunk.bin
```

```http
HTTP/1.1 202 Accepted
Range: 0-1048575
```

### PUT /v2/:repository/blobs/uploads/:uuid

Finalize the blob upload by specifying the digest.

```bash
curl -X PUT "https://api.hola.cloud/v2/my-project/my-app/blobs/uploads/abc-123-xyz?digest=sha256:a1b2c3d4..." \
  -H "Authorization: Basic <base64-credentials>" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @final-chunk.bin
```

```http
HTTP/1.1 201 Created
Location: /v2/my-project/my-app/blobs/sha256:a1b2c3d4...
Docker-Content-Digest: sha256:a1b2c3d4...
```

### PUT /v2/:repository/manifests/:tag

Push a manifest for a given tag.

```bash
curl -X PUT "https://api.hola.cloud/v2/my-project/my-app/manifests/latest" \
  -H "Authorization: Basic <base64-credentials>" \
  -H "Content-Type: application/vnd.docker.distribution.manifest.v2+json" \
  -d '{
    "schemaVersion": 2,
    "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
    "config": {
      "mediaType": "application/vnd.docker.container.image.v1+json",
      "size": 7023,
      "digest": "sha256:config-digest..."
    },
    "layers": [
      {
        "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
        "size": 72450000,
        "digest": "sha256:layer-digest-1..."
      }
    ]
  }'
```

```http
HTTP/1.1 201 Created
Location: /v2/my-project/my-app/manifests/sha256:manifest-digest...
Docker-Content-Digest: sha256:manifest-digest...
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Request succeeded |
| 201 | Resource created successfully |
| 202 | Upload accepted |
| 400 | Invalid request |
| 401 | Authentication required |
| 404 | Blob or manifest not found |
| 416 | Range not satisfiable |
