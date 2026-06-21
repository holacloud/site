# Run API Reference

Run is HolaCloud's container execution service. It exposes a REST API for container management and a Docker Registry v2 compatible interface for image storage.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/version` | Return the Run service version |
| `GET` | `/api/console` | Console data — repositories, images, state |
| `POST` | `/api/console/start` | Start a container |
| `POST` | `/api/console/stop` | Stop a container |
| `POST` | `/api/console/rollback` | Rollback to a previous image |
| `PUT` | `/api/console/env` | Save environment variables |
| `PUT` | `/api/console/volumes` | Save volume configuration |
| `GET` | `/v2/` | Docker Registry API version check |
| `HEAD` | `/v2/*` | Docker Registry blob/manifest check |
| `POST` | `/v2/*/blobs/uploads/` | Docker Registry blob upload |
| `PATCH` | `/v2/*/blobs/uploads/:uuid` | Docker Registry blob upload chunk |
| `PUT` | `/v2/*/blobs/uploads/:uuid` | Docker Registry blob upload finalize |
| `PUT` | `/v2/*/manifests/:tag` | Docker Registry manifest push |

## Authentication

Console API endpoints (`/api/console/*`) are public. Docker Registry endpoints require Basic Authentication using the DockerAuth mechanism.
