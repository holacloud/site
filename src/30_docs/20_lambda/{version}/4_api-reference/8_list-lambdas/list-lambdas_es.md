
# GET /api/v0/lambdas

Lista todas las funciones lambda pertenecientes a la cuenta autenticada.

## Autenticación

Requiere las cabeceras `Api-Key` y `Api-Secret`.

## Solicitud HTTP

```http
GET /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf"
```

## Respuesta

```json
[
  {
    "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
    "name": "hello-world",
    "runtime": "javascript",
    "active": true,
    "created_at": "2025-05-10T12:00:00Z"
  },
  {
    "id": "a2b3c4d5-e6f7-8901-bcde-f12345678901",
    "name": "process-orders",
    "runtime": "python",
    "active": true,
    "created_at": "2025-06-20T08:30:00Z"
  }
]
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Cabeceras de autenticación faltantes o inválidas |
