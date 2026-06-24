# Especificación OpenAPI

Devuelve el documento de especificación OpenAPI para la API Lambda.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud HTTP

```http
GET /api/v0/openapi HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: TU_TOKEN
```

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/api/v0/openapi" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

## Respuesta

Devuelve la especificación OpenAPI completa como un documento JSON.

```json
{
  "openapi": "3.0.3",
  "info": {
    "title": "Lambda API",
    "version": "1.0.0"
  },
  "paths": {
    "/api/v0/lambdas": {
      "get": {
        "summary": "List lambdas",
        "responses": {
          "200": {
            "description": "List of lambdas"
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
| 401 | Autenticación faltante o inválida |
