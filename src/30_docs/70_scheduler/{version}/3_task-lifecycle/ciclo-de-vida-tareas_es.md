# Ciclo de Vida de Tareas

El Scheduler mantiene una cola retrasada en memoria. Una tarea se encola con `id`, `future` o `delay`, `payload` y `labels` como array de strings. Los trabajadores reservan tareas con un lease y las eliminan para confirmar el procesamiento.

## Encolar

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "tarea-001",
    "payload": { "accion": "generar_informe" },
    "delay": "1h",
    "labels": ["equipo:analitica", "entorno:produccion"]
  }'
```

Respuesta: HTTP `202 Accepted` con cuerpo vacío.

## Listar

```json
{
  "scheduled": [
    {
      "id": "tarea-001",
      "future": "2025-06-22T00:00:00Z",
      "labels": ["equipo:analitica"]
    }
  ],
  "inflight": [],
  "scheduled_meta": { "page": 1, "per_page": 25, "total": 1, "total_pages": 1 },
  "inflight_meta": { "page": 1, "per_page": 25, "total": 0, "total_pages": 0 }
}
```

## Reservar

```json
{ "worktime": "60s" }
```

```json
{
  "id": "tarea-001",
  "payload": { "accion": "generar_informe" },
  "lease_expires_at": "2025-06-21T12:01:00Z",
  "labels": ["equipo:analitica"]
}
```

Si no hay tareas disponibles, la API devuelve `204 No Content`.

## Extender Lease

```json
{ "extension": "30s" }
```

```json
{ "lease_expires_at": "2025-06-21T12:01:30Z" }
```

## Confirmar

El trabajador confirma una tarea eliminándola:

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-001" \
  -H "X-API-Key: TU_API_KEY"
```

Respuesta: HTTP `204 No Content`.

## Reprogramar

```json
{ "delay": "5m" }
```

```json
{
  "id": "tarea-001",
  "future": "2025-06-21T12:06:00Z"
}
```

## SSE

El stream SSE emite solo eventos `snapshot` con `scheduled`, `inflight`, metas de paginacion y `health`.
