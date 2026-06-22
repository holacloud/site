# Documentación

Bienvenido a la documentación de HolaCloud. Aquí encontrarás guías, referencias de API y manuales operativos para todos los servicios de la plataforma.

## Primeros pasos

Si eres nuevo en HolaCloud, el mejor lugar para empezar es la [Consola](/es/docs/console/) — la interfaz de gestión web donde puedes aprovisionar recursos, administrar proyectos y monitorizar tus servicios.

Cada servicio incluye su propia guía de inicio, documentación conceptual y una referencia completa de API con detalles de cada endpoint, ejemplos de solicitud y códigos de error.

## Servicios

### [InceptionDB](/es/docs/inceptiondb/)
Base de datos NoSQL gestionada. Almacena y consulta documentos JSON con esquema dinámico, indexación y ejecución de JavaScript. Ideal para aplicaciones que necesitan modelos de datos flexibles.

### [Lambda](/es/docs/lambda/)
Plataforma de ejecución serverless. Despliega código que se ejecuta bajo demanda, escala automáticamente y se integra con el ecosistema HolaCloud. Sin servidores que gestionar.

### [Files](/es/docs/files/)
Almacenamiento de objetos seguro. Sube, descarga y organiza archivos con políticas de bucket configurables. Adecuado para almacenamiento multimedia, copias de seguridad y distribución de contenido.

### [Config](/es/docs/config/)
Gestión centralizada de configuración. Almacena, versiona y distribuye la configuración de tus aplicaciones en todos los entornos. Soporta parches JSON y sobrescrituras por usuario.

### [InstantLogs](/es/docs/logs/)
Ingesta y streaming de logs en tiempo real. Recoge datos estructurados y no estructurados de cualquier fuente, filtra en tiempo real e intégralos con tus pipelines de monitorización.

### [Tailon](/es/docs/queues/)
Sistema de colas de mensajes y workloads asíncronos. Entrega fiable de mensajes con consumidores long-poll, seguimiento de clientes y soporte para procesamiento asíncrono de alto rendimiento.

### [Scheduler](/es/docs/scheduler/)
Servicio distribuido de programación de tareas. Programa tareas únicas y recurrentes, gestiona leases de tareas y transmite eventos. Construido para workloads temporizados y en segundo plano.

### [KVNode](/es/docs/kvnode/)
Almacén clave-valor replicado con backends de persistencia configurables. Lecturas y escrituras de baja latencia, comprobaciones de salud y métricas por nodo.

### [Consola](/es/docs/console/)
Consola de gestión web para todos los servicios de HolaCloud. Provisiona proyectos, gestiona recursos y monitoriza el estado de los servicios desde una única interfaz.

### [Glue2](/es/docs/glue2/)
API Gateway y capa de autenticación. Enrutamiento central, autenticación y gestión de tráfico para la malla de servicios de HolaCloud. Soporta autenticación mediante API key y bearer token.

### [Holamail](/es/docs/holamail/)
Servicio SMTP de correo transaccional. Envía correos mediante relay SMTP con seguimiento de entrega fiable. Adecuado para notificaciones, restablecimientos de contraseña y mensajería transaccional.

### [Run](/es/docs/run/)
Servicio de ejecución de contenedores con registro Docker integrado. Inicia, detiene y revierte contenedores. Gestiona variables de entorno, volúmenes y ciclos de vida de despliegues.

## Referencia de API

Cada servicio incluye una sección completa de referencia de API con:

- Descripción de endpoints y métodos HTTP
- Requisitos de autenticación
- Ejemplos de solicitud y respuesta en JSON
- Tablas de códigos de error
- Ejemplos `curl` para cada endpoint

Navega por la barra lateral para acceder a cualquier servicio o usa los enlaces de productos en la navegación superior.
