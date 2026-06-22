# Listar Schedulers

Devuelve una lista de todos los schedulers.

## Autenticación

Este endpoint es público. No requiere autenticación.

## Parámetros de Consulta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| search | string | Filtrar schedulers por nombre (coincidencia parcial) |
| q | string | Consulta de búsqueda general |

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/schedulers?search=mi"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "schedulers": [
    {
      "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "display_name": "mi-scheduler",
      "ready": true,
      "scheduled": 5,
      "inflight": 0,
      "created_at": "2025-06-20T10:00:00Z",
      "updated_at": "2025-06-21T08:30:00Z"
    }
  ],
  "total": 1
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 500 | internal_error | Error interno del servidor |
