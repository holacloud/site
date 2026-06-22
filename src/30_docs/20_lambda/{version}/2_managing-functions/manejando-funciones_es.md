<h1>Manejando Funciones Lambda</h1>

Después de crear una lambda, usa la API de administración para inspeccionarla, actualizar sus campos soportados, listar todas las lambdas de la cuenta o eliminarla.

## Estructura de la Función

Las lambdas JavaScript exportan un manejador por defecto. El manejador recibe un objeto de solicitud y devuelve el cuerpo de respuesta que HolaCloud debe enviar.

```javascript
export default (req) => {
  return {
    body: {
      method: req.method,
      path: req.path,
      headers: req.headers,
      data: req.body
    }
  };
};
```

Las lambdas estáticas usan uno de los modos de lenguaje estático: `static-html`, `static-css` o `static-js`. En esos modos, `code` es el contenido servido para la lambda correspondiente.

## Actualizar una Lambda

Usa `PATCH /api/v0/lambdas/{lambda_id}` para actualizar `name`, `language`, `code`, `method` o `path`.

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/TU_LAMBDA_ID" \
  -H "X-Glue-Authentication: TU_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-updated",
    "method": "POST",
    "path": "/hello-updated",
    "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })"
  }'
```

Respuesta esperada:

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "created_timestamp": 1750507200,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-updated",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })",
  "method": "POST",
  "path": "/hello-updated"
}
```

## Ver Detalles de una Lambda

Obtén una lambda por ID:

```bash
curl "https://api.hola.cloud/api/v0/lambdas/TU_LAMBDA_ID" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

## Listar Todas las Lambdas

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

## Eliminar una Lambda

Elimina una lambda permanentemente:

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/TU_LAMBDA_ID" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

Una eliminación correcta devuelve la respuesta de la API para la lambda eliminada o una carga de confirmación, según la versión desplegada de la API.
