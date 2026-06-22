<h1>Invocando Funciones Lambda</h1>

HolaCloud Lambda soporta llamadas administrativas directas, llamadas públicas, rutas mux por propietario y un pequeño conjunto de endpoints de inspección del servicio.

## Invocación Administrativa

Usa `/api/v0/run/{lambda_id}` cuando el cliente pueda enviar `X-Glue-Authentication`. El endpoint acepta cualquier método HTTP.

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/TU_LAMBDA_ID" \
  -H "X-Glue-Authentication: TU_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

## Invocación Pública

Usa `/run/{lambda_id}` para webhooks, llamadas desde navegador y otros clientes que no deben enviar credenciales administrativas. El endpoint acepta cualquier método HTTP.

```bash
curl -X POST "https://api.hola.cloud/run/TU_LAMBDA_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

La lambda recibe los datos de solicitud mediante `req`, incluyendo método, path, headers, valores de query y body.

## Uso como Webhook

Configura un servicio externo para enviar eventos a:

```text
https://api.hola.cloud/run/TU_LAMBDA_ID
```

Ejemplo de manejador de webhook:

```javascript
export default (req) => {
  return {
    body: {
      received: true,
      event: req.headers["x-github-event"],
      payload: req.body
    }
  };
};
```

## Mux Router

El mux router reenvía rutas por propietario mediante `/mux/{owner_id}/*`. Acepta cualquier método HTTP.

```bash
curl -X GET "https://api.hola.cloud/mux/OWNER_ID/cualquier/ruta/aqui"
```

El `owner_id` identifica al propietario de HolaCloud. El resto del path se reenvía a la lógica de rutas de la lambda.

## Invocaciones en Curso

Consulta las invocaciones que se están ejecutando:

```bash
curl "https://api.hola.cloud/ongoing"
```

## Usuario Actual

El endpoint `/me` devuelve información del usuario autenticado:

```bash
curl "https://api.hola.cloud/me" \
  -H "X-Glue-Authentication: TU_TOKEN"
```

## Especificación OpenAPI

HolaCloud expone el documento OpenAPI de Lambda en:

```text
GET https://api.hola.cloud/openapi
```
