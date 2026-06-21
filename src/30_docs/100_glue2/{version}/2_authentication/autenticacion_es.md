<h1>Autenticación</h1>

Glue2 soporta tres mecanismos de autenticación: **Claves API**, **Sesiones** y **OAuth 2.0**. Cada endpoint puede configurarse para requerir autenticación (`Require`) o hacerla opcional (`Optional`). Las solicitudes no autenticadas a endpoints `Require` reciben una respuesta `401 Unauthorized`.

## Claves API

Las claves API proporcionan autenticación máquina a máquina. Cada clave consta de dos partes: un **Api-Key** (identificador público) y un **Api-Secret** (secreto privado). El secreto se almacena como un hash bcrypt y no se puede recuperar después de la creación.

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "mi-funcion", "runtime": "js", "code": "..."}'
```

Las claves API pueden limitarse a:

- **Proyecto**: La clave solo es válida para solicitudes al proyecto propietario.
- **Host**: Restringir a un host virtual específico.
- **Ruta**: Restringir a un patrón de prefijo de ruta (ej: `/api/v0/lambdas/*`).
- **Método**: Restringir a métodos HTTP específicos (`GET`, `POST`, etc.).

Si se establece algún ámbito, todas las condiciones deben coincidir. Una clave sin ámbito tiene acceso completo.

### Autenticación Opcional vs Requerida

- **Require**: El endpoint rechaza solicitudes no autenticadas con `401`.
- **Optional**: El endpoint acepta tanto solicitudes autenticadas como anónimas. Las solicitudes autenticadas reciben encabezados inyectados; las anónimas pasan sin ellos.

Cuando se proporciona una clave API válida, Glue2 inyecta el encabezado `X-Glue-Authentication` en la solicitud proxy, permitiendo que el backend identifique al cliente.

## Sesiones

La autenticación basada en sesiones es utilizada por los usuarios del navegador en la Consola de HolaCloud. Cuando un usuario inicia sesión a través de `auth.hola.cloud`, el servicio de autenticación crea una sesión almacenada en InceptionDB. El ID de sesión se devuelve como una cookie HTTP-only, Secure y SameSite.

```bash
# Las sesiones son gestionadas automáticamente por el navegador.
# Para acceso por API, usa claves API o tokens OAuth.
```

Las cookies de sesión se adjuntan automáticamente a las solicitudes a los subdominios de HolaCloud. Glue2 lee la cookie, busca la sesión en InceptionDB y la valida antes de reenviar la solicitud.

## OAuth 2.0 / Google Auth

Los usuarios de la Consola pueden autenticarse mediante Google OAuth 2.0. El flujo funciona de la siguiente manera:

1. El usuario hace clic en "Iniciar sesión con Google" en la página de inicio de sesión de la Consola.
2. El navegador redirige al endpoint de autorización de Google.
3. El usuario concede permiso y Google redirige de vuelta con un código de autorización.
4. La Consola intercambia el código por un token de acceso y un token ID.
5. El token ID se envía a Glue2, que lo valida y crea una sesión.

Los tokens OAuth también pueden usarse directamente para llamadas API:

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "Authorization: Bearer TU_TOKEN_OAUTH"
```

## Resumen del Flujo de Autenticación

```
Método        Credenciales              Uso típico           Ámbito
───────────────────────────────────────────────────────────────────────
Clave API    Api-Key + Api-Secret     Máquina a máquina     Proyecto/Host/Ruta/Método
Sesión       Cookie (http-only)       Usuarios de navegador Identidad de usuario
OAuth 2.0    Token Bearer             Consola / SSO         Identidad de Google
```

## Gestión de Claves API a Través de la Consola

Las claves API se gestionan en la Consola de HolaCloud en **Configuración > Claves API**. Desde allí puedes:

- Crear nuevos pares de claves (Api-Key + Api-Secret)
- Ver claves existentes (el secreto se muestra solo una vez al crearlo)
- Revocar claves
- Establecer ámbitos de clave

## Mejores Prácticas de Seguridad

- **Rota las claves regularmente**: Genera nuevas claves y actualiza tus aplicaciones periódicamente.
- **Mínimo privilegio**: Limita cada clave al conjunto mínimo de hosts, rutas y métodos necesarios.
- **Solo HTTPS**: Usa siempre `https://` — nunca envíes credenciales por HTTP plano.
- **Almacena secretos de forma segura**: Usa variables de entorno o un gestor de secretos. Nunca hardcodees secretos en el código fuente.
- **Revoca claves comprometidas inmediatamente**: Usa la Consola o la API para revocar claves si se exponen.

## Siguientes Pasos

- Aprende a crear y gestionar claves API en [Gestión de Claves API](../3_api-keys/gestionando-claves-api_es.md).
- Revisa la arquitectura y el enrutamiento en [Arquitectura y Enrutamiento](../1_architecture-routing/arquitectura-y-enrutamiento_es.md).
