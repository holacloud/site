
# Referencia de la API de Lambda

URL Base: `https://api.hola.cloud`

## Autenticación

Los endpoints de administración requieren las cabeceras `Api-Key` y `Api-Secret`. Los endpoints de invocación pública (`/run/{id}` y `/mux/{owner}/*`) no requieren autenticación.

Todas las solicitudes deben enviarse a través de HTTPS.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/api/v0/lambdas` | Listar todas las lambdas |
| POST | `/api/v0/lambdas` | Crear una nueva lambda |
| GET | `/api/v0/lambdas/{id}` | Obtener una lambda por ID |
| PATCH | `/api/v0/lambdas/{id}` | Modificar una lambda |
| DELETE | `/api/v0/lambdas/{id}` | Eliminar una lambda |
| POST | `/api/v0/run/{id}` | Ejecutar una lambda (admin) |
| POST | `/run/{id}` | Ejecutar una lambda (público) |
| GET | `/mux/{owner}/*` | Enrutar a las lambdas del propietario |

## Códigos de Error Comunes

| Código | Descripción |
|--------|-------------|
| 400 | Solicitud incorrecta — sintaxis mal formada o parámetros inválidos |
| 401 | No autorizado — credenciales de API faltantes o inválidas |
| 403 | Prohibido — las credenciales no tienen acceso al recurso |
| 404 | No encontrado — el recurso solicitado no existe |
| 409 | Conflicto — el recurso ya existe |
| 500 | Error interno del servidor |
