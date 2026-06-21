# POST /api/console/start

Inicia un nuevo contenedor a partir de una imagen especificada.

## Descripción

Crea e inicia un contenedor a partir de la imagen dada con variables de entorno, volúmenes y mapeos de puertos opcionales.

## Autenticación

Ninguna. Este endpoint es público.

## Cuerpo de la Solicitud

```json
{
  "image": "my-project/my-app:latest",
  "env": {
    "LOG_LEVEL": "debug",
    "DATABASE_URL": "postgres://user:pass@host:5432/db"
  },
  "volumes": [
    {
      "container_path": "/data",
      "size_gb": 10
    }
  ],
  "ports": {
    "80/tcp": 8080
  }
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "my-project/my-app:latest",
    "env": {"LOG_LEVEL": "debug"},
    "ports": {"80/tcp": 8080}
  }'
```

## Respuesta

```json
{
  "container_id": "abc123def456",
  "image": "my-project/my-app:latest",
  "status": "running",
  "ports": {"80/tcp": 8080},
  "started_at": "2026-06-21T10:00:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Contenedor iniciado correctamente |
| 400 | Cuerpo de solicitud inválido |
| 404 | Imagen no encontrada |
