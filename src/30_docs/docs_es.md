# Documentación de HolaCloud

Bienvenido a la nueva experiencia de documentación inspirada en AWS, Google Cloud y Azure. Aquí encontrarás recorridos rápidos, guías operativas y referencia técnica para cada servicio de HolaCloud.

## Navegación rápida
- **Guías generales**: [Identidad y acceso](#identidad-y-acceso), [Red y dominios](#red-y-dominios), [Observabilidad](#observabilidad-y-operaciones), [Seguridad y cumplimiento](#seguridad-y-cumplimiento), [Facturación y límites](#facturacion-y-limites).
- **Servicios**: [InceptionDB](10_inceptiondb/inceptiondb_es.md), [Lambda System](20_lambda/lambda_es.md), [Files](30_files/files_es.md), [Config](40_config/config_es.md), [InstantLogs](50_logs/instantlogs_es.md), [Tailon](60_queues/tailon_es.md).
- **Recursos comunes**: SDK/CLI, plantillas de infraestructura y ejemplos descargables en `statics/`.
- **Novedades**: notas de versión y calendario de deprecaciones (en preparación).

## Tabla de servicios y recursos
| Servicio | Quickstart | Guías paso a paso | Referencia | Operaciones |
| --- | --- | --- | --- | --- |
| InceptionDB | [Creación de clúster](10_inceptiondb/inceptiondb_es.md#inicio-rapido) | Migración, particionado, backup/restore | API de administración y límites | Métricas, alertas y runbooks |
| Lambda System | [Primera función](20_lambda/lambda_es.md#inicio-rapido) | Versiones/alias, orquestación, empaquetado | Límites y códigos de error | Observabilidad y control de versiones |
| Files | [Bucket y objetos](30_files/files_es.md#inicio-rapido) | Versionado, ciclo de vida, replicación | API/CLI y URLs firmadas | Costos, retención y recuperación |
| Config | [Espacio y lectura](40_config/config_es.md#inicio-rapido) | Feature flags, plantillas por entorno | API/CLI y auditoría | Rollbacks y supervisión |
| InstantLogs | [Ingesta y consultas](50_logs/instantlogs_es.md#inicio-rapido) | Filtrado avanzado, exportaciones | API/CLI y formatos | Retención y alertas |
| Tailon | [Cola y consumidor](60_queues/tailon_es.md#inicio-rapido) | DLQ, reintentos, fanout | API/CLI y límites | Idempotencia y monitoreo |

## Guías generales
### Identidad y acceso
- Autenticación mediante tokens de proyecto y roles definidos (Owner, Operator, Reader).
- Usa políticas de servicio para limitar acciones por entorno y aplica etiquetas `team` y `env` para segmentar permisos.
- Integra el SSO corporativo mediante OpenID Connect y rota credenciales cada 90 días.

### Red y dominios
- Expone endpoints mediante dominios gestionados por HolaCloud con TLS automático.
- Usa redes privadas para tráfico entre servicios y habilita listas de control de acceso por IP para consolas administrativas.
- Configura reglas de egress controlado para evitar filtraciones y centraliza DNS interno.

### Observabilidad y operaciones
- Publica métricas en el colector central (`holactl metrics push`) y crea dashboards base por servicio.
- Define alertas mínimas: disponibilidad, latencia p95 y errores 5xx; incluye runbooks vinculados.
- Registra auditoría de cambios y activa exportación de logs hacia InstantLogs para correlación.

### Seguridad y cumplimiento
- Cifra datos en reposo y en tránsito; usa claves administradas y rota certificados automáticamente.
- Habilita registro de auditoría y revisiones mensuales de accesos privilegiados.
- Etiqueta versiones como Preview/GA/Deprecated y documenta compatibilidad de SDK/CLI.

### Facturación y límites
- Consulta límites por servicio en la sección de referencia correspondiente.
- Configura alertas de presupuesto y reportes por etiqueta `team`.
- Usa tablas de costos por región y por tipo de recurso (clústeres, funciones, almacenamiento, mensajes).

## Recursos comunes
- **SDK/CLI**: `holactl` y SDKs por lenguaje con ejemplos en `statics/examples/` (añade tu token con `holactl login`).
- **Plantillas de infraestructura**: archivos base para desplegar entornos de desarrollo y producción.
- **Ejemplos descargables**: paquetes listos para ejecutar en los quickstarts de cada servicio.

## Novedades y cambios
- Publicaremos notas de versión trimestrales con mejoras, correcciones y cambios de límites.
- Incluiremos un calendario de deprecaciones y una hoja de ruta resumida por servicio.
