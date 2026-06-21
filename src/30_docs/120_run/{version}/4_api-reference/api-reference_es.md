# Referencia de la API de Run

Run es el servicio de ejecución de contenedores de HolaCloud. Expone una API REST para la gestión de contenedores y una interfaz compatible con Docker Registry v2 para el almacenamiento de imágenes.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| `GET` | `/version` | Devuelve la versión del servicio Run |
| `GET` | `/api/console` | Datos de la consola — repositorios, imágenes, estado |
| `POST` | `/api/console/start` | Iniciar un contenedor |
| `POST` | `/api/console/stop` | Detener un contenedor |
| `POST` | `/api/console/rollback` | Revertir a una imagen anterior |
| `PUT` | `/api/console/env` | Guardar variables de entorno |
| `PUT` | `/api/console/volumes` | Guardar configuración de volúmenes |
| `GET` | `/v2/` | Verificación de versión del Docker Registry |
| `HEAD` | `/v2/*` | Verificación de blob/manifiesto del Registry |
| `POST` | `/v2/*/blobs/uploads/` | Subida de blob al Registry |
| `PATCH` | `/v2/*/blobs/uploads/:uuid` | Fragmento de subida de blob |
| `PUT` | `/v2/*/blobs/uploads/:uuid` | Finalizar subida de blob |
| `PUT` | `/v2/*/manifests/:tag` | Subida de manifiesto al Registry |

## Autenticación

Los endpoints de la API de Consola (`/api/console/*`) son públicos. Los endpoints del Docker Registry requieren Autenticación Básica mediante el mecanismo DockerAuth.
