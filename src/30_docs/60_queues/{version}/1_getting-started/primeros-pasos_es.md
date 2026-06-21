# Primeros pasos con Tailon

Esta guía te explica cómo crear una cola, escribir mensajes, leerlos mediante long-poll, consultar las estadísticas de la cola y limpiar los recursos.

## Prerrequisitos

- Una cuenta de HolaCloud.
- [curl](https://curl.se/) instalado en tu máquina.

## Paso 1: Crear una Cola

Crea una nueva cola enviando una solicitud POST con un nombre:

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-cola"
  }'
```

Respuesta esperada:

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "mi-cola",
  "length": 0,
  "reads": 0,
  "writes": 0
}
```

Guarda el `id` devuelto — lo usarás en las siguientes solicitudes.

## Paso 2: Escribir Mensajes

Tailon acepta un flujo de objetos JSON, uno por línea (NDJSON). Escribe tres mensajes:

```bash
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/x-ndjson" \
  -d '{"type": "saludo", "text": "Hola"}
{"type": "saludo", "text": "Hello"}
{"type": "comando", "accion": "ping"}'
```

Respuesta esperada:

```json
{
  "written": 3
}
```

## Paso 3: Leer Mensajes

Usa una solicitud GET con long-poll para recuperar mensajes. El encabezado `Limit` controla cuántos mensajes devolver:

```bash
curl -X GET "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Limit: 2"
```

Respuesta esperada:

```json
[
  {
    "id": "m1",
    "type": "saludo",
    "text": "Hola",
    "timestamp": "2025-06-21T12:00:01Z"
  },
  {
    "id": "m2",
    "type": "saludo",
    "text": "Hello",
    "timestamp": "2025-06-21T12:00:01Z"
  }
]
```

Los mensajes se devuelven en orden FIFO. Repite la solicitud para recuperar los mensajes restantes.

## Paso 4: Consultar Estadísticas de la Cola

Inspecciona los detalles de la cola, incluyendo su longitud actual y el total de lecturas/escrituras:

```bash
curl "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

Respuesta esperada:

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "mi-cola",
  "length": 1,
  "reads": 2,
  "writes": 3
}
```

## Paso 5: Eliminar la Cola

Cuando hayas terminado, elimina la cola para liberar recursos:

```bash
curl -X DELETE "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

Respuesta esperada: HTTP `204 No Content`.

## Siguientes Pasos

- Aprende a gestionar múltiples colas y monitorear clientes activos en [Administrando Colas](../2_managing-queues/administrando-colas_es.md).
- Explora patrones avanzados de producción y consumo en [Produciendo y Consumiendo Mensajes](../3_producing-consuming/produciendo-consumiendo_es.md).
