# Docker Registry

Run 实现 Docker Registry v2 的面向 push 子集，用于上传 blob 和 manifest。它不是完整 pull-capable registry API。

## 认证

```bash
docker login run.hola.cloud
```

## 支持的 endpoints

- `GET /v2/`
- `HEAD /v2/*`
- `POST /v2/:repository/blobs/uploads/`
- `PATCH /v2/:repository/blobs/uploads/:uuid`
- `PUT /v2/:repository/blobs/uploads/:uuid`
- `PUT /v2/:repository/manifests/:reference`

Manifest 示例：

```bash
curl -X PUT "https://api.hola.cloud/v2/my-project/my-app/manifests/latest" \
  -H "Authorization: Basic <base64-credentials>" \
  -H "Content-Type: application/vnd.docker.distribution.manifest.v2+json" \
  --data-binary @manifest.json
```
