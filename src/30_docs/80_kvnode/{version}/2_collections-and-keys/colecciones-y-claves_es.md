# Colecciones Y Claves

KVNode organiza los datos en **colecciones**, cada una contiene pares **clave-valor**. Este documento cubre la gestión de colecciones y las operaciones con claves en detalle.

## Gestión de Colecciones

### Crear una Colección

```bash
curl -X POST "https://api.hola.cloud/v1/collections" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret" \
  -d '{"name": "mi-coleccion"}'
```

Respuesta:

```json
{"ok":true,"collection":"mi-coleccion"}
```

Las colecciones se crean al instante. Si la colección ya existe, la operación es idempotente.

### Listar Colecciones

```bash
curl "https://api.hola.cloud/v1/collections" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Respuesta:

```json
{"collections":["mi-coleccion","config","sesiones"]}
```

### Eliminar una Colección

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/mi-coleccion" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Respuesta:

```json
{"ok":true,"collection":"mi-coleccion"}
```

Eliminar una colección borra todas las claves almacenadas en ella. Esta operación no se puede deshacer.

## Operaciones con Claves

Las claves de KVNode son cadenas de texto que pueden contener cualquier carácter UTF-8. Los valores deben ser JSON válido.

### Establecer una Clave

```bash
curl -X POST "https://api.hola.cloud/v1/collections/mi-coleccion/keys/config:tema" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret" \
  -d '{"value": {"modo": "oscuro", "tamanoFuente": 14}}'
```

La respuesta incluye el número de secuencia y la versión:

```json
{"ok":true,"seq":3,"version":1}
```

Cada escritura exitosa devuelve un `seq` (secuencia global) y un `version` (versión por clave) que aumentan monotónicamente, útiles para control de concurrencia optimista.

### Obtener una Clave

```bash
curl "https://api.hola.cloud/v1/collections/mi-coleccion/keys/config:tema" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Respuesta:

```json
{"key":"config:tema","value":{"modo":"oscuro","tamanoFuente":14},"version":1,"updatedAt":"2025-01-01T00:00:00Z"}
```

Intentar obtener una clave que no existe devuelve un error 404.

### Eliminar una Clave

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/mi-coleccion/keys/config:tema" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Respuesta:

```json
{"ok":true,"seq":4,"version":2}
```

### Listar Claves con Prefijo y Límite

```bash
curl "https://api.hola.cloud/v1/collections/mi-coleccion/keys?prefix=config:&limit=20" \
  -H "apikey: tu-api-key" \
  -H "secret: tu-api-secret"
```

Tanto `prefix` como `limit` son parámetros de consulta opcionales:
- `prefix` — filtra las claves que comienzan con la cadena indicada
- `limit` — limita la cantidad de registros devueltos (por defecto: sin límite)

Respuesta:

```json
[
  {"key":"config:tema","value":{"modo":"oscuro","tamanoFuente":14},"version":1,"updatedAt":"2025-01-01T00:00:00Z"},
  {"key":"config:idioma","value":"es-ES","version":1,"updatedAt":"2025-01-01T00:00:01Z"}
]
```

## Convenciones para Nombres de Claves

Aunque KVNode acepta cualquier cadena UTF-8 como clave, recomendamos seguir estas convenciones:

- **Usa espacios de nombres con dos puntos**: `app:modulo:clave` (ej. `config:basedatos:host`)
- **Evita caracteres especiales** que puedan interferir con la codificación URL (`/`, `?`, `#`)
- **Mantén las claves legibles** — los nombres significativos simplifican la depuración
- **Sé consistente** — elige un patrón de nomenclatura y aplícalo en todas las colecciones

## Tipos de Valor

Todos los valores deben ser JSON válido. KVNode soporta:

- **Objetos**: `{"usuario": "alice", "rol": "admin"}`
- **Arreglos**: `["item1", "item2"]`
- **Cadenas**: `"hola mundo"`
- **Números**: `42` o `3.14`
- **Booleanos**: `true` o `false`
- **Null**: `null`

El tamaño máximo del valor está determinado por la configuración del backend de WAL.
