# Files

Almacenamiento de objetos seguro con control de acceso y URLs firmadas.

## Descripción y casos de uso
- **Media estática**: imágenes, videos y assets web con CDN opcional.
- **Backups y archivos empresariales**: retención configurable y versiones.
- **Integración por eventos**: dispara Lambdas o colas Tailon en cambios de objetos.

## Inicio rápido
1. **Prerrequisitos**: `holactl` instalado y token de proyecto.
2. **Crear bucket y cargar contenido**:
   ```bash
   holactl files bucket create assets --region us-east
   holactl files object put assets hero.png --file ./img/hero.png --acl private
   holactl files object get assets hero.png --out /tmp/hero.png
   ```
3. **Generar URL firmada**:
   ```bash
   holactl files url sign assets hero.png --expires 15m
   ```
4. **Habilitar versionado y ciclo de vida**:
   ```bash
   holactl files bucket versioning assets --enable
   holactl files bucket lifecycle apply assets --file policies/lifecycle.yml
   ```

## Conceptos y arquitectura
- **Buckets** por región con almacenamiento replicado y clases de costo (estándar, archive).
- **Controles de acceso**: ACL por objeto, políticas por bucket y etiquetas para gobernanza.
- **Eventos**: notificaciones hacia Lambda/Tailon/InstantLogs por operaciones `put`, `delete` o expiración.
- **Límites**: tamaño de objeto recomendado < 5 GB para subidas únicas; usa multipart para mayores.

## Guías paso a paso
- **Versionado y restauración**: activa versionado y recupera con `holactl files object restore --version-id`.
- **Ciclo de vida**: mueve a clase `archive` tras N días y elimina versiones huérfanas.
- **Replicación**: configura reglas cross-region con filtros por prefijo o etiqueta.
- **Costos**: habilita reportes por bucket y clase de almacenamiento; usa `--storage-class archive` para datos fríos.

## Operaciones y observabilidad
- **Métricas clave**: solicitudes por tipo, bytes transferidos, tasa de aciertos de caché, eventos fallidos.
- **Alertas base**: errores 5xx, latencia p95 alta, fallos en replicación, expiración de URLs firmadas.
- **Runbooks**: reintentar replicaciones, rotar claves de firma, limpiar objetos expirados.

## Seguridad y cumplimiento
- Cifrado en reposo y en tránsito habilitado por defecto.
- Políticas de bucket con condiciones por IP, usuario y etiquetas; soporte de bloqueo WORM opcional.
- Auditoría para todas las operaciones de escritura/lectura sensibles.

## Referencia rápida
- **CLI**: `holactl files bucket|object|url <action>`.
- **Errores comunes**: `ACCESS_DENIED` (revisa políticas), `EXPIRED_URL`, `OBJECT_TOO_LARGE` (usa multipart).
- **Límites**: 1.000 buckets por proyecto, URLs firmadas con expiración máx de 7 días, 5.000 reglas de ciclo de vida.

## Resolución de problemas
- **Replicación atrasada**: valida conectividad y revisa `holactl files replication status`.
- **Accesos no autorizados**: audita políticas y rota claves de firma.
- **Subidas lentas**: cambia a multipart, acerca al cliente a la región y habilita aceleración.
