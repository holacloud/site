<h1>Gestión de Claves API</h1>

Las claves API son el método principal de autenticación máquina a máquina para los servicios de HolaCloud. Cada clave consta de un identificador público basado en UUID (Api-Key) y un secreto basado en UUID (Api-Secret). El secreto se hashea con bcrypt y no se puede recuperar después de la creación.

## Crear una Clave API

### A Través de la Consola

1. Inicia sesión en [https://console.hola.cloud](https://console.hola.cloud).
2. Navega a **Configuración > Claves API**.
3. Haz clic en **Crear Clave API**.
4. Opcionalmente, establece los ámbitos de la clave (proyecto, host, ruta, método).
5. Haz clic en **Crear**.
6. Copia el Api-Key y el Api-Secret inmediatamente — el secreto se muestra solo una vez.

### A Través de la API de Serviceprojects

Las claves API pertenecen a proyectos. Primero, crea un proyecto (si no tienes uno):

```bash
curl -X POST "https://api.hola.cloud/api/v0/projects" \
  -H "Api-Key: TU_API_KEY_ADMIN" \
  -H "Api-Secret: TU_API_SECRET_ADMIN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mi Proyecto",
    "slug": "mi-proyecto"
  }'
```

Respuesta esperada:

```json
{
  "id": "p3b2c1a0-1234-5678-9abc-def012345678",
  "name": "Mi Proyecto",
  "slug": "mi-proyecto",
  "created_at": "2025-06-21T12:00:00Z"
}
```

Luego genera una clave API para el proyecto:

```bash
curl -X POST "https://api.hola.cloud/api/v0/projects/p3b2c1a0-1234-5678-9abc-def012345678/apikeys" \
  -H "Api-Key: TU_API_KEY_ADMIN" \
  -H "Api-Secret: TU_API_SECRET_ADMIN" \
  -H "Content-Type: application/json" \
  -d '{
    "scopes": {
      "hosts": ["mi-proyecto.hola.cloud"],
      "paths": ["/api/v0/lambdas/*"],
      "methods": ["GET", "POST"]
    }
  }'
```

Respuesta esperada:

```json
{
  "api_key": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "api_secret": "f0e1d2c3-b4a5-6789-0fed-cba987654321",
  "project_id": "p3b2c1a0-1234-5678-9abc-def012345678",
  "scopes": {
    "hosts": ["mi-proyecto.hola.cloud"],
    "paths": ["/api/v0/lambdas/*"],
    "methods": ["GET", "POST"]
  },
  "created_at": "2025-06-21T12:00:00Z"
}
```

## Estructura de la Clave

- **Api-Key**: UUID v4, actúa como el identificador público de la clave. Se envía en el encabezado `Api-Key`.
- **Api-Secret**: UUID v4, actúa como el secreto. Se envía en el encabezado `Api-Secret`. Se almacena como un hash bcrypt — HolaCloud no puede recuperarlo si se pierde.

## Ámbitos

Las claves API pueden restringirse a:

| Ámbito | Campo | Ejemplo |
|--------|-------|---------|
| Proyecto | `project_id` | `p3b2c1a0-...` |
| Host | `hosts` | `["mi-proyecto.hola.cloud"]` |
| Ruta | `paths` | `["/api/v0/lambdas/*"]` |
| Método | `methods` | `["GET", "POST"]` |

Cuando se establecen múltiples ámbitos, **todos** deben coincidir para que la solicitud sea aceptada. Los patrones de ruta admiten comodines estilo glob (`*` coincide con cualquier secuencia).

## Usar una Clave API

Una vez que tienes un Api-Key y un Api-Secret, inclúyelos en todas las solicitudes:

```bash
curl "https://mi-proyecto.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Secret: f0e1d2c3-b4a5-6789-0fed-cba987654321"
```

## Listar Claves API

```bash
curl "https://api.hola.cloud/api/v0/projects/p3b2c1a0-1234-5678-9abc-def012345678/apikeys" \
  -H "Api-Key: TU_API_KEY_ADMIN" \
  -H "Api-Secret: TU_API_SECRET_ADMIN"
```

```json
{
  "api_keys": [
    {
      "api_key": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "project_id": "p3b2c1a0-1234-5678-9abc-def012345678",
      "scopes": { "hosts": ["mi-proyecto.hola.cloud"] },
      "created_at": "2025-06-21T12:00:00Z",
      "last_used_at": "2025-06-21T14:30:00Z"
    }
  ]
}
```

Nota: El `api_secret` nunca se devuelve en las respuestas de listado — solo se muestra una vez en la creación.

## Revocar una Clave API

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/projects/p3b2c1a0-1234-5678-9abc-def012345678/apikeys/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Key: TU_API_KEY_ADMIN" \
  -H "Api-Secret: TU_API_SECRET_ADMIN"
```

Una revocación exitosa devuelve `200 OK`. La clave se invalida inmediatamente. Cualquier solicitud posterior que la use recibirá `401 Unauthorized`.

## Rotar una Clave API

Para rotar una clave, crea un nuevo par con los mismos ámbitos, actualiza tus aplicaciones para usar la nueva clave y luego revoca la clave antigua.

```bash
# Paso 1: Crear una nueva clave
curl -X POST "https://api.hola.cloud/api/v0/projects/ID_DE_TU_PROYECTO/apikeys" \
  -H "Api-Key: TU_API_KEY_ADMIN" \
  -H "Api-Secret: TU_API_SECRET_ADMIN" \
  -H "Content-Type: application/json" \
  -d '{"scopes": {"hosts": ["mi-proyecto.hola.cloud"]}}'

# Paso 2: Actualiza tu aplicación con el nuevo Api-Key y Api-Secret

# Paso 3: Eliminar la clave antigua
curl -X DELETE "https://api.hola.cloud/api/v0/projects/ID_DE_TU_PROYECTO/apikeys/API_KEY_ANTIGUA" \
  -H "Api-Key: TU_API_KEY_ADMIN" \
  -H "Api-Secret: TU_API_SECRET_ADMIN"
```

## Cómo Validan las Claves los Servicios

Los servicios backend no validan las claves API directamente. En su lugar, Glue2 realiza la validación usando el paquete `glueauth`:

1. Glue2 extrae los encabezados `Api-Key` y `Api-Secret` de la solicitud.
2. Busca el registro de la clave por el UUID del Api-Key en InceptionDB.
3. Compara el Api-Secret proporcionado contra el hash bcrypt almacenado.
4. Verifica que la solicitud coincida con los ámbitos de la clave (host, ruta, método).
5. Si es válido, inyecta el encabezado JWT `X-Glue-Authentication` y reenvía la solicitud.
6. Si no es válido, devuelve `401 Unauthorized` con un mensaje de error.

El servicio backend confía en el encabezado `X-Glue-Authentication` y lo usa para decisiones de autorización. Nunca necesita validar la clave API original por sí mismo.

## Recomendaciones de Seguridad

- **Trata el Api-Secret como una contraseña**: Nunca lo registres, lo confirmes en control de versiones ni lo compartas en canales inseguros.
- **Usa claves de corta duración**: Considera rotar las claves cada 90 días.
- **Restringe ámbitos**: Comienza con los ámbitos más estrechos y expande solo cuando sea necesario.
- **Monitorea el uso**: Usa el endpoint `/v0/stats` para monitorear los patrones de uso de claves API.
- **Audita las claves regularmente**: Elimina las claves no utilizadas para reducir la superficie de ataque.
