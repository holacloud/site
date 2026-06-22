# Docker Registry

Run implementa el subconjunto Docker Registry v2 orientado a push necesario para subir blobs y manifiestos. No es una API registry completa para pull.

## Autenticación

```bash
docker login run.hola.cloud
```

## Endpoints soportados

- `GET /v2/`
- `HEAD /v2/*`
- `POST /v2/:repository/blobs/uploads/`
- `PATCH /v2/:repository/blobs/uploads/:uuid`
- `PUT /v2/:repository/blobs/uploads/:uuid`
- `PUT /v2/:repository/manifests/:reference`

Ejemplo de manifiesto:

```bash
curl -X PUT "https://api.hola.cloud/v2/my-project/my-app/manifests/latest" \
  -H "Authorization: Basic <base64-credentials>" \
  -H "Content-Type: application/vnd.docker.distribution.manifest.v2+json" \
  --data-binary @manifest.json
```
