# Run

Run es el servicio de imágenes de contenedor y control de Console de HolaCloud. Expone un subconjunto orientado a push de Docker Registry v2 y una API pequeña de Console para inspeccionar repositorios y cambiar configuración de ejecución.

## Funcionalidades

### Subconjunto Registry v2 orientado a push

Run soporta los endpoints necesarios para subir blobs y manifiestos bajo `/v2`. No es una implementación completa de Docker Registry ni debe documentarse como registry general de pull.

```bash
docker login run.hola.cloud
docker build -t run.hola.cloud/my-project/my-app:latest .
docker push run.hola.cloud/my-project/my-app:latest
```

### API de Console

La API de Console trabaja con repositorios y referencias o digests de imagen:

- `GET /version`
- `GET /api/console?repository=`
- `POST /api/console/start`
- `POST /api/console/stop`
- `POST /api/console/rollback`
- `PUT /api/console/env`
- `PUT /api/console/volumes`

No existe `/v1/run/deploy`, `/api/console/run`, API de push/exec ni flujo con `container_id`.

## Primeros pasos

```bash
docker login run.hola.cloud
docker build -t run.hola.cloud/my-project/my-app:latest .
docker push run.hola.cloud/my-project/my-app:latest
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "reference": "latest"}'
```
