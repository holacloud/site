# Simular Api

Endpoint comodín que proporciona datos simulados para todos los servicios backend de la Consola durante el desarrollo y el modo de demostración.

## Descripción

La ruta `/fakeapi/*` coincide con cualquier ruta bajo `/fakeapi/` y devuelve respuestas simuladas. Esto permite que la interfaz de la Consola funcione sin un backend real. La estructura de rutas refleja el diseño real de la API para que el código de la interfaz sea idéntico entre desarrollo y producción.

## Sub-rutas

| Sub-ruta | Descripción |
|----------|-------------|
| `/fakeapi/auth/me` | Devuelve la sesión de usuario simulada actual |
| `/fakeapi/inceptionapi` | Simula respuestas de la API de InceptionDB |
| `/fakeapi/lambdasapi` | Simula la gestión de funciones Lambda |
| `/fakeapi/projectsapi` | Simula la configuración y listado de proyectos |
| `/fakeapi/filesapi` | Simula operaciones de almacenamiento de archivos |

## Autenticación

Ninguna. Estos endpoints simulados no requieren autenticación.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/fakeapi/auth/me"
```

```json
{
  "user": "demo@holacloud.com",
  "name": "Usuario de Demostración",
  "project": "default",
  "role": "admin",
  "authenticated": true
}
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/projectsapi"
```

```json
{
  "projects": [
    {
      "id": "default",
      "name": "Proyecto por Defecto",
      "plan": "free",
      "region": "us-east"
    }
  ]
}
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/lambdasapi"
```

```json
{
  "functions": [
    {
      "id": "fn-001",
      "name": "hello-world",
      "runtime": "nodejs18",
      "status": "active"
    }
  ]
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Datos simulados devueltos correctamente |
| 404 | Ruta bajo `/fakeapi/` no reconocida |
