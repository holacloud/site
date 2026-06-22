# Tráfico Estadísticas

Devuelve estadísticas de tráfico en tiempo real de la puerta de enlace.

## Descripción

Este endpoint proporciona métricas agregadas sobre las solicitudes que pasan por la puerta de enlace Glue2, incluyendo tasas de solicitud, percentiles de latencia y distribución de códigos de estado.

## Autenticación

Ninguna. Este endpoint es público.

## Solicitud

No se requiere cuerpo en la solicitud.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/v0/stats"
```

## Respuesta

```json
{
  "total_requests": 154203,
  "requests_per_second": 42.5,
  "latency_p50_ms": 12,
  "latency_p95_ms": 45,
  "latency_p99_ms": 120,
  "status_codes": {
    "2xx": 148000,
    "4xx": 5200,
    "5xx": 1003
  },
  "active_connections": 87,
  "uptime_seconds": 86400
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Estadísticas devueltas correctamente |
