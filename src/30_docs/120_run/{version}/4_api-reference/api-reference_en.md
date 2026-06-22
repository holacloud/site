# Run API Reference

Run exposes a small Console API and a push-oriented subset of Docker Registry v2 under `/v2`.

## Console API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/version` | Return the service version as plain text |
| `GET` | `/api/console?repository=` | Return Console data for a repository |
| `POST` | `/api/console/start` | Start a repository at a reference or digest |
| `POST` | `/api/console/stop` | Stop a repository |
| `POST` | `/api/console/rollback` | Roll back a repository to a reference or digest |
| `PUT` | `/api/console/env` | Save repository environment variables |
| `PUT` | `/api/console/volumes` | Save repository volume configuration |

The Console API does not use `container_id` and does not expose `/v1/run/deploy`, `/v1/run/push`, `/v1/run/exec`, `/api/console/run`, or similar run/exec endpoints.

## Registry Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/v2/` | Registry API version check |
| `HEAD` | `/v2/*` | Blob or manifest existence check used during push |
| `POST` | `/v2/*/blobs/uploads/` | Start blob upload |
| `PATCH` | `/v2/*/blobs/uploads/:uuid` | Upload blob data |
| `PUT` | `/v2/*/blobs/uploads/:uuid` | Finalize blob upload with digest |
| `PUT` | `/v2/*/manifests/:reference` | Push manifest |

This is a push-oriented subset, not a full pull-capable registry API.
