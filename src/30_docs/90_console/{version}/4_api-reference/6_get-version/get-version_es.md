# Obtener Versión

Devuelve la versión actual del servicio Consola.

## Descripción

Este endpoint devuelve la versión configurada de la Consola como texto plano. No requiere autenticación.

## Autenticación

Ninguna. Este endpoint es público.

## Solicitud

No se requiere cuerpo en la solicitud.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/version"
```

## Respuesta

```text
1.0.0
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Información de versión devuelta correctamente |
