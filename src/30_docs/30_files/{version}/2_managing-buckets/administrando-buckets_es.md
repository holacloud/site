# Administrando Buckets

Los buckets son los contenedores fundamentales en HolaCloud Files. Esta guía cubre todas las operaciones con buckets: crear, listar, inspeccionar, modificar y eliminar.

## Reglas de Nomenclatura de Buckets

Al crear un bucket, el nombre debe seguir estas reglas:

- Debe tener entre 3 y 63 caracteres
- Solo puede contener letras minúsculas, números y guiones (-)
- Debe comenzar y terminar con una letra o número
- No debe tener el formato de una dirección IP
- Debe ser único a nivel global entre todos los clientes de HolaCloud

## Crear un Bucket

Cree un nuevo bucket con una solicitud POST:

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-bucket-de-activos",
    "metadata": {
      "environment": "production",
      "project": "mi-app"
    }
  }'
```

Solicitud HTTP equivalente:

```http
POST /v1/buckets HTTP/1.1
Host: api.hola.cloud
Api-Key: SU_API_KEY
Api-Secret: SU_API_SECRET
Content-Type: application/json

{
  "name": "mi-bucket-de-activos",
  "metadata": {
    "environment": "production",
    "project": "mi-app"
  }
}
```

Respuesta esperada (HTTP 201):

```json
{
  "id": "bkt_xyz789",
  "name": "mi-bucket-de-activos",
  "metadata": {
    "environment": "production",
    "project": "mi-app"
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## Listar Buckets

Liste todos los buckets en su cuenta:

```bash
curl "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

Respuesta esperada:

```json
[
  {
    "id": "bkt_abc123",
    "name": "mi-primer-bucket",
    "createdAt": "2026-06-21T09:00:00Z",
    "size": 1024,
    "fileCount": 3
  },
  {
    "id": "bkt_xyz789",
    "name": "mi-bucket-de-activos",
    "createdAt": "2026-06-21T10:00:00Z",
    "size": 0,
    "fileCount": 0
  }
]
```

## Obtener Detalles de un Bucket

Recupere información detallada sobre un bucket específico por su ID:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_xyz789" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

Respuesta esperada:

```json
{
  "id": "bkt_xyz789",
  "name": "mi-bucket-de-activos",
  "metadata": {
    "environment": "production",
    "project": "mi-app"
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "modifiedAt": "2026-06-21T10:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## Modificar un Bucket

Actualice los metadatos de un bucket usando PATCH:

```bash
curl -X PATCH "https://api.hola.cloud/v1/buckets/bkt_xyz789" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "metadata": {
      "environment": "staging"
    }
  }'
```

Respuesta esperada:

```json
{
  "id": "bkt_xyz789",
  "name": "mi-bucket-de-activos",
  "metadata": {
    "environment": "staging",
    "project": "mi-app"
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "modifiedAt": "2026-06-21T11:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## Eliminar un Bucket

Un bucket debe estar **vacío** (sin archivos) antes de poder eliminarlo.

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_xyz789" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

Respuesta esperada: HTTP 204 Sin Contenido.

### Error: Bucket No Vacío

Si intenta eliminar un bucket que aún contiene archivos, recibirá:

```json
{
  "error": {
    "code": "BUCKET_NOT_EMPTY",
    "message": "No se puede eliminar el bucket 'mi-bucket-de-activos': el bucket contiene 3 archivo(s). Elimine todos los archivos antes de eliminar el bucket."
  }
}
```

## Manejo de Errores

### Nombre de Bucket Duplicado

```json
{
  "error": {
    "code": "BUCKET_ALREADY_EXISTS",
    "message": "Ya existe un bucket con el nombre 'mi-bucket-de-activos'."
  }
}
```

### Nombre de Bucket Inválido

```json
{
  "error": {
    "code": "INVALID_BUCKET_NAME",
    "message": "El nombre del bucket debe tener entre 3 y 63 caracteres, contener solo letras minúsculas, números y guiones, y comenzar/terminar con una letra o número."
  }
}
```

### Bucket No Encontrado

```json
{
  "error": {
    "code": "BUCKET_NOT_FOUND",
    "message": "Bucket 'bkt_inexistente' no encontrado."
  }
}
```

## Resumen de Códigos de Estado HTTP

| Código | Descripción |
|--------|-------------|
| 200 | Éxito (GET, PATCH) |
| 201 | Creado (POST) |
| 204 | Sin Contenido (DELETE) |
| 400 | Solicitud Incorrecta -- entrada inválida |
| 404 | No Encontrado -- el bucket no existe |
| 409 | Conflicto -- el bucket ya existe o no está vacío |
| 500 | Error Interno del Servidor |
