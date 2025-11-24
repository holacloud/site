# Tailon

Sistema de colas para trabajo asíncrono con entrega garantizada y visibilidad completa.

## Descripción y casos de uso
- **Procesamiento de tareas**: imágenes, pagos, workflows de datos.
- **Integración de servicios**: desacopla productores y consumidores entre microservicios.
- **Fanout y buffering**: reparte eventos a múltiples consumidores y absorbe picos.

## Inicio rápido
1. **Prerrequisitos**: `holactl` instalado y permisos sobre el proyecto.
2. **Crear cola y productor**:
   ```bash
   holactl tailon queue create jobs --visibility-timeout 30s --dlq jobs-dlq
   holactl tailon message send jobs --body '{"task":"resize","id":1}'
   ```
3. **Consumir mensajes**:
   ```bash
   holactl tailon message receive jobs --max-messages 5 --wait 10s
   holactl tailon message delete jobs --receipt-handle <handle>
   ```
4. **Configurar reintentos y DLQ**:
   ```bash
   holactl tailon queue policy jobs --max-retries 5 --backoff exponential
   holactl tailon queue inspect jobs-dlq
   ```

## Conceptos y arquitectura
- **Colas** con réplicas y particiones automáticas para throughput horizontal.
- **Mensajes** con atributos y trazas; expiración configurable.
- **DLQ** para mensajes fallidos y **visibilidad** para bloquear mientras se procesa.
- **Límites**: tamaño recomendado de mensaje < 256 KB; usa punteros a Files para cargas grandes.

## Guías paso a paso
- **Idempotencia**: incluye `message-id` único y registra progreso en InceptionDB para reintentos seguros.
- **Patrones de fanout**: usa colas secundarias con `holactl tailon queue subscribe` o publica múltiples mensajes derivados.
- **Programación**: combina con Lambda `schedule` o usa atributos de retraso (`--delay`).
- **Monitoreo de DLQ**: inspecciona y reprocesa con `holactl tailon dlq replay` tras corregir el error.

## Operaciones y observabilidad
- **Métricas clave**: mensajes en vuelo, tasa de reintentos, edad de mensaje más antiguo, tasa de aciertos DLQ.
- **Alertas base**: crecimiento acelerado de DLQ, tiempo de visibilidad excedido, latencia de entrega alta.
- **Runbooks**: aumentar particiones, ajustar visibilidad, habilitar consumidores adicionales y limpiar mensajes expirados.

## Seguridad y cumplimiento
- Control de acceso por cola y cifrado en tránsito y reposo.
- Auditoría de operaciones administrativas y de lectura/escritura sensibles.
- Integración con políticas de red privada para consumidores internos.

## Referencia rápida
- **CLI**: `holactl tailon queue|message|dlq <action>`.
- **Errores comunes**: `VISIBILITY_TIMEOUT` (aumenta el tiempo o confirma antes), `UNAUTHORIZED`, `MESSAGE_TOO_LARGE`.
- **Límites**: 2.000 colas por proyecto, retención máx 14 días, 20 suscripciones fanout por cola.

## Resolución de problemas
- **DLQ creciendo**: inspecciona causa raíz, corrige consumidor y reprocesa con backoff.
- **Duplicados**: activa entrega exactamente una vez con claves idempotentes y confirma rápido.
- **Picos de latencia**: escala consumidores, ajusta particiones y revisa políticas de red.
