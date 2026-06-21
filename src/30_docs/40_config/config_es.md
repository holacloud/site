# Config

Config es un servicio centralizado de gestión de configuraciones dentro del ecosistema Hola.Cloud. Permite almacenar, recuperar y administrar configuraciones de aplicaciones de forma jerárquica, versionada y segura.

## Beneficios Clave

### Configuración Jerárquica
Config soporta una estructura jerárquica de niveles: proyecto, entorno y servicio. Esto permite definir configuraciones base y sobrescribir valores específicos por entorno (desarrollo, staging, producción) o por servicio, reduciendo la duplicación y simplificando la administración.

### Basado en JSON
Todas las configuraciones se almacenan como documentos JSON, lo que facilita su lectura, escritura e integración con cualquier lenguaje de programación o herramienta. Los esquemas son flexibles, permitiendo almacenar cualquier estructura JSON válida.

### Versionado
Cada actualización de configuración crea una nueva versión. Puedes rastrear cambios a lo largo del tiempo, comparar revisiones y revertir a una versión anterior si es necesario.

### Fácil Integración
Config se integra sin problemas con otros servicios de Hola.Cloud como InceptionDB, Lambda e InstantLogs. Centralizar las configuraciones entre servicios reduce el trabajo repetitivo y mantiene tu infraestructura consistente.

### Acceso Seguro
El servicio expone dos superficies API. La API de administración (`/v0/configs`) es de acceso público para operaciones de gestión. La API de usuario (`/v1/config`) requiere una clave API, asegurando que solo clientes autorizados lean o modifiquen la configuración en tiempo de ejecución.

## Resumen de la API

| Método | Ruta | Descripción | Autenticación |
|--------|------|-------------|---------------|
| GET | `/v0/configs` | Listar todas las configuraciones (admin) | Pública |
| POST | `/v0/configs` | Crear una nueva configuración (admin) | Pública |
| GET | `/v0/configs/{id}` | Obtener una configuración por ID (admin) | Pública |
| DELETE | `/v0/configs/{id}` | Eliminar una configuración (admin) | Pública |
| PATCH | `/v0/configs/{id}` | Actualización parcial de una configuración (admin) | Pública |
| GET | `/v1/config` | Obtener la configuración de usuario actual | API Key |
| PATCH | `/v1/config` | Actualizar la configuración de usuario actual | API Key |

URL base: `https://api.hola.cloud`

## Mejores Casos de Uso

### Configuración de Microservicios
Centraliza y versiona la configuración de todos tus microservicios. Cada servicio obtiene su propia configuración al iniciar y recibe actualizaciones sin necesidad de redespliegue.

### Despliegues Multi-Entorno
Administra configuraciones separadas para desarrollo, staging y producción desde un único panel de control. Sobrescribe solo los valores que difieren entre entornos.

### Feature Flags
Almacena banderas de funcionalidad como entradas de configuración y actualízalas en tiempo de ejecución sin cambios de código. El endpoint PATCH soporta actualizaciones parciales, facilitando la activación o desactivación de banderas individuales.
