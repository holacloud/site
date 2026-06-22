# Primeros Pasos

Esta guía te muestra las operaciones básicas de KVNode: verificación de salud, gestión de colecciones y operaciones con clave-valor.

## Verificación de Salud

Verifica que el nodo esté funcionando:

```bash
curl "https://api.hola.cloud/healthz"
```

Respuesta esperada:

```json
{"ok":true,"node":"node-1","role":"primary"}
```

## Verificación de Disponibilidad

Comprueba si el nodo está listo para recibir tráfico (verifica la conexión con el nodo padre en réplicas):

```bash
curl "https://api.hola.cloud/readyz"
```

Respuesta esperada cuando está listo:

```json
{"ok":true,"node":"node-1","role":"primary","ready":true,"checks":{"wal_replayed":true,"parent_connected":true}}
```

## Crear una Colección

Las colecciones son contenedores de pares clave-valor. Crea una con una solicitud POST:

```bash
curl -X POST "https://api.hola.cloud/v1/collections" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret" \
  -d '{"name": "mi-coleccion"}'
```

Respuesta esperada:

```json
{"ok":true,"collection":"mi-coleccion"}
```

## Establecer un Par Clave-Valor

Almacena un valor JSON bajo una clave dentro de una colección:

```bash
curl -X POST "https://api.hola.cloud/v1/collections/mi-coleccion/keys/usuario:alice" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret" \
  -d '{"value": {"usuario": "alice", "rol": "admin"}}'
```

Respuesta esperada:

```json
{"ok":true,"seq":1,"version":1}
```

## Obtener una Clave

Recupera el valor de una clave específica:

```bash
curl "https://api.hola.cloud/v1/collections/mi-coleccion/keys/usuario:alice" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Respuesta esperada:

```json
{"key":"usuario:alice","value":{"usuario":"alice","rol":"admin"},"version":1,"updatedAt":"2025-01-01T00:00:00Z"}
```

## Listar Claves con Prefijo

Lista todas las claves en una colección, opcionalmente filtradas por prefijo y limitadas en cantidad:

```bash
curl "https://api.hola.cloud/v1/collections/mi-coleccion/keys?prefix=usuario:&limit=10" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Respuesta esperada:

```json
[{"key":"usuario:alice","value":{"usuario":"alice","rol":"admin"},"version":1}]
```

## Eliminar una Clave

Elimina una clave de una colección:

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/mi-coleccion/keys/usuario:alice" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Respuesta esperada:

```json
{"ok":true,"seq":2,"version":2}
```

## Resumen

Has realizado las operaciones básicas de KVNode: verificación de salud, creación de colecciones y CRUD completo sobre pares clave-valor. A partir de aquí, explora la configuración de replicación y la integración con los SDKs multi-lenguaje.
