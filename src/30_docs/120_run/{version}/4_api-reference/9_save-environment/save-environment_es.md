# Guardar Entorno

Guarda las variables de entorno para un contenedor.

## Descripción

Actualiza la configuración de variables de entorno para un contenedor especificado. El contenedor debe ser reiniciado para que los cambios surtan efecto.

## Autenticación

Ninguna. Este endpoint es público.

## Cuerpo de la Solicitud

```json
{
  "container_id": "abc123def456",
  "env": {
    "LOG_LEVEL": "info",
    "DATABASE_URL": "postgres://user:pass@host:5432/db",
    "REDIS_URL": "redis://host:6379"
  }
}
```

## Ejemplo

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Content-Type: application/json" \
  -d '{
    "container_id": "abc123def456",
    "env": {
      "LOG_LEVEL": "info",
      "DATABASE_URL": "postgres://user:pass@host:5432/db"
    }
  }'
```

## Respuesta

```json
{
  "container_id": "abc123def456",
  "env": {
    "LOG_LEVEL": "info",
    "DATABASE_URL": "postgres://user:pass@host:5432/db",
    "REDIS_URL": "redis://host:6379"
  },
  "updated": true
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Variables de entorno guardadas correctamente |
| 400 | Cuerpo de solicitud inválido |
| 404 | Contenedor no encontrado |
