<h1>Invocando Funciones Lambda</h1>

HolaCloud Lambda ofrece múltiples métodos de invocación para adaptarse a diferentes casos de uso — desde llamadas síncronas directas hasta integraciones con webhooks y enrutamiento de dominios personalizados.

## Invocación Síncrona

Todas las invocaciones son síncronas: HolaCloud ejecuta tu función y devuelve la respuesta una vez que la ejecución finaliza.

### Invocación Administrativa (Autenticada)

Usa el endpoint administrativo para llamadas internas donde controlas el cliente:

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/ID_DE_TU_FUNCION" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "clave": "valor"
  }'
```

### Invocación Pública (Sin Autenticación)

Usa el endpoint público para servicios externos o aplicaciones del lado del cliente:

```bash
curl -X POST "https://api.hola.cloud/run/ID_DE_TU_FUNCION" \
  -H "Content-Type: application/json" \
  -d '{
    "clave": "valor"
  }'
```

El endpoint público no requiere autenticación, lo que lo hace ideal para webhooks y APIs públicas.

## Uso como Webhook

Cualquier función Lambda puede funcionar como endpoint de webhook. Configura tu servicio externo (Stripe, GitHub, SendGrid, etc.) para que envíe eventos a:

```
https://api.hola.cloud/run/ID_DE_TU_FUNCION
```

La función recibe el payload del webhook como cuerpo de la solicitud y puede devolver una respuesta que el servicio externo interpretará.

Ejemplo: webhook de GitHub que registra eventos de push:

```javascript
export default (req) => {
  const evento = req.headers["x-github-event"];
  const payload = req.body;

  console.log(`Recibido ${evento} desde GitHub`);

  return {
    status_code: 200,
    body: { recibido: true }
  };
};
```

## El Enrutador Mux

El enrutador mux asigna dominios o subdominios personalizados a funciones Lambda específicas por propietario. El patrón de ruta es:

```
GET /mux/{owner_id}/*
```

Cuando una solicitud llega al mux, HolaCloud la enruta a la función Lambda correspondiente a ese usuario. Esto permite escenarios como:

- Plataformas SaaS multiinquilino donde cada inquilino tiene un subdominio
- Mapeo de dominios personalizados para funciones desplegadas por usuarios
- Enrutamiento basado en rutas a diferentes funciones bajo el mismo propietario

```bash
curl "https://api.hola.cloud/mux/ID_DEL_PROPIETARIO/cualquier/ruta/aqui" \
  -H "Content-Type: application/json"
```

El `owner_id` es el identificador de la cuenta de HolaCloud. El resto de la ruta se reenvía a la función, que lo recibe como parte de `req`.

## Monitoreo de Ejecuciones en Curso

Consulta qué funciones se están ejecutando actualmente usando el endpoint público `ongoing`:

```bash
curl "https://api.hola.cloud/ongoing"
```

Respuesta esperada:

```json
[
  {
    "function_id": "f3b2c1a0-1234-5678-9abc-def012345678",
    "function_name": "hello-world",
    "started_at": "2025-06-21T12:00:00Z",
    "duration_ms": 234
  }
]
```

Este endpoint es útil para monitoreo operativo y depuración de funciones de larga duración.

## Endpoint de Usuario Actual

El endpoint `/me` devuelve información sobre el usuario autenticado:

```bash
curl "https://api.hola.cloud/me" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

## Especificación OpenAPI

HolaCloud expone una especificación OpenAPI legible por máquina en:

```
GET https://api.hola.cloud/openapi
```

Puede usarse con herramientas como Swagger UI, Postman o generadores de código para explorar e interactuar con la API de Lambda de forma programática.
