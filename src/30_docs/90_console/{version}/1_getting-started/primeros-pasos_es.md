
# Primeros pasos con la Consola

La Consola es el plano de control único para todos los servicios de HolaCloud. Proporciona una interfaz web unificada para administrar bases de datos, funciones, archivos, configuración, registros, colas y más — todo desde un panel central.

![Panel de la Consola](img/console.png)

## Iniciar sesión

Navegue a [https://console.hola.cloud](https://console.hola.cloud) e inicie sesión con sus credenciales de HolaCloud. Si está usando el modo de demostración o desarrollo, los endpoints mock de la API estarán activos, permitiéndole explorar todas las funcionalidades sin un backend real.

```bash
# Ejemplo: verificar sesión mediante la API mock
curl -X GET "https://console.hola.cloud/fakeapi/authapi/auth/me"
```

Respuesta esperada:

```json
{
  "email": "fulanez@gmail.com",
  "id": "user-1234",
  "nick": "fulanez"
}
```

## Navegar por el panel principal

Después de iniciar sesión, el panel principal muestra:

- **Estado de los servicios** — indicadores de salud para cada servicio activo de HolaCloud (InceptionDB, Lambda, Files, Config, InstantLogs, Queues, Scheduler).
- **Actividad reciente** — un feed cronológico de las acciones realizadas en su proyecto.
- **Métricas clave** — resúmenes de uso como conteo de solicitudes, tamaño de almacenamiento e invocaciones de funciones.

Use la barra lateral izquierda para saltar entre secciones: Panel, Bases de Datos, Funciones, Archivos, Configuración, Registros, Colas, Programador y Configuración del Proyecto.

## Ver estado del servicio y actividad reciente

Cada tarjeta de servicio en el panel muestra su estado actual (operativo, degradado o fuera de servicio). Haga clic en una tarjeta para acceder a la vista detallada de ese servicio. El feed de actividad reciente se puede filtrar por tipo de servicio y rango de tiempo.

```bash
# Obtener actividad del proyecto mediante la API mock
curl -X GET "https://console.hola.cloud/fakeapi/projectsapi/v0/projects"
```

```json
[
  { "id": "project-00000000-0000-0000-0000-000000000001", "name": "Hello" }
]
```

## Cambiar entre proyectos

Si pertenece a varios proyectos, use el selector de proyectos en la parte superior de la barra lateral. Toda la interfaz cambia al contexto del proyecto seleccionado y sus servicios.

```bash
# Listar proyectos disponibles
curl -X GET "https://console.hola.cloud/fakeapi/projectsapi/v0/projects"
```

```json
[
  { "id": "project-00000000-0000-0000-0000-000000000001", "name": "Hello" }
]
```

## Acceder a documentación y soporte

La Consola incluye un menú de ayuda integrado (icono de signo de interrogación en la barra superior) con enlaces a:

- Este sitio de documentación
- Referencia de la API
- Foro de la comunidad
- Sistema de tickets de soporte
- Página de estado del servicio

## Resumen

La Consola actúa como el plano de control único para todos los servicios de HolaCloud. Ya sea que esté trabajando en producción o explorando en modo de demostración, la Consola le brinda visibilidad y control total sobre sus recursos en la nube desde un solo lugar.
