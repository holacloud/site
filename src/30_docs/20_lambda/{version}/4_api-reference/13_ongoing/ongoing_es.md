# Listar Invocaciones en Curso

Devuelve una lista de las invocaciones de lambda actualmente en ejecución para la cuenta autenticada.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud HTTP

```http
GET /api/v0/ongoing HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: TU_TOKEN
```

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/api/v0/ongoing" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

## Respuesta

```json
[
  {
    "id": "inv_abc123",
    "lambda_id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
    "lambda_name": "hello-world",
    "started_at": "2025-07-01T12:00:00Z",
    "method": "POST",
    "path": "/hello-world"
  }
]
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Autenticación faltante o inválida |
