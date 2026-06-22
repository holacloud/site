# Documentación

Bienvenido a la documentación de HolaCloud. Aquí encontrarás guías y referencias de API para los servicios de la plataforma.

## Primeros pasos

Si eres nuevo en HolaCloud, empieza por la [Consola](/es/docs/console/) y las guías de cada servicio. Algunos servicios siguen evolucionando, por lo que conviene leer la referencia API junto con las notas respaldadas por la implementación actual.

Cada servicio incluye su propia guía de inicio, documentación conceptual y una referencia API con detalles de endpoints, ejemplos de solicitud y códigos de error.

## Servicios

### [InceptionDB](/es/docs/inceptiondb/)
Base de datos NoSQL documental. Almacena y consulta documentos JSON con esquema dinámico, colecciones e índices.

### [Lambda](/es/docs/lambda/)
Plataforma de ejecución de funciones. Despliega funciones JavaScript y activos estáticos detrás de endpoints HTTP.

### [Files](/es/docs/files/)
API de archivos. Sube, descarga, lista, inspecciona y elimina archivos organizados en buckets.

### [Config](/es/docs/config/)
API de configuración. Almacena y actualiza entradas de configuración JSON por usuario.

### [InstantLogs](/es/docs/logs/)
Ingesta y streaming de logs en tiempo real. Recoge logs raw o framed, filtra con expresiones regulares y transmite eventos.

### [Tailon](/es/docs/queues/)
API de colas de mensajes. Crea colas, escribe mensajes JSON y lee JSON delimitado por saltos de línea desde colas con nombre.

### [Scheduler](/es/docs/scheduler/)
Cola de tareas diferidas. Encola tareas diferidas puntuales, resérvalas con leases y transmite snapshots del scheduler.

### [KVNode](/es/docs/kvnode/)
API clave-valor con colecciones, claves, estado de nodo, métricas y sincronización root/replica.

### [Consola](/es/docs/console/)
Superficie web de consola para servicios HolaCloud, incluidas las APIs fake actuales usadas por la UI.

### [Glue2](/es/docs/glue2/)
Gateway por host y capa de unión de servicios con virtual hosts, inyección de contexto de autenticación, estado y estadísticas de tráfico.

### [Holamail](/es/docs/holamail/)
Servidor SMTP básico de pruebas. Acepta y registra mensajes SMTP para pruebas y desarrollo.

### [Run](/es/docs/run/)
Controles de runtime de contenedores con un subconjunto Docker Registry orientado a push. Inicia, detiene y revierte runtimes de repositorio.

## Referencia de API

Cada servicio incluye una sección de referencia API con:

- Descripción de endpoints y métodos HTTP
- Requisitos de autenticación
- Ejemplos de solicitud y respuesta en JSON
- Tablas de códigos de error
- Ejemplos `curl` para cada endpoint

Navega por la barra lateral para acceder a cualquier servicio o usa los enlaces de productos en la navegación superior.
