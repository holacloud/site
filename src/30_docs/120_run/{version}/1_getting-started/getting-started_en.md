# Getting Started

Run is a container execution service with a built-in Docker Registry v2 compatible API. It allows you to deploy, run, and manage container images with environment variables and volumes.

## Authentication

Run uses **DockerAuth** with Basic Auth for registry authentication. Use your HolaCloud API key and secret as the username and password.

### Logging into the Registry

```bash
docker login https://registry.hola.cloud
```

You will be prompted for your HolaCloud API key (username) and API secret (password).

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
docker build -t registry.hola.cloud/my-project/my-app:v1 .
```

Push to the registry:

```bash
docker push registry.hola.cloud/my-project/my-app:v1
```

## Starting a Container

Start a container from a pushed image via the Console API:

```bash
curl -X POST "https://api.hola.cloud/api/console/run" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "registry.hola.cloud/my-project/my-app:v1",
    "env": {
      "LOG_LEVEL": "info",
      "DATABASE_URL": "postgres://..."
    },
    "volumes": [
      {
        "container_path": "/data",
        "volume_name": "my-data"
      }
    ]
  }'
```

Expected response:

```json
{
  "id": "c8a9f3b2-1234-5678-9abc-def012345678",
  "image": "registry.hola.cloud/my-project/my-app:v1",
  "status": "running",
  "created_at": "2025-06-21T12:00:00Z"
}
```

## Stopping a Running Container

Stop a container by its ID:

```bash
curl -X DELETE "https://api.hola.cloud/api/console/run/c8a9f3b2-1234-5678-9abc-def012345678" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

## Rolling Back to a Previous Image Version

Deploy a previous image version by specifying a different tag:

```bash
curl -X PUT "https://api.hola.cloud/api/console/run/c8a9f3b2-1234-5678-9abc-def012345678" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "registry.hola.cloud/my-project/my-app:v0"
  }'
```

## Managing Environment Variables

Update environment variables for a running container:

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "env": {
      "LOG_LEVEL": "debug",
      "DATABASE_URL": "postgres://..."
    }
  }'
```

## Managing Volumes

Attach volumes to a container:

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "volumes": [
      {
        "container_path": "/data",
        "volume_name": "my-data",
        "read_only": false
      }
    ]
  }'
```

## Checking Running State

Query the current state of all containers:

```bash
curl "https://api.hola.cloud/api/console/run" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Expected response:

```json
[
  {
    "id": "c8a9f3b2-1234-5678-9abc-def012345678",
    "image": "registry.hola.cloud/my-project/my-app:v1",
    "status": "running",
    "env": { "LOG_LEVEL": "info" },
    "volumes": [{ "container_path": "/data", "volume_name": "my-data" }],
    "created_at": "2025-06-21T12:00:00Z"
  }
]
```

## Complete Workflow Example

A full end-to-end workflow:

1. **Login** to the registry:
   ```bash
   docker login https://registry.hola.cloud
   ```

2. **Build** your image:
   ```bash
   docker build -t registry.hola.cloud/my-project/my-app:v1 .
   ```

3. **Push** the image:
   ```bash
   docker push registry.hola.cloud/my-project/my-app:v1
   ```

4. **Start** a container:
   ```bash
   curl -X POST "https://api.hola.cloud/api/console/run" \
     -H "Api-Key: YOUR_API_KEY" \
     -H "Api-Secret: YOUR_API_SECRET" \
     -H "Content-Type: application/json" \
     -d '{"image": "registry.hola.cloud/my-project/my-app:v1"}'
   ```

5. **Verify** it is running:
   ```bash
   curl "https://api.hola.cloud/api/console/run" \
     -H "Api-Key: YOUR_API_KEY" \
     -H "Api-Secret: YOUR_API_SECRET"
   ```

6. **Stop** the container when done:
   ```bash
   curl -X DELETE "https://api.hola.cloud/api/console/run/CONTAINER_ID" \
     -H "Api-Key: YOUR_API_KEY" \
     -H "Api-Secret: YOUR_API_SECRET"
   ```
