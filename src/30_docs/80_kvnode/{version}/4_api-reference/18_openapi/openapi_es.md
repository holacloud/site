# Especificación OpenAPI

Devuelve la especificación OpenAPI para la API KVNode.

## Autenticación

Este endpoint es público. No requiere autenticación.

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/openapi.json"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "openapi": "3.0.0",
  "info": {
    "title": "KVNode API",
    "version": "1.0.0"
  },
  "paths": {
    "/healthz": {
      "get": {
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "Node is healthy"
          }
        }
      }
    }
  }
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 500 | internal_error | Error interno del servidor |
