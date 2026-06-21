# Tailon Colas

Tailon es un servicio ligero de colas de mensajes en memoria que forma parte del ecosistema HolaCloud. Diseñado para la velocidad y la simplicidad, Tailon permite la comunicación asíncrona entre microservicios, la distribución de eventos y el almacenamiento en búfer de cargas de trabajo sin la sobrecarga de un gestor de mensajes completo.

## Características Clave

### Rendimiento en Memoria
Los mensajes se almacenan en memoria para operaciones de escritura y lectura ultrarrápidas. Tailon está optimizado para escenarios de alto rendimiento y baja latencia donde la persistencia no es la prioridad principal.

### Consumo por Long-Poll
Los consumidores recuperan mensajes mediante solicitudes HTTP de long-poll. Esto reduce la necesidad de sondeos constantes mientras mantiene las conexiones ligeras. El encabezado opcional `Limit` controla cuántos mensajes se devuelven por solicitud.

### Escritura en Streaming
Los productores pueden enviar múltiples mensajes en una sola solicitud utilizando JSON delimitado por saltos de línea (NDJSON), lo que hace que la ingesta por lotes sea eficiente y sencilla.

### Multi-Productor / Multi-Consumidor
Cualquier número de productores puede escribir en la misma cola, y cualquier número de consumidores puede leer de ella. Tailon maneja el acceso concurrente de forma segura sin configuración adicional.

### Monitoreo de Clientes Activos
Tailon expone el endpoint `/v1/clients` para que puedas inspeccionar qué consumidores de streaming están actualmente conectados a tus colas.

## Casos de Uso

### Colas de Tareas Asíncronas
Descarga operaciones pesadas o lentas — como envío de correos, procesamiento de imágenes o generación de informes — a trabajadores en segundo plano que consumen de una cola de Tailon.

### Distribución de Eventos
Transmite eventos desde un solo productor a múltiples servicios consumidores. Cada consumidor lee de forma independiente, permitiendo patrones de fan-out.

### Nivelación de Carga
Almacena en búfer las solicitudes entrantes durante picos de tráfico. Los trabajadores consumen a su propio ritmo, evitando que los servicios descendentes se vean abrumados.

### Desacoplamiento de Microservicios
Reemplaza las llamadas HTTP directas entre servicios con comunicación basada en colas. Los productores y consumidores evolucionan de forma independiente mientras el contrato del mensaje sea estable.

Tailon de HolaCloud es una solución de colas simple pero potente para aplicaciones modernas que necesitan un paso de mensajes rápido, fiable y escalable.
