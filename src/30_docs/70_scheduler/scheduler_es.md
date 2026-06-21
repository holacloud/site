# Scheduler

El Scheduler es un servicio distribuido de programación de tareas que forma parte del ecosistema HolaCloud. Proporciona encolado de tareas basado en tiempo, reserva por concesión para la seguridad de los trabajadores y monitoreo en tiempo real a través de Server-Sent Events (SSE). Diseñado como una alternativa ligera a los sistemas cron tradicionales, el Scheduler se integra perfectamente con el resto de tus servicios de HolaCloud.

## Características Clave

### Almacenamiento de Tareas en Memoria
Las tareas se mantienen en memoria para una programación y despacho rápidos. El servicio está optimizado para cargas de trabajo de alto rendimiento y baja latencia donde las tareas se procesan poco después de su hora programada.

### Programación Basada en Tiempo
Cada tarea tiene un campo `future` (o `delay`) que controla cuándo estará disponible. Las tareas permanecen inactivas hasta que llega su hora programada, momento en el cual se vuelven visibles para los trabajadores.

### Reserva por Concesión
Cuando un trabajador reserva una tarea, la tarea entra en un período de concesión definido por `worktime`. Si el trabajador completa y confirma la tarea dentro de la concesión, se elimina. Si la concesión expira, la tarea vuelve a estar disponible para que otro trabajador la reserve. Esto garantiza procesamiento al menos una vez sin intervención manual.

### Streaming SSE
El endpoint `/schedulers/{id}/tasks/stream` ofrece un flujo SSE en tiempo real de eventos de tareas. Los trabajadores y paneles de monitoreo pueden suscribirse para recibir nuevas tareas, finalizaciones y expiraciones de concesión a medida que ocurren.

### Etiquetado y Filtrado
Las tareas pueden anotarse con `labels` arbitrarios (pares clave-valor) al encolarse. Las operaciones de listado admiten filtrado por etiquetas, lo que facilita la inspección de tareas por proyecto, prioridad o grupo de trabajadores.

## Casos de Uso

### Reemplazo de Cron
Programa trabajos recurrentes o únicos sin gestionar archivos crontab. Cada llamada programada crea una tarea que estará disponible en el momento especificado.

### Procesamiento de Trabajos Retrasados
Encola trabajos con un retraso — por ejemplo, enviar un correo de recordatorio 24 horas después de que un usuario se registre. El Scheduler retiene la tarea hasta que llegue su momento.

### Trabajadores Distribuidos
Múltiples trabajadores pueden reservar y procesar tareas del mismo scheduler. El mecanismo de concesión asegura que una tarea se procese exactamente una vez en condiciones normales, y se reintente automáticamente si un trabajador falla.

### Colas de Reintento
Cuando una tarea falla, se puede reprogramar con un retraso mayor. El endpoint de `reschedule` del Scheduler facilita la implementación de retroceso exponencial.

El Scheduler de HolaCloud es una alternativa ligera y fácil de usar a los programadores de trabajos tradicionales, construida para aplicaciones nativas en la nube.
