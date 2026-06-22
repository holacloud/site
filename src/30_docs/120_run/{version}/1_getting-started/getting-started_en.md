# Getting Started

Run exposes a push-oriented Docker Registry v2 subset and a Console API for repository state, starts, stops, rollbacks, environment variables, and volumes.

## Authentication

Registry endpoints use Docker Basic Auth credentials configured by `docker login`.

```bash
docker login run.hola.cloud
```

## Building and Pushing an Image

Create a simple application with a `Dockerfile`:

```dockerfile
FROM alpine:latest
COPY app.sh /app.sh
RUN chmod +x /app.sh
CMD ["/app.sh"]
```

Build and tag the image:

```bash
docker build -t run.hola.cloud/my-project/my-app:v1 .
```

Push to Run:

```bash
docker push run.hola.cloud/my-project/my-app:v1
```

## Reading Console Data

Query the current Console data for a repository:

```bash
curl "https://api.hola.cloud/api/console?repository=my-project/my-app"
```

## Starting a Repository

Start a repository from a pushed tag or digest:

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "reference": "v1"
  }'
```

You can also use a digest:

```json
{
  "repository": "my-project/my-app",
  "digest": "sha256:a1b2c3d4..."
}
```

## Stopping a Repository

```bash
curl -X POST "https://api.hola.cloud/api/console/stop" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app"}'
```

## Rolling Back

Rollback uses the same repository plus `reference` or `digest` shape:

```bash
curl -X POST "https://api.hola.cloud/api/console/rollback" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "reference": "v0"
  }'
```

## Managing Environment Variables

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "env": [
      {"key": "LOG_LEVEL", "desired_value": "debug"},
      {"key": "DATABASE_URL", "desired_value": "postgres://..."}
    ]
  }'
```

## Managing Volumes

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "volumes": [
      {"name": "my-data", "target": "/data", "mode": "rw"}
    ]
  }'
```

## Complete Workflow Example

```bash
docker login run.hola.cloud
docker build -t run.hola.cloud/my-project/my-app:v1 .
docker push run.hola.cloud/my-project/my-app:v1
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "reference": "v1"}'
curl "https://api.hola.cloud/api/console?repository=my-project/my-app"
```
