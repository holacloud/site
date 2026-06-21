# Administrando Configuraciones

La API de administración (`/v0/configs`) proporciona operaciones CRUD completas para gestionar configuraciones. Estos endpoints son de acceso público y no requieren clave API.

## Listar Configuraciones

Obtén una lista de todas las configuraciones almacenadas:

```bash
curl "https://api.hola.cloud/v0/configs"
```

Respuesta esperada:

```json
[
  {
    "id": "cfg_abc123",
    "project": "my-app",
    "environment": "production",
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-06-20T14:22:00Z"
  },
  {
    "id": "cfg_def456",
    "project": "my-app",
    "environment": "staging",
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-06-19T09:00:00Z"
  }
]
```

## Crear una Configuración

Crea una nueva entrada de configuración:

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{
    "project": "my-app",
    "environment": "production",
    "config": {
      "database": {
        "host": "db.example.com",
        "port": 5432
      }
    }
  }'
```

Respuesta esperada:

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "config": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    }
  },
  "created_at": "2025-06-20T14:22:00Z",
  "updated_at": "2025-06-20T14:22:00Z"
}
```

## Obtener una Configuración

Recupera una configuración por su ID:

```bash
curl "https://api.hola.cloud/v0/configs/cfg_abc123"
```

Respuesta esperada:

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "config": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    }
  },
  "created_at": "2025-01-15T10:30:00Z",
  "updated_at": "2025-06-20T14:22:00Z"
}
```

## Actualizar Parcialmente una Configuración

Aplica una actualización parcial a una configuración existente:

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{
    "config": {
      "database": {
        "host": "db-nuevo.example.com"
      }
    }
  }'
```

Respuesta esperada:

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "config": {
    "database": {
      "host": "db-nuevo.example.com",
      "port": 5432
    }
  },
  "created_at": "2025-01-15T10:30:00Z",
  "updated_at": "2025-06-20T15:00:00Z"
}
```

## Eliminar una Configuración

Elimina una configuración permanentemente:

```bash
curl -X DELETE "https://api.hola.cloud/v0/configs/cfg_abc123"
```

Respuesta esperada: `HTTP 204 No Content`

## Referencia de Solicitudes HTTP

```http
GET /v0/configs HTTP/1.1
Host: api.hola.cloud

POST /v0/configs HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "project": "my-app",
  "environment": "production",
  "config": { ... }
}

GET /v0/configs/{id} HTTP/1.1
Host: api.hola.cloud

PATCH /v0/configs/{id} HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "config": { ... }
}

DELETE /v0/configs/{id} HTTP/1.1
Host: api.hola.cloud
```
