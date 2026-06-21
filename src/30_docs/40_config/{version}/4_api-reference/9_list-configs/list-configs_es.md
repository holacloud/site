# GET /v0/configs

Listar todas las configuraciones. Este endpoint es de acceso público (API de administración).

## Autenticación

No requiere autenticación. Este es un endpoint público.

## Parámetros de Consulta

| Parámetro | Tipo | Valor por Defecto | Descripción |
|-----------|------|-------------------|-------------|
| `limit` | integer | 100 | Número máximo de configuraciones a devolver |
| `offset` | integer | 0 | Número de configuraciones a omitir |
| `project` | string | "" | Filtrar configuraciones por nombre de proyecto |

## Solicitud

```bash
curl "https://api.hola.cloud/v0/configs?limit=10&offset=0"
```

## Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "configs": [
    {
      "id": "cfg_abc123",
      "project": "my-app",
      "environment": "production",
      "data": {
        "database": {
          "host": "db.example.com",
          "port": 5432
        }
      },
      "createdAt": "2026-06-21T10:00:00Z",
      "updatedAt": "2026-06-21T10:00:00Z"
    }
  ],
  "total": 1
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 500 | Internal Server Error | Ocurrió un error inesperado |
