# KVNode

KVNode es un almacén clave-valor distribuido y replicado, diseñado para alta disponibilidad y acceso de baja latencia. Forma parte del ecosistema HolaCloud y expone una API HTTP simple para almacenar y recuperar valores JSON.

## Características Principales

### Persistencia Basada en WAL
Cada escritura se registra en un Write-Ahead Log (WAL) antes de aplicarse al almacenamiento en memoria. Esto garantiza durabilidad y permite la replicación sin interrupciones entre nodos.

### Backends de WAL Intercambiables
KVNode soporta múltiples backends de WAL — memoria, Kafka, PostgreSQL, Redis y MongoDB — para que elijas la capa de persistencia que mejor se adapte a tu infraestructura.

### SDKs Multi-Lenguaje
Hay SDKs oficiales disponibles para Go, Python, Java, JavaScript, Kotlin, PHP y Node.js, facilitando la integración de KVNode en cualquier pila tecnológica.

### Consistencia Fuerte
Las escrituras se replican de forma síncrona o asíncrona según el backend de WAL, con semántica linealizable para operaciones sobre una sola clave.

## Casos de Uso

- **Almacenamiento de Configuración**: Guarda la configuración de tu aplicación como pares clave-valor con replicación automática entre nodos.
- **Feature Flags**: Administra interruptores de funcionalidad de forma centralizada y propaga los cambios en tiempo real.
- **Descubrimiento de Servicios**: Registra y descubre endpoints de servicio con lecturas de baja latencia.
- **Almacén de Sesiones**: Almacena sesiones de usuario con búsquedas rápidas y replicación integrada para tolerancia a fallos.

## Resumen de la API

KVNode expone una API RESTful en `https://api.hola.cloud`. Todos los endpoints que modifican datos requieren autenticación interna mediante clave y secreto de API, o mediante el encabezado de autenticación del gateway Glue.

| Método | Ruta | Descripción | Autenticación |
|--------|------|-------------|---------------|
| GET | /healthz | Verificación de salud | Pública |
| GET | /readyz | Disponibilidad (verifica conexión con el padre) | Pública |
| GET | /v1/status | Estado del nodo (colecciones, replicación, uptime) | Interna |
| GET | /v1/metrics | Métricas del nodo (escrituras, lecturas) | Interna |
| GET | /v1/collections | Listar colecciones | Interna |
| POST | /v1/collections | Crear colección | Interna |
| DELETE | /v1/collections/{col} | Eliminar colección | Interna |
| GET | /v1/collections/{col}/keys | Listar claves (límite, prefijo) | Interna |
| GET | /v1/collections/{col}/keys/* | Obtener valor de clave | Interna |
| POST | /v1/collections/{col}/keys/* | Establecer valor de clave | Interna |
| DELETE | /v1/collections/{col}/keys/* | Eliminar clave | Interna |
| POST | /v1/replicate | Flujo de replicación (NDJSON) | Interna |
