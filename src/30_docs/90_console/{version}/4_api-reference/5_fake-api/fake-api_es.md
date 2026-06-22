# API Falsa

Endpoints fijos que proporcionan datos simulados para la Consola durante el desarrollo y el modo de demostración.

## Descripción

La Consola expone solo las rutas fake API listadas abajo. Las rutas no listadas no forman parte de la fake API de la Consola.

## Sub-rutas

| Método | Ruta | Descripción |
|--------|------|-------------|
| `GET` | `/fakeapi/authapi/auth/me` | Devuelve el usuario simulado actual |
| `GET` | `/fakeapi/inceptionapi/v1/databases` | Lista bases de datos simuladas |
| `GET` | `/fakeapi/inceptionapi/v1/databases/{databaseid}/collections` | Lista colecciones simuladas |
| `GET` | `/fakeapi/lambdasapi/api/v0/lambdas` | Lista lambdas simuladas |
| `GET` | `/fakeapi/projectsapi/v0/projects` | Lista proyectos simulados |
| `GET` | `/fakeapi/projectsapi/v0/projects/{projectid}` | Devuelve un proyecto simulado |
| `GET` | `/fakeapi/filesapi/v1/buckets` | Lista buckets simulados |

## Autenticación

Ninguna. Estos endpoints simulados no requieren autenticación.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/fakeapi/authapi/auth/me"
```

```json
{
  "email": "fulanez@gmail.com",
  "id": "user-1234",
  "nick": "fulanez"
}
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/projectsapi/v0/projects"
```

```json
[
  { "id": "project-00000000-0000-0000-0000-000000000001", "name": "Hello" }
]
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/lambdasapi/api/v0/lambdas"
```

```json
[
  { "id": "a663e1f1-e5cb-4fc2-b846-53f2cc7574c9", "method": "GET", "name": "Index", "path": "/" }
]
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Datos simulados devueltos correctamente |
| 404 | Ruta bajo `/fakeapi/` no reconocida |
