# Especificación OpenAPI

Devuelve la especificación OpenAPI para la API Scheduler.

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
  "openapi": "3.1.0",
  "info": {
    "title": "Scheduler API",
    "version": "1.0.0"
  },
  "paths": {}
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 500 | internal_error | Error interno del servidor |
