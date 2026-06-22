# Referencia de la API Lambda

URL base: `https://api.hola.cloud`

## Autenticación

Los endpoints de administración requieren la cabecera `X-Glue-Authentication`. Los endpoints públicos de invocación (`/run/{lambda_id}` y `/mux/{owner_id}/*`) no requieren autenticación.

Todas las solicitudes deben enviarse por HTTPS.

## Objeto Lambda

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | string | Identificador de la lambda |
| `created_timestamp` | number | Fecha de creación como timestamp Unix |
| `owner` | string | Identificador del propietario |
| `project_id` | string | Identificador del proyecto |
| `name` | string | Nombre de la lambda |
| `language` | string | Uno de `javascript`, `static-html`, `static-css` o `static-js` |
| `code` | string | Código fuente o contenido estático |
| `method` | string | Método HTTP asociado a la lambda |
| `path` | string | Path HTTP asociado a la lambda |

## Endpoints

| Método | Path | Descripción |
|--------|------|-------------|
| POST | `/api/v0/lambdas` | Crear una lambda |
| GET | `/api/v0/lambdas` | Listar lambdas |
| GET | `/api/v0/lambdas/{lambda_id}` | Obtener una lambda por ID |
| PATCH | `/api/v0/lambdas/{lambda_id}` | Modificar una lambda |
| DELETE | `/api/v0/lambdas/{lambda_id}` | Eliminar una lambda |
| ANY | `/api/v0/run/{lambda_id}` | Ejecutar una lambda con autenticación |
| ANY | `/run/{lambda_id}` | Ejecutar una lambda públicamente |
| ANY | `/mux/{owner_id}/*` | Enrutar una solicitud dentro del alcance de un propietario |
| GET | `/ongoing` | Listar invocaciones en curso |
| GET | `/me` | Devolver el usuario autenticado |
| GET | `/openapi` | Devolver el documento OpenAPI |

## Códigos de Error Comunes

| Código | Descripción |
|--------|-------------|
| 400 | Solicitud inválida |
| 401 | Autenticación faltante o inválida |
| 403 | Prohibido |
| 404 | Recurso no encontrado |
| 500 | Error interno del servidor |
