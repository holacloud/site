<h1>Lambda</h1>

HolaCloud Lambda permite publicar pequeños manejadores HTTP que se ejecutan dentro de HolaCloud. Una lambda almacena su código fuente, metadatos de ruta, propietario y asociación de proyecto, y responde mediante la API de Lambda o las rutas públicas de ejecución.

## Funciones Principales

### Lenguajes Soportados
Lambda acepta estos valores de `language`: `javascript`, `static-html`, `static-css` y `static-js`.

### Enrutamiento HTTP
Cada lambda puede almacenar `method` y `path`. Usa estos campos para describir cómo debe alcanzarse la lambda mediante el mux router o la capa de rutas de tu aplicación.

### Invocación Pública y Administrativa
Usa endpoints administrativos autenticados para crear, inspeccionar, actualizar, eliminar y ejecutar lambdas. Usa las rutas públicas de ejecución y mux cuando una lambda deba ser llamada por webhooks, navegadores u otros clientes externos.

## Casos de Uso

- **APIs HTTP**: Crea pequeños manejadores de solicitudes apoyados en servicios de HolaCloud.
- **Webhooks**: Recibe eventos de terceros mediante una URL pública `/run/{lambda_id}`.
- **Respuestas Estáticas**: Sirve fragmentos HTML, CSS o JavaScript de cliente con los modos estáticos.
- **Rutas Mux**: Enruta paths por propietario mediante `/mux/{owner_id}/*`.

## Campos de Lambda

Los registros Lambda usan estos campos: `id`, `created_timestamp`, `owner`, `project_id`, `name`, `language`, `code`, `method` y `path`.
