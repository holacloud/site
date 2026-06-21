# Replicación y Alta Disponibilidad

KVNode logra alta disponibilidad mediante un modelo de **replicación líder-seguidor**. Los cambios escritos en un nodo principal (líder) se transmiten en tiempo real a los nodos réplica (seguidores) a través de un protocolo de replicación basado en NDJSON.

## Cómo Funciona la Replicación

1. Un cliente escribe un par clave-valor en cualquier nodo del clúster.
2. El nodo receptor persiste la escritura en su WAL y la aplica al almacenamiento en memoria.
3. Si el nodo tiene un **padre** configurado, reenvía la escritura upstream mediante el endpoint `/v1/replicate`.
4. Si el nodo tiene **hijos**, transmite la escritura a todas las réplicas conectadas a través del flujo de replicación.
5. Cada réplica aplica el cambio en su propio almacenamiento, manteniendo una copia eventualmente consistente.

## Protocolo de Replicación

El endpoint de replicación (`POST /v1/replicate`) usa **NDJSON** (JSON Delimitado por Nueva Línea) como formato de transmisión. El protocolo soporta estos tipos de comando:

| Comando | Descripción |
|---------|-------------|
| `set` | Una clave fue creada o actualizada |
| `delete` | Una clave fue eliminada |
| `baseline_begin` | Inicio de una instantánea completa de colección |
| `baseline_end` | Fin de una instantánea completa de colección |
| `ping` | Heartbeat de mantenimiento |

### Ejemplo: Flujo de Replicación

```
{"type":"baseline_begin","collection":"my-collection","seq":0}
{"type":"set","collection":"my-collection","key":"usuario:alice","value":{"rol":"admin"},"version":1,"timestamp":1700000000}
{"type":"baseline_end","collection":"my-collection","seq":1}
{"type":"set","collection":"my-collection","key":"usuario:bob","value":{"rol":"usuario"},"version":1,"timestamp":1700000001}
{"type":"ping","timestamp":1700000010}
```

## Estado del Nodo

El endpoint `/v1/status` expone la topología de replicación:

```bash
curl "https://api.hola.cloud/v1/status" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Respuesta:

```json
{
  "node": "node-1",
  "role": "primary",
  "parent": "",
  "collections": {
    "mi-coleccion": {
      "keys": 42,
      "lastSeq": 128,
      "walSizeBytes": 65536
    }
  },
  "children": 2,
  "uptimeSeconds": 86400,
  "replication": {
    "enabled": true,
    "parent_connected": false
  }
}
```

## Métricas del Nodo

Monitorea el rendimiento con el endpoint `/v1/metrics`:

```bash
curl "https://api.hola.cloud/v1/metrics" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Respuesta:

```json
{
  "writes_total": 15000,
  "reads_total": 82000,
  "replication_commands_sent_total": 15000,
  "replication_commands_received_total": 0,
  "children_connected": 2,
  "parent_connected": false
}
```

## Verificaciones de Disponibilidad

El endpoint `/readyz` verifica que un nodo esté listo para recibir tráfico. Para las réplicas, esto significa que la conexión con el padre debe estar establecida y el WAL completamente reproducido. Para los primarios, la disponibilidad es inmediata después del inicio.

```bash
curl "https://api.hola.cloud/readyz"
```

Respuesta cuando no está listo:

```json
{"ok":false,"node":"node-2","role":"replica","ready":false,"checks":{"wal_replayed":true,"parent_connected":false}}
```

## Backends de WAL Intercambiables

KVNode soporta múltiples backends de WAL, seleccionables al iniciar:

| Backend | Descripción | Caso de Uso |
|---------|-------------|-------------|
| **Memoria** | Búfer en proceso (no durable) | Desarrollo, pruebas, cachés efímeras |
| **Kafka** | Log durable via Apache Kafka | Despliegues de producción de alto rendimiento |
| **PostgreSQL** | WAL almacenado en tablas de PostgreSQL | Entornos que ya usan Postgres |
| **Redis** | WAL almacenado en streams de Redis | Arquitecturas de baja latencia centradas en Redis |
| **MongoDB** | WAL almacenado en colecciones de MongoDB | Despliegues centrados en MongoDB |

Cada backend implementa la misma interfaz de WAL, por lo que cambiar entre ellos no requiere modificaciones a nivel de aplicación.
