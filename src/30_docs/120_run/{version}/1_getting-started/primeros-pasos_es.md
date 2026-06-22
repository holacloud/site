# Primeros Pasos

Run es un servicio de ejecución de contenedores con una API compatible con Docker Registry v2 incorporada. Permite desplegar, ejecutar y gestionar imágenes de contenedores con variables de entorno y volúmenes.

## Autenticación

Run utiliza **DockerAuth** con autenticación básica para la autenticación en el registro. Usa tu clave de API y secreto de HolaCloud como usuario y contraseña.

### Iniciar sesión en el Registro

```bash
docker login https://registry.hola.cloud
```

Se te solicitará tu clave de API de HolaCloud (usuario) y tu secreto de API (contraseña).

## Construir y Subir una Imagen

Crea una aplicación simple con un `Dockerfile`:

```dockerfile
FROM alpine:latest
COPY app.sh /app.sh
RUN chmod +x /app.sh
CMD ["/app.sh"]
```

Construye y etiqueta la imagen:

```bash
docker build -t registry.hola.cloud/mi-proyecto/mi-app:v1 .
```

Súbela al registro:

```bash
docker push registry.hola.cloud/mi-proyecto/mi-app:v1
```

## Iniciar un Contenedor

Inicia un contenedor desde una imagen subida mediante la API de Consola:

```bash
curl -X POST "https://api.hola.cloud/api/console/run" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "registry.hola.cloud/mi-proyecto/mi-app:v1",
    "env": {
      "LOG_LEVEL": "info",
      "DATABASE_URL": "postgres://..."
    },
    "volumes": [
      {
        "container_path": "/data",
        "volume_name": "mis-datos"
      }
    ]
  }'
```

Respuesta esperada:

```json
{
  "id": "c8a9f3b2-1234-5678-9abc-def012345678",
  "image": "registry.hola.cloud/mi-proyecto/mi-app:v1",
  "status": "running",
  "created_at": "2025-06-21T12:00:00Z"
}
```

## Detener un Contenedor en Ejecución

Detén un contenedor por su ID:

```bash
curl -X DELETE "https://api.hola.cloud/api/console/run/c8a9f3b2-1234-5678-9abc-def012345678" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

## Revertir a una Versión Anterior de la Imagen

Despliega una versión anterior especificando una etiqueta diferente:

```bash
curl -X PUT "https://api.hola.cloud/api/console/run/c8a9f3b2-1234-5678-9abc-def012345678" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "registry.hola.cloud/mi-proyecto/mi-app:v0"
  }'
```

## Gestionar Variables de Entorno

Actualiza las variables de entorno de un contenedor en ejecución:

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "env": {
      "LOG_LEVEL": "debug",
      "DATABASE_URL": "postgres://..."
    }
  }'
```

## Gestionar Volúmenes

Adjunta volúmenes a un contenedor:

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "volumes": [
      {
        "container_path": "/data",
        "volume_name": "mis-datos",
        "read_only": false
      }
    ]
  }'
```

## Consultar el Estado de Ejecución

Consulta el estado actual de todos los contenedores:

```bash
curl "https://api.hola.cloud/api/console/run" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

Respuesta esperada:

```json
[
  {
    "id": "c8a9f3b2-1234-5678-9abc-def012345678",
    "image": "registry.hola.cloud/mi-proyecto/mi-app:v1",
    "status": "running",
    "env": { "LOG_LEVEL": "info" },
    "volumes": [{ "container_path": "/data", "volume_name": "mis-datos" }],
    "created_at": "2025-06-21T12:00:00Z"
  }
]
```

## Flujo de Trabajo Completo

Un flujo de trabajo completo de principio a fin:

1. **Inicia sesión** en el registro:
   ```bash
   docker login https://registry.hola.cloud
   ```

2. **Construye** tu imagen:
   ```bash
   docker build -t registry.hola.cloud/mi-proyecto/mi-app:v1 .
   ```

3. **Sube** la imagen:
   ```bash
   docker push registry.hola.cloud/mi-proyecto/mi-app:v1
   ```

4. **Inicia** un contenedor:
   ```bash
   curl -X POST "https://api.hola.cloud/api/console/run" \
     -H "Api-Key: TU_API_KEY" \
     -H "Api-Secret: TU_API_SECRET" \
     -H "Content-Type: application/json" \
     -d '{"image": "registry.hola.cloud/mi-proyecto/mi-app:v1"}'
   ```

5. **Verifica** que esté funcionando:
   ```bash
   curl "https://api.hola.cloud/api/console/run" \
     -H "Api-Key: TU_API_KEY" \
     -H "Api-Secret: TU_API_SECRET"
   ```

6. **Detén** el contenedor cuando termines:
   ```bash
   curl -X DELETE "https://api.hola.cloud/api/console/run/CONTAINER_ID" \
     -H "Api-Key: TU_API_KEY" \
     -H "Api-Secret: TU_API_SECRET"
   ```
