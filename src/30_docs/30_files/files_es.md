# Files

HolaCloud Files es un servicio de almacenamiento de objetos compatible con S3, diseñado para almacenar y recuperar cualquier cantidad de datos en cualquier momento. Proporciona una API RESTful simple para gestionar buckets y archivos, con escalabilidad, durabilidad y seguridad integradas.

## Características Clave

### API Compatible con S3
HolaCloud Files expone una API familiar compatible con S3, lo que facilita la migración de aplicaciones existentes o la integración con herramientas que utilizan el protocolo S3.

### Escalabilidad
Files se escala automáticamente para manejar cualquier cantidad de datos, desde unos pocos gigabytes hasta petabytes, sin necesidad de aprovisionamiento manual ni planificación de capacidad.

### Durabilidad y Disponibilidad
Los datos se almacenan de forma redundante en múltiples zonas de disponibilidad, garantizando un 99.999999999% de durabilidad y un 99.99% de disponibilidad.

### Seguridad
Todos los datos se cifran en reposo mediante AES-256 y en tránsito mediante TLS 1.3. El acceso se controla mediante claves API con permisos detallados.

### Preparado para Versionado
El servicio está diseñado con soporte para versionado, permitiéndole conservar, recuperar y restaurar cada versión de cada archivo en sus buckets.

## Casos de Uso

### Copias de Seguridad y Recuperación ante Desastres
Files proporciona un destino confiable para copias de seguridad de bases de datos, instantáneas del sistema y archivos de recuperación ante desastres. Sus garantías de durabilidad aseguran que sus datos sobrevivan a fallos de hardware y cortes regionales.

### Activos Estáticos para Aplicaciones Web
Aloje imágenes, CSS, JavaScript y otros activos estáticos con baja latencia y alto rendimiento. Files se integra perfectamente con HolaCloud CDN para la entrega global de contenido.

### Contenido Generado por el Usuario
Maneje las cargas de archivos de sus usuarios -- fotos de perfil, documentos, videos y más -- con una capa de almacenamiento escalable y segura.

### Data Lakes y Analítica
Almacene datos estructurados y no estructurados a escala de petabytes para cargas de trabajo analíticas. Files se integra con HolaCloud Compute y HolaCloud Analytics para pipelines de procesamiento de datos.

### Almacenamiento y Streaming de Medios
Almacene archivos de audio, video e imágenes para aplicaciones multimedia. Files soporta solicitudes de rango para la entrega eficiente de contenido parcial.

## Integración con HolaCloud

HolaCloud Files funciona perfectamente con otros servicios de HolaCloud:

- **HolaCloud Compute** -- Adjunte volúmenes de Files a sus instancias de cómputo para almacenamiento persistente.
- **HolaCloud CDN** -- Entregue contenido de archivos globalmente mediante caché perimetral.
- **HolaCloud Logs** -- Archive registros de acceso directamente en buckets de Files.
- **HolaCloud InceptionDB** -- Almacene instantáneas de copias de seguridad de bases de datos en Files.

## Referencia de API

Todos los endpoints utilizan la URL base `https://api.hola.cloud` y requieren autenticación mediante los encabezados `Api-Key` y `Api-Secret`.

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | /v1/buckets | Listar todos los buckets |
| POST | /v1/buckets | Crear un nuevo bucket |
| GET | /v1/buckets/{id} | Obtener detalles de un bucket |
| PATCH | /v1/buckets/{id} | Modificar metadatos de un bucket |
| DELETE | /v1/buckets/{id} | Eliminar un bucket |
| GET | /v1/buckets/{id}/list/* | Listar archivos en un bucket |
| PUT | /v1/buckets/{id}/files/* | Subir un archivo |
| GET | /v1/buckets/{id}/files/* | Descargar un archivo |
| DELETE | /v1/buckets/{id}/files/* | Eliminar un archivo |
| HEAD | /v1/buckets/{id}/files/* | Obtener metadatos de un archivo |
