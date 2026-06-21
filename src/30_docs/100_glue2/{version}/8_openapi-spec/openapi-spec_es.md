# GET /openapi.json

Devuelve la especificación OpenAPI de la puerta de enlace Glue2.

## Descripción

Este endpoint sirve el documento de especificación OpenAPI 3.0 que describe todas las rutas, formatos de solicitud y esquemas de respuesta expuestos por la puerta de enlace. La especificación se genera automáticamente a partir de las definiciones de ruta.

## Autenticación

Ninguna. Este endpoint es público.

## Solicitud

No se requiere cuerpo en la solicitud.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/openapi.json"
```

## Respuesta

```json
{
  "openapi": "3.0.3",
  "info": {
    "title": "Glue2 API Gateway",
    "version": "2.3.1",
    "description": "Puerta de enlace central de API para servicios HolaCloud"
  },
  "paths": {
    "/version": {
      "get": {
        "summary": "Obtener versión de la puerta de enlace",
        "responses": {
          "200": {
            "description": "Información de versión"
          }
        }
      }
    },
    "/v0/virtualhosts": {
      "get": {
        "summary": "Listar hosts virtuales",
        "responses": {
          "200": {
            "description": "Tabla de enrutamiento"
          }
        }
      }
    },
    "/v0/stats": {
      "get": {
        "summary": "Obtener estadísticas de tráfico",
        "responses": {
          "200": {
            "description": "Estadísticas de tráfico"
          }
        }
      }
    },
    "/v0/status": {
      "get": {
        "summary": "Estado de salud del backend",
        "responses": {
          "200": {
            "description": "Estado del servicio"
          }
        }
      }
    }
  }
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Especificación OpenAPI devuelta correctamente |
