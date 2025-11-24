# Lambda System

Ejecución de funciones sin servidores con escalado automático y facturación por consumo.

## Descripción y casos de uso
- **Eventos y colas**: procesar mensajes de Tailon o webhooks externos.
- **Cron y tareas programadas**: automatizar mantenimientos y pipelines ligeros.
- **APIs ligeras**: exponer endpoints rápidos con autenticación integrada.

## Inicio rápido
1. **Prerrequisitos**: `holactl` instalado, token y runtime soportado (Node.js, Python, Go).
2. **Crear la primera función**:
   ```bash
   holactl lambda init hello --runtime node18 --template http
   holactl lambda deploy hello --source ./hello
   holactl lambda invoke hello --payload '{"name":"Mundo"}'
   ```
3. **Versiones y alias**:
   ```bash
   holactl lambda publish-version hello --description "Primera release"
   holactl lambda alias set hello --alias prod --version 1
   ```
4. **Programar con cron**:
   ```bash
   holactl lambda schedule create nightly-backup --function hello --cron "0 2 * * *"
   ```

## Conceptos y arquitectura
- **Runtimes administrados** con aislamiento por contenedor y warm pools para arranques rápidos.
- **Desencadenadores**: colas Tailon, HTTP, cron y eventos de Files/Config.
- **Políticas de ejecución**: permisos mínimos para acceder a otros servicios; cada alias hereda políticas dedicadas.
- **Límites base**: 512 MB–4 GB de memoria, tiempo máx 15 min, payload síncrono 6 MB.

## Guías paso a paso
- **Variables seguras**: almacena secretos en Config y márcalos como `--secret` al vincular a la función.
- **Orquestación**: usa colas y estados en InceptionDB para flujos de larga duración.
- **Testing y empaquetado**: incluye dependencias locales, ejecuta `holactl lambda test --local` y valida tiempos de arranque.
- **Despliegues azules/verdes**: crea alias `blue` y `green` y mueve el peso de tráfico con `holactl lambda alias shift`.

## Operaciones y observabilidad
- **Métricas clave**: invocaciones, duración p95, cold starts, errores, reintentos y DLQ.
- **Alertas base**: tasa de errores > 1%, cold starts sostenidos, saturación de concurrencia.
- **Runbooks**: incrementar concurrencia reservada, limpiar versiones antiguas, reprocesar mensajes desde DLQ.

## Seguridad y cumplimiento
- Ejecución en entornos aislados con políticas mínimas necesarias.
- Soporte de firmas para webhooks y validez de tokens de identidad.
- Auditoría de despliegues y cambios de alias/versiones.

## Referencia rápida
- **CLI**: `holactl lambda <action>` (deploy, invoke, publish-version, alias, schedule, logs).
- **Errores comunes**: `TIMEOUT` (aumenta tiempo o optimiza), `UNAUTHORIZED` (revisa políticas de ejecución), `PAYLOAD_TOO_LARGE`.
- **Límites**: 1.000 funciones por proyecto, 100 versiones activas por función, DLQ retenida 14 días por defecto.

## Resolución de problemas
- **Tiempos altos de arranque**: usa runtimes con warm pool y revisa tamaño del paquete.
- **Mensajes en DLQ**: inspecciona evento, corrige código y reprocesa con `holactl lambda dlq replay`.
- **Erros intermitentes 5xx**: activa reintentos exponenciales y revisa dependencias externas.
