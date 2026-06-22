# Produciendo Consumiendo

Este documento cubre el formato de los mensajes, estrategias de escritura, lectura con long-poll, manejo de errores y patrones comunes de consumo.

## Formato del Mensaje

Cada mensaje en Tailon es un objeto JSON arbitrario. Tailon no impone un esquema — tú decides la estructura. Sin embargo, el servidor asigna metadatos a cada mensaje:

```json
{
  "id": "msg-001",
  "type": "pedido.creado",
  "order_id": "ord-1234",
  "amount": 49.99,
  "timestamp": "2025-06-21T12:00:01Z"
}
```

Los campos `id` y `timestamp` se añaden automáticamente cuando el mensaje se encola.

## Escribir Mensajes

### Escribir un Solo Mensaje

Envía un único objeto JSON con `Content-Type: application/json`:

```bash
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "pedido.creado",
    "order_id": "ord-1234",
    "amount": 49.99
  }'
```

```http
POST /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "type": "pedido.creado",
  "order_id": "ord-1234",
  "amount": 49.99
}
```

Respuesta esperada:

```json
{
  "written": 1
}
```

### Escribir Múltiples Mensajes (NDJSON)

Usa `Content-Type: application/x-ndjson` para enviar varios mensajes en una sola solicitud, un objeto JSON por línea:

```bash
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/x-ndjson" \
  -d '{"type": "pageview", "url": "/inicio", "user": "u1"}
{"type": "pageview", "url": "/precios", "user": "u2"}
{"type": "pageview", "url": "/docs", "user": "u1"}'
```

Respuesta esperada:

```json
{
  "written": 3
}
```

## Leer Mensajes

Tailon utiliza **long-poll** — la solicitud se bloquea hasta que al menos un mensaje esté disponible o se alcance un tiempo de espera del servidor. Usa el encabezado `Limit` para controlar el tamaño del lote.

```bash
curl -X GET "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Limit: 10"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Limit: 10
```

Respuesta esperada (array de mensajes):

```json
[
  {
    "id": "msg-001",
    "type": "pedido.creado",
    "order_id": "ord-1234",
    "amount": 49.99,
    "timestamp": "2025-06-21T12:00:01Z"
  },
  {
    "id": "msg-002",
    "type": "pageview",
    "url": "/inicio",
    "user": "u1",
    "timestamp": "2025-06-21T12:00:02Z"
  }
]
```

Los mensajes se devuelven en orden FIFO y se eliminan de la cola al ser entregados.

## Manejo de Errores

| Código | Significado |
|--------|------------|
| 200 | Éxito (mensajes devueltos o escritos) |
| 204 | Sin contenido — la cola existe pero está vacía (long-poll agotado) |
| 400 | Cuerpo o encabezados de solicitud inválidos |
| 404 | Cola no encontrada |
| 409 | Conflicto en la operación de la cola |
| 500 | Error interno del servidor |

Los errores de productor suelen deberse a JSON mal formado. Los errores de consumo suelen indicar una cola inexistente o un valor `Limit` inválido.

## Patrones de Consumo

### Consumidor Único

Un solo trabajador consulta la cola en un bucle, procesando un lote a la vez:

```bash
while true; do
  curl -X GET "https://api.hola.cloud/v1/queues/id-de-mi-cola" -H "Limit: 5"
  # procesar cada mensaje
done
```

### Múltiples Consumidores

Varios trabajadores pueden consultar la misma cola. Cada mensaje se entrega a exactamente un consumidor porque los mensajes se eliminan al leer.

### Fan-Out con Múltiples Colas

Si necesitas que el mismo mensaje llegue a varios consumidores, crea colas separadas y haz que el productor escriba en cada una, o construye un proceso distribuidor que lea de una cola y escriba en otras.

### Prioridad con Colas Dedicadas

Usa colas separadas para diferentes prioridades de mensajes. Los trabajadores de alta prioridad consultan la cola de alta prioridad con más frecuencia.
