
# Drop Database

Elimina permanentemente una base de datos con todas sus colecciones y documentos. Esta acción no se puede deshacer.

## Autenticación

Requiere las cabeceras `Api-Key`, `Api-Secret` y `X-Project`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | uuid | Identificador único de la base de datos a eliminar |

## Solicitud HTTP

```http
DELETE /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
```

## Ejemplo

```bash
curl -X DELETE "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp"
```

## Respuesta

```json
{
  "message": "Base de datos eliminada exitosamente",
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 403 | Acceso denegado al proyecto |
| 404 | Base de datos no encontrada |
