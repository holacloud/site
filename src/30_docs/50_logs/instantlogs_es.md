# InstantLogs

Plataforma de logs en tiempo real para ingestión, consulta y alertas basadas en eventos.

## Descripción y casos de uso
- **Observabilidad centralizada** para aplicaciones y funciones Lambda.
- **Auditoría** de cambios y accesos en servicios críticos.
- **Análisis en streaming**: detección de anomalías y respuestas automatizadas.

## Inicio rápido
1. **Prerrequisitos**: `holactl` instalado y token de proyecto.
2. **Enviar los primeros logs**:
   ```bash
   holactl logs stream create app-demo --retention 7d
   holactl logs send app-demo --file logs/sample.log --format json
   ```
3. **Consultar en tiempo real**:
   ```bash
   holactl logs query app-demo --expr 'status >= 500' --tail
   holactl logs export app-demo --expr 'service="checkout"' --out errors.ndjson
   ```
4. **Crear alerta**:
   ```bash
   holactl logs alert create checkout-5xx --stream app-demo --expr 'status >= 500' \
     --threshold 10 --window 5m --action tailon:dlq
   ```

## Conceptos y arquitectura
- **Streams** con retención configurable y almacenamiento indexado para filtros rápidos.
- **Esquema flexible**: acepta JSON o texto con parseo por plantillas y extracción automática de campos comunes.
- **Enrutamiento**: exportación hacia buckets de Files, colas Tailon o webhooks.
- **Límites**: tasa recomendada inicial de 5k eventos/s por stream; aumenta bajo petición.

## Guías paso a paso
- **Filtrado avanzado**: usa expresiones con campos estructurados (`status`, `latency_ms`, `trace_id`).
- **Exportación y archivado**: programa exportaciones diarias a Files en clase `archive`.
- **Alertas basadas en eventos**: combina condiciones y acciones (Tailon DLQ, webhooks, Lambda).
- **Retención y costos**: ajusta retención por stream y habilita compresión en exportaciones.

## Operaciones y observabilidad
- **Métricas clave**: tasa de ingestión, latencia de indexación, errores de parseo, tamaño de retención.
- **Alertas base**: ingestión rechazada, backlog de indexación, fallos en acciones de alerta.
- **Runbooks**: aumentar particiones de ingestión, depurar plantillas de parseo, reenviar lotes fallidos.

## Seguridad y cumplimiento
- Cifrado de datos en tránsito y reposo, con controles de acceso por stream.
- Redacción de campos sensibles con políticas de máscara.
- Auditoría completa de consultas y cambios de configuración.

## Referencia rápida
- **CLI**: `holactl logs stream|query|alert|export <action>`.
- **Errores comunes**: `RATE_LIMITED` (reduce volumen o solicita capacidad), `INVALID_EXPRESSION`, `ACTION_FAILED`.
- **Límites**: 500 streams por proyecto, retención máxima 365 días, tamaño de evento recomendado < 256 KB.

## Resolución de problemas
- **Eventos descartados**: revisa `holactl logs dlq inspect` y corrige el formato.
- **Consultas lentas**: agrega índices en campos frecuentes y reduce ventanas de tiempo.
- **Alertas ruidosas**: aplica umbrales por frecuencia y añade condiciones adicionales.
