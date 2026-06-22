# Referencia de API de Run

Run expone una API pequeña de Console y un subconjunto Docker Registry v2 orientado a push bajo `/v2`.

## Endpoints de Console

| Método | Ruta | Descripción |
|--------|------|-------------|
| `GET` | `/version` | Devuelve la versión como texto plano |
| `GET` | `/api/console?repository=` | Devuelve datos de Console para un repositorio |
| `POST` | `/api/console/start` | Inicia un repositorio con una referencia o digest |
| `POST` | `/api/console/stop` | Detiene un repositorio |
| `POST` | `/api/console/rollback` | Revierte un repositorio a una referencia o digest |
| `PUT` | `/api/console/env` | Guarda variables de entorno del repositorio |
| `PUT` | `/api/console/volumes` | Guarda configuración de volúmenes del repositorio |

La API no usa `container_id` y no expone `/v1/run/deploy`, `/v1/run/push`, `/v1/run/exec`, `/api/console/run` ni endpoints similares de run/exec.

## Endpoints de Registry

| Método | Ruta | Descripción |
|--------|------|-------------|
| `GET` | `/v2/` | Verificación de API Registry |
| `HEAD` | `/v2/*` | Verificación de blob o manifiesto durante push |
| `POST` | `/v2/*/blobs/uploads/` | Inicia subida de blob |
| `PATCH` | `/v2/*/blobs/uploads/:uuid` | Sube datos del blob |
| `PUT` | `/v2/*/blobs/uploads/:uuid` | Finaliza subida con digest |
| `PUT` | `/v2/*/manifests/:reference` | Sube manifiesto |

Es un subconjunto orientado a push, no una API registry completa para pull.
