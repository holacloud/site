# Primeros pasos con Files

Esta guía le muestra las operaciones básicas de HolaCloud Files: crear un bucket, subir un archivo, listar contenidos, descargar y eliminar.

## Prerrequisitos

Antes de comenzar, necesita:

- Una cuenta de HolaCloud con una clave API y un secreto API
- `curl` instalado en su máquina

Sus credenciales API se pasan como encabezados en cada solicitud:

```
Api-Key: su-api-key
Api-Secret: su-api-secret
```

Todas las solicitudes utilizan la URL base `https://api.hola.cloud`.

## Paso 1: Crear un Bucket

Los buckets son contenedores para sus archivos. Cada bucket debe tener un nombre único a nivel global.

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-primer-bucket"
  }'
```

Respuesta esperada:

```json
{
  "id": "bkt_abc123",
  "name": "mi-primer-bucket",
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## Paso 2: Subir un Archivo

Sube un archivo de texto al bucket usando PUT. La ruta del archivo se especifica en la URL después de `/files/`.

```bash
echo "¡Hola, HolaCloud Files!" > hola.txt

curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/hola.txt" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: text/plain" \
  --data-binary @hola.txt
```

Respuesta esperada:

```json
{
  "path": "hola.txt",
  "size": 24,
  "contentType": "text/plain",
  "uploadedAt": "2026-06-21T10:01:00Z",
  "etag": "\"d41d8cd98f00b204e9800998ecf8427e\""
}
```

## Paso 3: Listar Contenidos del Bucket

Lista todos los archivos en el bucket con un filtro de prefijo opcional:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

Respuesta esperada:

```json
{
  "files": [
    {
      "path": "hola.txt",
      "size": 24,
      "contentType": "text/plain",
      "modifiedAt": "2026-06-21T10:01:00Z"
    }
  ],
  "prefix": "",
  "total": 1
}
```

## Paso 4: Descargar el Archivo

Descarga el archivo usando GET:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/hola.txt" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -o hola_descargado.txt

cat hola_descargado.txt
```

Salida:

```
¡Hola, HolaCloud Files!
```

## Paso 5: Eliminar el Archivo

Elimina el archivo usando DELETE:

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123/files/hola.txt" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

Respuesta esperada: HTTP 204 Sin Contenido.

## Paso 6: Eliminar el Bucket

El bucket debe estar vacío antes de poder eliminarlo. Como eliminamos el archivo, esto tendrá éxito:

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

Respuesta esperada: HTTP 204 Sin Contenido.

## Resumen

Ha creado exitosamente un bucket, subido un archivo, listado los contenidos del bucket, descargado el archivo y limpiado los recursos. Ahora está listo para integrar HolaCloud Files en sus aplicaciones.
