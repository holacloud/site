# Estadísticas de Tráfico

Devuelve un arreglo de estadísticas por host. Cada elemento usa los campos `prettyStats` de Glue2 e incluye filas agregadas para `*` y `-noroute-`.

```bash
curl -X GET "https://api.hola.cloud/v0/stats"
```

```json
[
  {
    "host": "mi-proyecto.hola.cloud",
    "served_requests": 154203,
    "serving_time": "32.4s",
    "latency_avg": "210µs",
    "uptime": "24h0m0s",
    "start_timestamp": "2026-06-21T10:00:00Z"
  }
]
```
