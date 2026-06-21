# GET /v0/status

Devuelve el estado de salud de todos los servicios backend.

## Descripción

Este endpoint realiza verificaciones de salud contra cada servicio backend registrado e informa su estado actual. Es utilizado por los sistemas de monitoreo y el panel de la Consola.

## Autenticación

Ninguna. Este endpoint es público.

## Solicitud

No se requiere cuerpo en la solicitud.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/v0/status"
```

## Respuesta

```json
{
  "services": [
    {
      "name": "inceptiondb",
      "status": "healthy",
      "latency_ms": 5,
      "last_checked": "2026-06-21T10:00:00Z"
    },
    {
      "name": "lambda",
      "status": "healthy",
      "latency_ms": 12,
      "last_checked": "2026-06-21T10:00:00Z"
    },
    {
      "name": "files",
      "status": "degraded",
      "latency_ms": 350,
      "last_checked": "2026-06-21T10:00:00Z"
    },
    {
      "name": "kvnode",
      "status": "healthy",
      "latency_ms": 3,
      "last_checked": "2026-06-21T10:00:00Z"
    }
  ],
  "overall": "degraded"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Información de estado devuelta correctamente |
