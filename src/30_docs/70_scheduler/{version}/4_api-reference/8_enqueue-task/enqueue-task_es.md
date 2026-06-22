# Encolar Tarea

Encola una nueva tarea para ejecutarse en un horario programado.

## Autenticación

Requiere autenticación. Pasa tu clave API mediante el encabezado `X-API-Key` o `Authorization: Bearer`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | string | ID de tarea requerido |
| future | string | Marca de tiempo ISO 8601 para la disponibilidad |
| delay | string | Duración Go desde ahora, como `60s` o `5m` (alternativa a future) |
| payload | object | Carga útil JSON arbitraria para el trabajador |
| labels | array de strings | Etiquetas opcionales para filtrado |

```json
{
  "payload": {
    "type": "enviar_correo",
    "to": "usuario@example.com",
    "template": "bienvenida"
  },
  "delay": "60s",
  "labels": ["proyecto:incorporacion", "prioridad:alta"]
}
```

## Ejemplo de Solicitud

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "payload": {
      "type": "enviar_correo",
      "to": "usuario@example.com",
      "template": "bienvenida"
    },
    "delay": "60s",
    "labels": ["proyecto:incorporacion", "prioridad:alta"]
  }'
```

## Ejemplo de Respuesta

```http
HTTP/1.1 202 Accepted
```

El cuerpo de la respuesta está vacío.

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | invalid_json | JSON inválido |
| 400 | validation_error | Falta id, future/delay inválido o labels inválidas |
| 401 | unauthorized | Clave API faltante o inválida |
| 409 | task_already_exists | La tarea ya existe |
| 500 | internal_error | Error interno del servidor |
