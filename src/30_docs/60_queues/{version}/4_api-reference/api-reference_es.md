
# Referencia de la API de Tailon

URL base: `https://api.hola.cloud`

## Autenticación

Los endpoints de Tailon son **públicos** por defecto. No se requieren encabezados de autenticación para la mayoría de las operaciones.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/v1/queues` | Listar todas las colas |
| POST | `/v1/queues` | Crear una nueva cola |
| GET | `/v1/queues/{id}` | Obtener detalles de una cola |
| DELETE | `/v1/queues/{id}` | Eliminar una cola |
| POST | `/v1/queues/{id}` | Escribir mensajes en una cola |
| GET | `/v1/queues/{id}` | Leer mensajes de una cola (long-poll) |
| GET | `/v1/clients` | Listar clientes de streaming activos |

## Códigos de Error Comunes

| Código | Descripción |
|--------|-------------|
| 400 | Solicitud incorrecta — JSON malformado o parámetros inválidos |
| 404 | No encontrado — la cola solicitada no existe |
| 409 | Conflicto — conflicto en la operación de cola |
| 500 | Error interno del servidor |
