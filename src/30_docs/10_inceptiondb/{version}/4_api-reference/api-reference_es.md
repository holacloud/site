
# Referencia de la API de InceptionDB

URL Base: `https://api.hola.cloud`

## Autenticación

Los endpoints de administración de InceptionDB requieren autenticación mediante dos cabeceras:

- `Api-Key` — Su clave de API (UUID)
- `Api-Secret` — Su secreto de API (UUID)
- `X-Project` — Limita la solicitud a un proyecto específico

Todas las solicitudes deben enviarse a través de HTTPS.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/v1/databases` | Listar todas las bases de datos |
| POST | `/v1/databases` | Crear una nueva base de datos |
| GET | `/v1/databases/{id}` | Obtener una base de datos por ID |
| DELETE | `/v1/databases/{id}` | Eliminar una base de datos |
| PATCH | `/v1/databases/{id}` | Actualizar una base de datos |
| GET | `/v1/databases/{id}/collections` | Listar colecciones en una base de datos |
| POST | `/v1/databases/{id}/collections` | Crear una colección |
| POST | `/v1/databases/{id}/collections/{col}` | Operaciones con documentos (insertar, buscar, eliminar, modificar) |

## Códigos de Error Comunes

| Código | Descripción |
|--------|-------------|
| 400 | Solicitud incorrecta — sintaxis mal formada o parámetros inválidos |
| 401 | No autorizado — credenciales de API faltantes o inválidas |
| 403 | Prohibido — las credenciales no tienen acceso al recurso |
| 404 | No encontrado — el recurso solicitado no existe |
| 409 | Conflicto — el recurso ya existe |
| 500 | Error interno del servidor |
