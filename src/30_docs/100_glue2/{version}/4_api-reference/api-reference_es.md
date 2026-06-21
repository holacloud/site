# Referencia de la API de Glue2

Glue2 es la puerta de enlace central de API para HolaCloud. La mayoría de los endpoints son públicos, excepto aquellos bajo `/v0/secret` que requieren autenticación glueauth.Require.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| `GET` | `/version` | Devuelve la versión de la puerta de enlace |
| `GET` | `/v0/virtualhosts` | Lista la tabla de enrutamiento |
| `GET` | `/v0/stats` | Obtiene estadísticas de tráfico |
| `GET` | `/v0/status` | Verifica el estado de los servicios backend |
| `GET` | `/openapi.json` | Especificación OpenAPI |
