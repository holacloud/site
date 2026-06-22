# Logs

InstantLogs es un servicio de agregación y transmisión de logs en tiempo real dentro del ecosistema Hola.Cloud. Proporciona ingesta de logs multi-tenant, filtrado basado en expresiones regulares, transmisión en vivo, exportación en NDJSON y almacenamiento persistente, todo con baja latencia.

## Beneficios Clave

### Agregación de Logs en Tiempo Real
Ingiere logs desde múltiples fuentes en tiempo real. InstantLogs recopila, procesa y almacena las entradas de log a medida que llegan, brindándote visibilidad inmediata de tus sistemas.

### Arquitectura Multi-Tenant
Organiza logs en loggers aislados. Cada logger tiene sus propias credenciales, política de retención y controles de acceso, lo que hace que InstantLogs sea adecuado para equipos, clientes o microservicios que operan en la misma cuenta.

### Filtrado por Expresiones Regulares
Filtra flujos de logs usando expresiones regulares. Reduce a las entradas que coinciden con un patrón, excluye ruido y concéntrate en lo que importa sin salir del endpoint de streaming.

### Transmisión en Vivo
Transmite logs a medida que se ingieren a través de conexiones HTTP de larga duración. Elige entre formatos NDJSON y texto plano. Los eventos aparecen en tiempo casi real con una sobrecarga mínima.

### Almacenamiento Persistente
Todos los logs ingeridos se almacenan de forma duradera y están disponibles para reproducción. Puedes re-transmitir datos históricos o exportarlos para análisis sin conexión.

### Modelo de Autenticación Dual
InstantLogs utiliza dos modos de autenticación. Las operaciones de gestión (crear, listar, eliminar loggers, administrar claves API) requieren una clave API (encabezados `Api-Key` + `Api-Secret`). Las operaciones de datos (ingesta, transmisión, filtrado, estadísticas) utilizan un Logger Secret (encabezado `X-Instantlogs-Event-Secret` o `Authorization: Bearer`).

## Resumen de la API

### Endpoints de Gestión (Autenticación con API Key)

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/v1/loggers` | Listar todos los loggers |
| POST | `/v1/loggers` | Crear un nuevo logger |
| GET | `/v1/loggers/{id}` | Obtener detalles de un logger |
| DELETE | `/v1/loggers/{id}` | Eliminar un logger |
| POST | `/v1/loggers/{id}/apiKeys` | Crear una nueva clave API para un logger |
| DELETE | `/v1/loggers/{id}/apiKeys/{key}` | Eliminar una clave API de un logger |

### Endpoints de Datos (Autenticación con Logger Secret)

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | `/v1/loggers/{id}/ingest` | Ingresar entradas de log |
| GET | `/v1/loggers/{id}/filter` | Transmitir y filtrar logs |
| GET | `/v1/loggers/{id}/events` | Transmitir eventos (NDJSON o texto) |
| GET | `/v1/loggers/{id}/stats` | Obtener estadísticas del logger |

### Eventos Enmarcados (Autenticación con Event Secret)

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | `/v1/ingest/events` | Ingresar eventos enmarcados |

URL base: `https://api.hola.cloud`

## Mejores Casos de Uso

### Registro Centralizado
Agrega logs de todos tus servicios, contenedores y servidores en un único flujo. Depura problemas de producción más rápido con búsqueda y filtrado unificados.

### Monitoreo en Tiempo Real
Transmite logs de aplicación e infraestructura en vivo. Detecta anomalías a medida que ocurren y activa alertas basadas en patrones de log.

### Pistas de Auditoría
Almacena registros de log inmutables para cumplimiento normativo y auditoría. Reproduce flujos históricos bajo demanda.

### SaaS Multi-Tenant
Aísla logs por cliente o equipo usando loggers separados. Cada inquilino obtiene sus propias credenciales y políticas de retención.
