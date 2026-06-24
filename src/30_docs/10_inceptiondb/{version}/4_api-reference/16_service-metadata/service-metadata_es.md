# Metadatos del servicio

Estos endpoints exponen metadatos y diagnósticos del servicio.

## Árbol de rutas

```http
GET /doc HTTP/1.1
Host: api.hola.cloud
```

Devuelve un árbol de texto con las rutas registradas por el servicio.

## Raíz de la API

```http
GET /v1 HTTP/1.1
Host: api.hola.cloud
```

Devuelve la respuesta de bienvenida de v1.

## Release

```http
GET /release HTTP/1.1
Host: api.hola.cloud
```

Devuelve metadatos de release del servicio desplegado.

## OpenAPI

```http
GET /openapi.json HTTP/1.1
Host: api.hola.cloud
```

Devuelve el documento OpenAPI servido por la API.

## Debug de bases de datos

```http
GET /v1/debug/dbs HTTP/1.1
Host: api.hola.cloud
```

Devuelve un listado debug de bases de datos. Este endpoint está pensado para diagnósticos, no para tráfico de aplicación.
