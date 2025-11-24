# Config

Servicio centralizado de configuración y feature flags con versionado y auditoría.

## Descripción y casos de uso
- **Parámetros por entorno**: valores diferenciados para dev/stage/prod.
- **Feature flags**: habilita o deshabilita funciones por usuario, región o porcentaje.
- **Secretos de aplicación**: almacena secretos cifrados y consúmelos desde Lambda u otros servicios.

## Inicio rápido
1. **Prerrequisitos**: `holactl` instalado y roles adecuados.
2. **Crear espacio y primer key**:
   ```bash
   holactl config space create app-core --description "Configuración central"
   holactl config set app-core DB_HOST=db.prod.local --env prod
   holactl config set app-core DB_HOST=db.dev.local --env dev
   ```
3. **Leer desde aplicación**:
   ```bash
   holactl config get app-core DB_HOST --env prod --format env
   ```
4. **Crear feature flag**:
   ```bash
   holactl config flag create checkout-rollout --space app-core --percentage 10 --env prod
   ```

## Conceptos y arquitectura
- **Espacios** con versionado automático y aislamiento por entorno.
- **Claves y flags** con etiquetas (`team`, `service`) y metadata para auditoría.
- **Controles de acceso** por entorno y tipo de dato (secreto o texto plano).
- **Límites**: 10.000 claves por espacio, 200 versiones retenidas por defecto.

## Guías paso a paso
- **Plantillas por entorno**: exporta valores a YAML y aplica en masa con `holactl config apply`.
- **Rotación segura**: marca claves como `--secret`, rota con `holactl config rotate` y notifica a los consumidores.
- **Integración CI/CD**: usa `holactl config export --format env` en pipelines y evita imprimir secretos.
- **Rollback**: vuelve a una versión previa con `holactl config rollback --version N` y valida con auditoría.

## Operaciones y observabilidad
- **Métricas clave**: tasa de lecturas/escrituras, latencia de entrega, fallos de autorización.
- **Alertas base**: cambios masivos en ventanas cortas, fallos de replicación, flags sin entorno.
- **Runbooks**: bloquear escrituras en incidentes, validar firmas de contenido, restaurar desde respaldo.

## Seguridad y cumplimiento
- Cifrado de secretos en HSM administrado y control de acceso granular.
- Auditoría completa de lectura/escritura y exportación hacia InstantLogs.
- Políticas obligatorias para marcar datos sensibles y exigir rotación periódica.

## Referencia rápida
- **CLI**: `holactl config <action>` (space, set, get, flag, apply, rollback, export).
- **Errores comunes**: `VERSION_CONFLICT` (reintenta con la última versión), `UNAUTHORIZED`, `SECRET_ONLY_OVER_TLS`.
- **Límites**: 50 espacios por proyecto, 5 regiones replicadas máximo, tamaño de valor recomendado < 16 KB.

## Resolución de problemas
- **Valores desactualizados en clientes**: verifica caching TTL y ejecuta `holactl config invalidate-cache`.
- **Flags aplicados de forma desigual**: revisa segmentación y exporta auditoría para confirmar delivery.
- **Rotación fallida**: restaura la versión previa y vuelve a ejecutar con sesiones activas notificadas.
