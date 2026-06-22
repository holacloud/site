<h1>Primeros Pasos con Lambda</h1>

Esta guía muestra cómo crear una lambda JavaScript, listarla e invocarla mediante rutas administrativas y públicas.

## Requisitos Previos

- Una cuenta de HolaCloud con un token `X-Glue-Authentication`.
- [curl](https://curl.se/) instalado en tu máquina.

## Paso 1: Crear una Lambda

Crea una lambda simple `hello-world`:

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: TU_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "language": "javascript",
    "method": "GET",
    "path": "/hello-world",
    "code": "export default (req) => ({ body: { message: \"Hello, World!\", method: req.method, path: req.path } })"
  }'
```

Respuesta esperada:

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "created_timestamp": 1750507200,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-world",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Hello, World!\", method: req.method, path: req.path } })",
  "method": "GET",
  "path": "/hello-world"
}
```

Guarda el `id` devuelto; es el `lambda_id` usado por los endpoints de ejecución, consulta, actualización y eliminación.

## Paso 2: Listar Lambdas

Verifica que la lambda fue creada:

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

## Paso 3: Ejecutar la Lambda

Invócala mediante la ruta administrativa autenticada:

```bash
curl -X GET "https://api.hola.cloud/api/v0/run/TU_LAMBDA_ID" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

Respuesta esperada:

```json
{
  "body": {
    "message": "Hello, World!",
    "method": "GET",
    "path": "/hello-world"
  }
}
```

También puedes invocar la lambda mediante su ruta pública:

```bash
curl -X GET "https://api.hola.cloud/run/TU_LAMBDA_ID"
```

## Paso 4: Revisar Invocaciones en Curso

Lista las invocaciones que se están ejecutando:

```bash
curl "https://api.hola.cloud/ongoing"
```

## Siguientes Pasos

- Aprende a actualizar código y campos de ruta en [Manejando Funciones Lambda](../2_managing-functions/manejando-funciones_es.md).
- Explora rutas públicas de ejecución, mux, `/me` y `/openapi` en [Invocando Funciones Lambda](../3_invoking-functions/invocando-funciones_es.md).
