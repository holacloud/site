# POST /api/console/stop

Detiene un contenedor en ejecución.

## Descripción

Detiene gracefulmente un contenedor por su ID. Se concede un período de gracia antes de la terminación forzada.

## Autenticación

Ninguna. Este endpoint es público.

## Cuerpo de la Solicitud

```json
{
  "container_id": "abc123def456"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/api/console/stop" \
  -H "Content-Type: application/json" \
  -d '{"container_id": "abc123def456"}'
```

## Respuesta

```json
{
  "container_id": "abc123def456",
  "status": "stopped",
  "stopped_at": "2026-06-21T11:00:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Contenedor detenido correctamente |
| 400 | Cuerpo de solicitud inválido |
| 404 | Contenedor no encontrado |
