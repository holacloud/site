# Referencia de la API de la Consola

La Consola expone un conjunto de endpoints de API simulados utilizados por la interfaz web durante el desarrollo y el modo de demostración. Estos endpoints simulan servicios backend y no requieren autenticación.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| `GET` | `/version` | Devuelve la versión actual de la Consola |
| `GET` | `/fakeapi/*` | Endpoints simulados para desarrollo |

### GET /version

Devuelve la cadena de versión del servicio Consola.

### GET /fakeapi/*

Endpoint comodín que proporciona datos simulados para varios servicios backend, incluyendo autenticación, InceptionDB, Lambda, Files y gestión de proyectos.
