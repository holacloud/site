# Docker Registry API

Run 实现了 Docker Registry v2 API 规范，用于推送和拉取容器镜像。

## 身份验证

Docker Registry 端点需要使用 DockerAuth 机制进行基本身份验证。通常通过 `docker login` 配置凭据。

```bash
docker login run.hola.cloud
```

## 端点

### GET /v2/

检查注册表是否正在运行并支持 API v2。

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

检查 blob 是否存在。

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

开始 blob 上传。返回用于后续操作的上传 UUID。

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

上传 blob 的一个块。

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

通过指定 digest 完成 blob 上传。

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

为指定标签推送清单。

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

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 请求成功 |
| 201 | 资源创建成功 |
| 202 | 上传已接受 |
| 400 | 请求无效 |
| 401 | 需要身份验证 |
| 404 | blob 或清单未找到 |
| 416 | 范围无法满足 |
