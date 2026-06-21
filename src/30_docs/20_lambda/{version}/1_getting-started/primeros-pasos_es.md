<h1>Primeros pasos con Lambda</h1>

Esta guía te explica cómo crear tu primera función Lambda, invocarla e inspeccionar el resultado.

## Prerrequisitos

- Una cuenta de HolaCloud con credenciales API (Api-Key y Api-Secret).
- [curl](https://curl.se/) instalado en tu máquina.

## Paso 1: Obtén tus Credenciales API

Ingresa al panel de control de HolaCloud y dirígete a la sección de claves API. Genera un nuevo par de claves — recibirás un **Api-Key** y un **Api-Secret**. Guarda estos valores de forma segura, ya que se usan para autenticar todas las solicitudes administrativas.

## Paso 2: Crea una Función Lambda

Crea una función simple "hola-mundo" escrita en JavaScript:

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hola-mundo",
    "runtime": "js",
    "code": "export default (req) => { return { body: \"¡Hola, Mundo!\" } }"
  }'
```

Respuesta esperada:

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "name": "hola-mundo",
  "runtime": "js",
  "status": "active",
  "created_at": "2025-06-21T12:00:00Z"
}
```

Guarda el `id` devuelto — lo necesitarás para invocar y administrar tu función.

## Paso 3: Lista tus Lambdas

Verifica que la función se haya creado listando todas las lambdas de tu cuenta:

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

## Paso 4: Ejecuta la Función

Invoca la función de forma síncrona usando el endpoint administrativo:

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/ID_DE_TU_FUNCION" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{}'
```

Respuesta esperada:

```json
{
  "body": "¡Hola, Mundo!",
  "status_code": 200
}
```

También puedes invocar la función desde un endpoint público (sin autenticación):

```bash
curl -X POST "https://api.hola.cloud/run/ID_DE_TU_FUNCION" \
  -H "Content-Type: application/json" \
  -d '{}'
```

## Paso 5: Consulta Invocaciones Activas

Lista las invocaciones en ejecución usando el endpoint público:

```bash
curl "https://api.hola.cloud/ongoing"
```

## Siguientes Pasos

- Aprende a actualizar el código de tu función y configurar variables de entorno en [Gestionando Funciones Lambda](../2_managing-functions/manejando-funciones_es.md).
- Explora los patrones de invocación, incluyendo webhooks y el enrutador mux, en [Invocando Funciones Lambda](../3_invoking-functions/invocando-funciones_es.md).
