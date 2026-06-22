# Primeros pasos

Run expone un subconjunto Docker Registry v2 orientado a push y una API de Console para estado de repositorios, start, stop, rollback, variables de entorno y volúmenes.

## Autenticación

Los endpoints de registry usan Basic Auth configurado con `docker login`.

```bash
docker login run.hola.cloud
```

## Construir y subir una imagen

```bash
docker build -t run.hola.cloud/my-project/my-app:v1 .
docker push run.hola.cloud/my-project/my-app:v1
```

## Leer datos de Console

```bash
curl "https://api.hola.cloud/api/console?repository=my-project/my-app"
```

## Iniciar un repositorio

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "reference": "v1"
  }'
```

También puedes usar un digest:

```json
{
  "repository": "my-project/my-app",
  "digest": "sha256:a1b2c3d4..."
}
```

## Detener un repositorio

```bash
curl -X POST "https://api.hola.cloud/api/console/stop" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app"}'
```

## Rollback

```bash
curl -X POST "https://api.hola.cloud/api/console/rollback" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "reference": "v0"}'
```

## Variables de entorno

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "env": [
      {"key": "LOG_LEVEL", "desired_value": "debug"}
    ]
  }'
```

## Volúmenes

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "volumes": [
      {"name": "my-data", "target": "/data", "mode": "rw"}
    ]
  }'
```
