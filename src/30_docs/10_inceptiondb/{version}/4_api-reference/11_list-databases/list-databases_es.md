
# GET /v1/databases

Lista todas las bases de datos accesibles para el proyecto autenticado.

## Autenticación

Requiere las cabeceras `Api-Key`, `Api-Secret` y `X-Project`.

## Solicitud HTTP

```http
GET /v1/databases HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
```

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/v1/databases" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp"
```

## Respuesta

```json
[
  {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "production-db",
    "created_at": "2025-06-15T10:30:00Z",
    "collection_count": 12
  },
  {
    "id": "b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "name": "staging-db",
    "created_at": "2025-06-10T08:00:00Z",
    "collection_count": 5
  }
]
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 403 | Acceso denegado al proyecto |
