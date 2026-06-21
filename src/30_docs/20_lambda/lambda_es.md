<h1>Lambda</h1>

HolaCloud Lambda es un servicio de cómputo serverless que te permite ejecutar código en respuesta a eventos sin necesidad de aprovisionar ni administrar servidores. Sube tus funciones en JavaScript, Go o Python, y HolaCloud se encarga de todo, desde el escalado automático hasta la alta disponibilidad.

## Características Principales

### Soporte Multi-Lenguaje
Escribe funciones en el lenguaje que prefieras. HolaCloud Lambda es compatible con JavaScript (Node.js), Go y Python, ofreciendo una interfaz de invocación consistente en todos los lenguajes.

### Escalado Automático
Las funciones Lambda se escalan automáticamente según el volumen de solicitudes entrantes. Ya sea que tengas una solicitud al día o miles por segundo, HolaCloud ajusta la capacidad al instante — solo pagas por lo que usas.

### Soporte para Webhooks
Cada función Lambda es automáticamente accesible mediante una URL pública. Úsala como endpoint de webhook para servicios externos como Stripe, GitHub o Slack, o intégrala con tus propias aplicaciones.

### Facturación por Uso
Sin costos fijos ni cargos por inactividad. Solo pagas por el tiempo de cómputo que tu código consume durante la ejecución, redondeado al milisegundo más cercano.

## Casos de Uso

- **APIs**: Crea backends REST o GraphQL ligeros sin administrar infraestructura.
- **Webhooks**: Recibe y procesa eventos de servicios externos con un endpoint público dedicado.
- **Tareas en Segundo Plano**: Externaliza procesos como envío de correos, procesamiento de imágenes o generación de reportes.
- **Procesamiento de Datos**: Transforma, filtra o enriquece flujos de datos en tiempo real.

## Integración en el Ecosistema HolaCloud

Lambda funciona junto a otros servicios de HolaCloud para potenciar aplicaciones serverless completas:

- **InceptionDB** — almacena y consulta datos persistentes desde tus funciones.
- **HolaAuth** — autentica usuarios y protege endpoints.
- **Mux Router** — asigna dominios personalizados a funciones y propietarios específicos.

Todos los servicios comparten un gateway de API común, un modelo de autenticación unificado y las mismas herramientas CLI, lo que facilita la creación de aplicaciones integradas.
