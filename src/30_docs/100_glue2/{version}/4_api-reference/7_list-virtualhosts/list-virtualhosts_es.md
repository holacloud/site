# Listar Virtualhosts

Devuelve la tabla de enrutamiento de hosts virtuales actual.

## Descripción

Este endpoint lista todos los hosts virtuales configurados y sus servicios backend correspondientes. Es utilizado por la Consola de HolaCloud para mostrar la configuración de enrutamiento.

## Autenticación

Ninguna. Este endpoint es público.

## Solicitud

No se requiere cuerpo en la solicitud.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/v0/virtualhosts"
```

## Respuesta

```json
{
  "virtualhosts": [
    {
      "host": "my-project.hola.cloud",
      "target": "inceptiondb",
      "port": 8080,
      "tls": true
    },
    {
      "host": "api.my-project.hola.cloud",
      "target": "lambda",
      "port": 8081,
      "tls": true
    }
  ]
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Tabla de enrutamiento devuelta correctamente |
