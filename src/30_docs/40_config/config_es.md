# Config

Config expone una API JSON pequeña para configuración dentro del ecosistema Hola.Cloud. Almacena entradas de configuración y permite que usuarios autenticados lean o actualicen su configuración de ejecución.

## Beneficios Clave

### Configuración por Entradas
La API v0 almacena una configuración como un Thing con `id` y `entries`. La API de usuario v1 devuelve la configuración del usuario actual como un mapa `entries`.

### Basado en JSON
Todas las configuraciones se almacenan como documentos JSON, lo que facilita su lectura, escritura e integración con cualquier lenguaje de programación o herramienta. Los esquemas son flexibles, permitiendo almacenar cualquier estructura JSON válida.

### Actualizaciones Simples
Las actualizaciones reemplazan o parchean los datos de entrada aceptados por la API.

### Fácil Integración
Config se integra sin problemas con otros servicios de Hola.Cloud como InceptionDB, Lambda e InstantLogs. Centralizar las configuraciones entre servicios reduce el trabajo repetitivo y mantiene tu infraestructura consistente.

### Acceso Seguro
El servicio expone dos superficies API. La API v0 (`/v0/configs`) gestiona Things de configuración. La API de usuario v1 (`/v1/config`) requiere `X-Glue-Authentication` y lee o actualiza la configuración de ejecución del usuario autenticado.

## Resumen de la API

| Método | Ruta | Descripción | Autenticación |
|--------|------|-------------|---------------|
| GET | `/v0/configs` | Listar todas las configuraciones (admin) | Pública |
| POST | `/v0/configs` | Crear una nueva configuración (admin) | Pública |
| GET | `/v0/configs/{id}` | Obtener una configuración por ID (admin) | Pública |
| DELETE | `/v0/configs/{id}` | Eliminar una configuración (admin) | Pública |
| PATCH | `/v0/configs/{id}` | Actualización parcial de una configuración (admin) | Pública |
| GET | `/v1/config` | Obtener el mapa `entries` de configuración del usuario actual | Auth Glue |
| PATCH | `/v1/config` | Actualizar el mapa `entries` de configuración del usuario actual | Auth Glue |

URL base: `https://api.hola.cloud`

## Mejores Casos de Uso

### Configuración de Ejecución
Obtén la configuración del usuario autenticado desde `/v1/config` al iniciar o durante la ejecución.

### Feature Flags
Almacena banderas de funcionalidad como entradas de configuración y actualízalas en tiempo de ejecución sin cambios de código.
