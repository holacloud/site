# GET /version

Devuelve la versión actual de la puerta de enlace Glue2.

## Descripción

El endpoint de versión proporciona información de compilación y lanzamiento sobre la instancia de la puerta de enlace en ejecución.

## Autenticación

Ninguna. Este endpoint es público.

## Solicitud

No se requiere cuerpo en la solicitud.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/version"
```

## Respuesta

```json
{
  "service": "glue2",
  "version": "2.3.1",
  "commit": "a1b2c3d4",
  "build_time": "2026-06-20T12:00:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Información de versión devuelta correctamente |
