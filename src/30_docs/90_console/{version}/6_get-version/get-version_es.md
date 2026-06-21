# GET /version

Devuelve la versión actual del servicio Consola.

## Descripción

Este endpoint es utilizado por la interfaz de la Consola para mostrar la información de versión en el pie de página y las páginas de diagnóstico. No requiere autenticación.

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
  "version": "1.0.0",
  "build": "20260601-abcdef",
  "mode": "development"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Información de versión devuelta correctamente |
