# OpenAPI Specification

Devuelve la especificación OpenAPI del servicio Run.

## Solicitud

```http
GET /openapi.json
```

## Ejemplo

```bash
curl "https://api.hola.cloud/openapi.json"
```

## Respuesta

Devuelve la especificación OpenAPI como JSON.

```json
{
  "openapi": "3.0.0",
  "info": {
    "title": "HolaCloud Run API",
    "version": "1.0.0"
  },
  "paths": {}
}
```
