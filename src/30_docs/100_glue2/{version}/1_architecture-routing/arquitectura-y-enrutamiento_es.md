<h1>Arquitectura y Enrutamiento</h1>

Glue2 es el proxy inverso central que se sitúa frente a todos los servicios de HolaCloud. Todo el tráfico entrante — ya sea de navegadores, clientes API o servicios internos — pasa a través de Glue2 antes de llegar a su destino.

![Arquitectura de Glue2](img/glue2-architecture.png)

## Proxy Inverso

Glue2 opera como un proxy inverso de capa 7. Termina TLS, inspecciona cada solicitud, autentica al cliente y luego reenvía la solicitud al servicio backend correspondiente. Ningún tráfico externo llega directamente a los servicios de HolaCloud; Glue2 es el único punto de entrada.

```
Cliente ──▶ Glue2 (verificación de auth) ──▶ Servicio Backend
               │
               ├──▶ InceptionDB
               ├──▶ Lambda
               ├──▶ Files
               ├──▶ Config
               ├──▶ KVNode
               ├──▶ Scheduler
               └──▶ Logs
```

## Enrutamiento por Host Virtual

Glue2 enruta las solicitudes según el encabezado `Host`. Cada proyecto de HolaCloud recibe un subdominio único (`<proyecto>.hola.cloud`), y Glue2 mapea ese subdominio al servicio backend del proyecto.

Los servicios de plataforma integrados usan subdominios reservados:

| Host | Servicio |
|------|----------|
| `inceptiondb.hola.cloud` | InceptionDB |
| `api.hola.cloud` | API pública (Lambda, Files, Config, etc.) |
| `console.hola.cloud` | Consola de HolaCloud |
| `auth.hola.cloud` | Autenticación / Servicio de sesión |
| `logs.hola.cloud` | InstantLogs |

Cuando llega una solicitud para `mi-proyecto.hola.cloud`, Glue2 busca el proyecto con ese slug, resuelve su servicio backend principal y reenvía la solicitud.

## Flujo del Proxy

1. El cliente envía una solicitud HTTP a `<proyecto>.hola.cloud`.
2. Glue2 termina la conexión TLS y lee el encabezado `Host`.
3. Glue2 verifica la autenticación (cookie de sesión, clave API o token OAuth). Si el endpoint requiere autenticación y no se proporciona, la solicitud es rechazada con `401 Unauthorized`.
4. Glue2 resuelve el host virtual a un ID de proyecto y un servicio backend destino.
5. Glue2 inyecta los encabezados de autenticación y proyecto en la solicitud.
6. La solicitud se reenvía al servicio backend.
7. La respuesta del backend se transmite de vuelta al cliente y Glue2 registra la transacción.

## Encabezados Inyectados

Al reenviar a un backend, Glue2 inyecta:

| Encabezado | Descripción |
|------------|-------------|
| `X-Glue-Authentication` | JWT codificado en Base64 con ID de usuario, sesión, roles y método de auth |
| `X-Holacloud-Project-Id` | El UUID del proyecto derivado del host virtual |
| `X-Holacloud-Tenant-Project-Id` | ID de proyecto con alcance de inquilino para aislamiento multiinquilino |
| `X-Forwarded-For` | Dirección IP original del cliente |
| `X-Forwarded-Proto` | Protocolo original (`http` o `https`) |

Los servicios backend usan estos encabezados para identificar al cliente y aplicar autorización. El JWT de `X-Glue-Authentication` está firmado con un secreto compartido conocido solo por Glue2 y los backends de confianza.

## Endpoints de Administración V0

Glue2 expone endpoints de administración bajo `/v0/`:

### Listar Hosts Virtuales

```bash
curl "https://api.hola.cloud/v0/virtualhosts" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

```json
{
  "virtual_hosts": [
    {
      "host": "mi-proyecto.hola.cloud",
      "project_id": "p3b2c1a0-1234-5678-9abc-def012345678",
      "backend": "svc-1"
    }
  ]
}
```

### Estadísticas de Tráfico

```bash
curl "https://api.hola.cloud/v0/stats" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

### Estado de Salud del Backend

```bash
curl "https://api.hola.cloud/v0/status" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

```json
{
  "backends": [
    {
      "name": "inceptiondb",
      "status": "saludable",
      "latency_ms": 12
    }
  ]
}
```

## Balanceo de Carga y Failover

Glue2 distribuye el tráfico entre múltiples instancias backend para cada servicio. Si una instancia no está saludable (falla las comprobaciones de salud o devuelve errores), Glue2 la elimina de la rotación y reintenta en una instancia saludable. Las comprobaciones de salud se ejecutan cada 10 segundos. Puedes monitorear la salud del backend a través del endpoint `/v0/status`.

## Siguientes Pasos

- Aprende sobre los mecanismos de autenticación en [Autenticación](../2_authentication/autenticacion_es.md).
- Aprende a gestionar claves API en [Gestión de Claves API](../3_api-keys/gestionando-claves-api_es.md).
