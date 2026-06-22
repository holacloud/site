# Glue2

Glue2 es la puerta de enlace central de API para HolaCloud. Autentica cada solicitud entrante, la enruta al servicio backend correcto y recopila registros de acceso. Todo el tráfico hacia los servicios de HolaCloud pasa a través de Glue2.

## Autenticación

Glue2 soporta tres mecanismos de autenticación:

### Sesiones
Sesiones basadas en navegador autenticadas mediante HolaAuth. Tras el inicio de sesión, la cookie de sesión se utiliza para autenticar solicitudes posteriores.

### Claves API
Acceso máquina a máquina mediante un par de clave y secreto de API. Se envían como los encabezados HTTP `apikey` y `secret` y se validan contra las credenciales del proyecto.

### OAuth
La integración con OAuth 2.0 permite la autenticación a través de proveedores externos. Glue2 valida los tokens e inyecta información de identidad del usuario en las solicitudes proxy.

## Enrutamiento por Host Virtual

Glue2 enruta las solicitudes según el encabezado `Host`. Cada proyecto en HolaCloud obtiene un subdominio único (`project.hola.cloud`), y Glue2 mapea las solicitudes entrantes al servicio backend correspondiente para ese proyecto.

```
┌─────────────┐     ┌──────────┐     ┌──────────────┐
│  Cliente    │────▶│  Glue2   │────▶│  Svc Proyecto │
└─────────────┘     └──────────┘     └──────────────┘
                           │
                           ├────▶ InceptionDB
                           ├────▶ Lambda
                           ├────▶ KVNode
                           └────▶ Files / Config / etc.
```

## Encabezados Inyectados

Cuando Glue2 reenvía una solicitud a un servicio backend, inyecta los siguientes encabezados:

| Encabezado | Descripción |
|------------|-------------|
| `X-Glue-Authentication` | Contexto de autenticación codificado en Base64 (ID de usuario, sesión, roles) |
| `X-Holacloud-Project-Id` | El ID del proyecto extraído del host virtual |
| `X-Forwarded-For` | La dirección IP original del cliente |

Los servicios backend utilizan estos encabezados para autorización y aislamiento de inquilinos.

## Límite de Velocidad

Glue2 aplica límites de velocidad por proyecto y por endpoint. Los límites son configurables en la configuración del proyecto y pueden basarse en solicitudes por segundo, conexiones simultáneas o cuotas diarias. Cuando se excede un límite, Glue2 devuelve una respuesta `429 Too Many Requests`.

## Registros de Acceso

Todas las solicitudes que pasan por Glue2 se registran con:

- Marca de tiempo y duración de la solicitud
- IP del cliente y agente de usuario
- Método, ruta y código de estado de la solicitud
- ID del proyecto y método de autenticación
- Servicio backend destino y tiempo de respuesta

Los registros se transmiten a InstantLogs para monitoreo y análisis en tiempo real.
