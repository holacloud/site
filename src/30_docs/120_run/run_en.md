# Run

Run is HolaCloud's container execution service. It manages Docker containers through a Docker Registry v2 compatible interface, allowing you to deploy and run applications in isolated environments.

## Features

### Docker Registry v2 Compatible
Run is a fully compliant Docker Registry v2. Push and pull images using standard Docker commands — no custom tooling required.

```bash
docker login run.hola.cloud
docker build -t run.hola.cloud/my-project/my-app:latest .
docker push run.hola.cloud/my-project/my-app:latest
```

### Console Management
Manage containers, images, and deployments through the HolaCloud Console. Start, stop, and restart containers, view logs, and monitor resource usage — all from a web interface.

### Environment Variables
Configure your containers with environment variables. Set them at deployment time through the Console or API for flexible, configuration-driven deployments.

```bash
curl -X POST "https://api.hola.cloud/v1/run/deploy" \
  -H "Authorization: Bearer your-token" \
  -d '{
    "image": "run.hola.cloud/my-project/my-app:latest",
    "env": {
      "DATABASE_URL": "postgres://...",
      "LOG_LEVEL": "debug"
    }
  }'
```

### Volumes
Attach persistent volumes to your containers for stateful workloads. Volumes are backed by HolaCloud's distributed storage and survive container restarts.

## Use Cases

### Full Application Deployment
Deploy web applications, APIs, and microservices as Docker containers. Run handles networking, health checks, and automatic restarts.

### Development Environments
Spin up isolated development environments for each branch or developer. Use Run to create ephemeral containers that match your production setup.

### CI/CD Pipelines
Integrate Run into your CI/CD workflow. Push images to the registry as part of your build process, then deploy them automatically to staging or production environments.

## Getting Started

1. **Authenticate** with the registry:
   ```bash
   docker login run.hola.cloud
   ```

2. **Build and push** your image:
   ```bash
   docker build -t run.hola.cloud/my-project/my-app:latest .
   docker push run.hola.cloud/my-project/my-app:latest
   ```

3. **Deploy** through the Console or API:
   ```bash
   curl -X POST "https://api.hola.cloud/v1/run/deploy" \
     -H "Authorization: Bearer your-token" \
     -d '{"image": "run.hola.cloud/my-project/my-app:latest"}'
   ```

Your container will start and be available at the assigned endpoint.
