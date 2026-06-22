# Docker Registro

Run implementa la especificación de Docker Registry v2 para subir y descargar imágenes de contenedores.

## Autenticación

Los endpoints de Docker Registry requieren Autenticación Básica mediante el mecanismo DockerAuth. Las credenciales se configuran típicamente mediante `docker login`.

```bash
docker login run.hola.cloud
```

## Endpoints

### GET /v2/

Verifica si el registro está funcionando y admite la API v2.

```bash
curl -X GET "https://api.hola.cloud/v2/" \
  -H "Authorization: Basic <credenciales-base64>"
```

```http
HTTP/1.1 200 OK
Docker-Distribution-API-Version: registry/2.0
```

```json
{}
```

### HEAD /v2/:repository/blobs/:digest

Verifica si un blob existe.

```bash
curl -X HEAD "https://api.hola.cloud/v2/my-project/my-app/blobs/sha256:a1b2c3d4..." \
  -H "Authorization: Basic <credenciales-base64>"
```

```http
HTTP/1.1 200 OK
Content-Length: 72450000
Docker-Content-Digest: sha256:a1b2c3d4...
```

### POST /v2/:repository/blobs/uploads/

Inicia una subida de blob. Devuelve un UUID de subida para operaciones posteriores.

```bash
curl -X POST "https://api.hola.cloud/v2/my-project/my-app/blobs/uploads/" \
  -H "Authorization: Basic <credenciales-base64>"
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

Sube un fragmento del blob.

```bash
curl -X PATCH "https://api.hola.cloud/v2/my-project/my-app/blobs/uploads/abc-123-xyz" \
  -H "Authorization: Basic <credenciales-base64>" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @blob-chunk.bin
```

```http
HTTP/1.1 202 Accepted
Range: 0-1048575
```

### PUT /v2/:repository/blobs/uploads/:uuid

Finaliza la subida del blob especificando el digest.

```bash
curl -X PUT "https://api.hola.cloud/v2/my-project/my-app/blobs/uploads/abc-123-xyz?digest=sha256:a1b2c3d4..." \
  -H "Authorization: Basic <credenciales-base64>" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @final-chunk.bin
```

```http
HTTP/1.1 201 Created
Location: /v2/my-project/my-app/blobs/sha256:a1b2c3d4...
Docker-Content-Digest: sha256:a1b2c3d4...
```

### PUT /v2/:repository/manifests/:tag

Sube un manifiesto para una etiqueta determinada.

```bash
curl -X PUT "https://api.hola.cloud/v2/my-project/my-app/manifests/latest" \
  -H "Authorization: Basic <credenciales-base64>" \
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

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Solicitud exitosa |
| 201 | Recurso creado correctamente |
| 202 | Subida aceptada |
| 400 | Solicitud inválida |
| 401 | Autenticación requerida |
| 404 | Blob o manifiesto no encontrado |
| 416 | Rango no satisfacible |
