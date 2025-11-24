# InceptionDB

Base de datos NoSQL distribuida para cargas con alta concurrencia y baja latencia.

## Descripción y casos de uso
- **Aplicaciones en tiempo real**: monitoreo financiero, IoT y analítica streaming.
- **Catálogos de contenidos**: esquemas flexibles para e-commerce y media.
- **Grafo y relaciones**: recomendaciones, redes sociales y permisos jerárquicos.

## Inicio rápido
1. **Prerrequisitos**: `holactl` instalado, token de proyecto y red privada habilitada.
2. **Crear clúster**:
   ```bash
   holactl inceptiondb cluster create --name demo --nodes 3 --region eu-west
   holactl inceptiondb cluster status demo
   ```
3. **Conectar y probar** (SDK o CLI):
   ```bash
   holactl inceptiondb collection create users --cluster demo --pk userId
   holactl inceptiondb doc put users --cluster demo --file examples/users.json
   holactl inceptiondb query run --cluster demo --query 'SELECT * FROM users LIMIT 5'
   ```
4. **Observa las métricas iniciales**:
   ```bash
   holactl metrics watch --service inceptiondb --cluster demo
   ```

## Conceptos y arquitectura
- **Colecciones** con particiones automáticas por clave primaria y réplicas en zonas distintas.
- **Índices secundarios** y **materializaciones** para consultas de baja latencia.
- **Acceso** mediante roles por colección y políticas de atributos (`team`, `env`).
- **Límites predeterminados**: 10 TB por clúster, p99 < 20 ms para lecturas, 10k ops/s por partición.

## Guías paso a paso
- **Migración de datos**: exporta desde tu origen a `CSV/JSON`, carga con `holactl inceptiondb bulk load`, valida con `holactl inceptiondb validate`.
- **Particionado y keys**: usa claves compuestas (`tenant#userId`) para distribuir tráfico y evita hot partitions revisando métricas de `skew`.
- **Backup y restore**: programa snapshots horarios (`holactl inceptiondb backup schedule`) y restaura a un clúster nuevo para pruebas.
- **Performance tuning**: activa caché de lectura (`--read-cache on`), ajusta niveles de consistencia y revisa índices no usados.

## Operaciones y observabilidad
- **Métricas clave**: latencia p95/p99, tasa de throttling, tamaño por partición, replicación atrasada.
- **Alertas base**: indisponibilidad del clúster, `% throttled > 2%`, replicación atrasada > 30s, uso de disco > 80%.
- **Runbooks**: escalar nodos (`holactl inceptiondb cluster scale --nodes N`), regenerar índices, rotar claves de cifrado.

## Seguridad y cumplimiento
- Cifrado en tránsito (TLS) y en reposo con claves administradas por HolaCloud.
- Políticas de acceso por colección y auditoría de todas las operaciones administrativas.
- Registros inmutables para accesos de lectura/escritura y soporte de retención configurable.

## Referencia rápida
- **CLI**: `holactl inceptiondb <resource> <action>` (cluster, collection, doc, query, backup).
- **Errores comunes**: `IDEMPOTENT_REQUIRED` (reintentos sin idempotency key), `THROTTLED` (aumenta throughput o reintenta exponencial).
- **Límites**: 1.000 colecciones por clúster, 500 índices secundarios, tamaño de documento recomendado < 512 KB.

## Resolución de problemas
- **Consultas lentas**: revisa plan de ejecución con `holactl inceptiondb query explain` y agrega índices.
- **Hot partitions**: redistribuye claves o habilita particiones manuales.
- **Conflictos de escritura**: usa operaciones idempotentes y consistencia transaccional donde aplique.
