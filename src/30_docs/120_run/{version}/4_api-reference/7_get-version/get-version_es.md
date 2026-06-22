# Obtener Versión

Devuelve la versión actual del servicio Run.

## Descripción

Proporciona información de versión y compilación sobre la instancia del servicio Run en ejecución.

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
  "service": "run",
  "version": "1.5.2",
  "commit": "e5f6g7h8",
  "build_time": "2026-06-20T12:00:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Información de versión devuelta correctamente |
