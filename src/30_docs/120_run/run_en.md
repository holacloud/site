# Run

Run is HolaCloud's container image and console-control service. It exposes a push-oriented subset of the Docker Registry v2 API and a small Console API for inspecting repositories and changing runtime configuration.

## Features

### Push-Oriented Registry v2 Subset

Run supports the registry endpoints needed to push image blobs and manifests under `/v2`. It is not a full Docker Registry implementation and should not be documented as a general pull registry.

```bash
docker login run.hola.cloud
docker build -t run.hola.cloud/my-project/my-app:latest .
docker push run.hola.cloud/my-project/my-app:latest
```

### Console Management API

The Console API works with repositories and image references or digests:

- `GET /version`
- `GET /api/console?repository=`
- `POST /api/console/start`
- `POST /api/console/stop`
- `POST /api/console/rollback`
- `PUT /api/console/env`
- `PUT /api/console/volumes`

There is no `/v1/run/deploy`, `/api/console/run`, push/exec API, or `container_id` workflow.

### Environment Variables and Volumes

Environment and volume configuration is saved by repository:

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "env": [
      {"key": "LOG_LEVEL", "desired_value": "debug"}
    ]
  }'
```

## Getting Started

1. Authenticate with the registry:

```bash
docker login run.hola.cloud
```

2. Build and push your image:

```bash
docker build -t run.hola.cloud/my-project/my-app:latest .
docker push run.hola.cloud/my-project/my-app:latest
```

3. Start the repository at a pushed reference:

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "reference": "latest"
  }'
```
