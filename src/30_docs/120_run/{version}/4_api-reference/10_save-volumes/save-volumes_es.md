# PUT /api/console/volumes

Guarda la configuración de volúmenes para un contenedor.

## Descripción

Actualiza la configuración de volúmenes persistentes para un contenedor especificado. Los volúmenes están respaldados por el almacenamiento distribuido de HolaCloud y sobreviven a los reinicios del contenedor.

## Autenticación

Ninguna. Este endpoint es público.

## Cuerpo de la Solicitud

```json
{
  "container_id": "abc123def456",
  "volumes": [
    {
      "container_path": "/data",
      "size_gb": 20,
      "mount_path": "/mnt/data"
    },
    {
      "container_path": "/config",
      "size_gb": 5,
      "mount_path": "/mnt/config"
    }
  ]
}
```

## Ejemplo

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Content-Type: application/json" \
  -d '{
    "container_id": "abc123def456",
    "volumes": [
      {"container_path": "/data", "size_gb": 20}
    ]
  }'
```

## Respuesta

```json
{
  "container_id": "abc123def456",
  "volumes": [
    {
      "volume_id": "vol-001",
      "container_path": "/data",
      "size_gb": 20,
      "mount_path": "/mnt/data",
      "status": "active"
    }
  ],
  "updated": true
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Configuración de volúmenes guardada correctamente |
| 400 | Cuerpo de solicitud inválido |
| 404 | Contenedor no encontrado |
