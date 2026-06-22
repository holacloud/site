# Docker Registry

Run implements the push-oriented subset of Docker Registry v2 needed by Docker clients to upload image blobs and manifests. It is not a full pull-capable registry API.

## Authentication

Registry endpoints require Basic Authentication configured through `docker login`.

```bash
docker login run.hola.cloud
```

## Supported Endpoints

### GET /v2/

Check that the registry API is available.

```bash
curl -X GET "https://api.hola.cloud/v2/" \
  -H "Authorization: Basic <base64-credentials>"
```

### HEAD /v2/*

Check blob or manifest existence during a push workflow.

### POST /v2/:repository/blobs/uploads/

Start a blob upload and receive an upload UUID.

```bash
curl -X POST "https://api.hola.cloud/v2/my-project/my-app/blobs/uploads/" \
  -H "Authorization: Basic <base64-credentials>"
```

### PATCH /v2/:repository/blobs/uploads/:uuid

Upload blob data.

```bash
curl -X PATCH "https://api.hola.cloud/v2/my-project/my-app/blobs/uploads/abc-123" \
  -H "Authorization: Basic <base64-credentials>" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @layer.tar
```

### PUT /v2/:repository/blobs/uploads/:uuid

Finalize a blob upload with its digest.

```bash
curl -X PUT "https://api.hola.cloud/v2/my-project/my-app/blobs/uploads/abc-123?digest=sha256:a1b2c3d4..." \
  -H "Authorization: Basic <base64-credentials>" \
  -H "Content-Type: application/octet-stream"
```

### PUT /v2/:repository/manifests/:reference

Push a manifest for a tag or reference.

```bash
curl -X PUT "https://api.hola.cloud/v2/my-project/my-app/manifests/latest" \
  -H "Authorization: Basic <base64-credentials>" \
  -H "Content-Type: application/vnd.docker.distribution.manifest.v2+json" \
  --data-binary @manifest.json
```

## Not Supported

Do not rely on Run as a full Registry v2 pull implementation. The documented registry surface is the push-oriented subset above.
