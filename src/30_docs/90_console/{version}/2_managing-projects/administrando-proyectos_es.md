
# Administrando Proyectos y Servicios

Los proyectos son la unidad organizativa principal en HolaCloud. Agrupan instancias de servicios, miembros del equipo, claves API y facturación. Esta guía explica cómo crear y gestionar proyectos a través de la Consola.

## Crear y cambiar entre proyectos

Para crear un nuevo proyecto, abra el selector de proyectos en la barra lateral y haga clic en **Nuevo Proyecto**. Proporcione un nombre y una descripción opcional. Una vez creado, la Consola cambia automáticamente al nuevo proyecto.

```bash
# Crear un proyecto mediante la API de la Consola
curl -X POST "https://console.hola.cloud/api/v1/projects" \
  -H "Authorization: Bearer <su-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mi App de Producción",
    "description": "Entorno de producción para mi aplicación"
  }'
```

```json
{
  "id": "prod-app-123",
  "name": "Mi App de Producción",
  "description": "Entorno de producción para mi aplicación",
  "createdAt": "2026-06-20T12:00:00Z"
}
```

Para cambiar de proyecto, haga clic en el nombre del proyecto actual en la barra lateral y seleccione uno diferente del menú desplegable.

```bash
# Listar todos los proyectos accesibles
curl -X GET "https://console.hola.cloud/api/v1/projects" \
  -H "Authorization: Bearer <su-token>"
```

```json
{
  "projects": [
    { "id": "default", "name": "Default", "role": "admin" },
    { "id": "prod-app-123", "name": "Mi App de Producción", "role": "admin" }
  ]
}
```

## Administrar instancias de servicios por proyecto

Cada proyecto contiene sus propias instancias de los servicios de HolaCloud. Navegue a la sección del servicio (por ejemplo, Bases de Datos, Funciones, Archivos, Configuración, Registros, Colas, Programador) para gestionar las instancias dentro del proyecto actual.

Por ejemplo, para listar las bases de datos InceptionDB en el proyecto actual:

```bash
curl -X GET "https://console.hola.cloud/api/v1/projects/prod-app-123/databases" \
  -H "Authorization: Bearer <su-token>"
```

```json
{
  "databases": [
    { "id": "db-001", "name": "users-db", "engine": "inceptiondb", "status": "active" },
    { "id": "db-002", "name": "analytics-db", "engine": "inceptiondb", "status": "active" }
  ]
}
```

El mismo patrón aplica para funciones Lambda, buckets de Files, almacenes de Config, streams de InstantLogs, Colas y trabajos del Programador.

## Gestión de claves API por proyecto

Las claves API autentican el acceso programático a los servicios de HolaCloud. Adminístrelas en **Configuración del Proyecto > Claves API**.

```bash
# Crear una clave API
curl -X POST "https://console.hola.cloud/api/v1/projects/prod-app-123/api-keys" \
  -H "Authorization: Bearer <su-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Clave CI/CD",
    "scopes": ["lambda:invoke", "files:read", "config:read"]
  }'
```

```json
{
  "id": "key-abc123",
  "name": "Clave CI/CD",
  "key": "hc_api_abc123...",
  "scopes": ["lambda:invoke", "files:read", "config:read"],
  "createdAt": "2026-06-20T14:00:00Z"
}
```

Puede revocar claves en cualquier momento desde la misma página de configuración.

## Ver facturación y uso del proyecto

La sección **Facturación** muestra las métricas de uso actuales, el historial de facturas y el desglose de costos por servicio.

```bash
curl -X GET "https://console.hola.cloud/api/v1/projects/prod-app-123/billing" \
  -H "Authorization: Bearer <su-token>"
```

```json
{
  "currentMonth": {
    "period": "2026-06",
    "total": 125.40,
    "services": {
      "inceptiondb": 45.00,
      "lambda": 60.00,
      "files": 15.40,
      "config": 5.00
    }
  },
  "invoices": [
    { "id": "inv-001", "period": "2026-05", "total": 110.20, "status": "paid" }
  ]
}
```

## Gestión de miembros del equipo

Invite miembros del equipo desde **Configuración del Proyecto > Equipo**. Asigne roles: Administrador, Desarrollador o Visor.

```bash
curl -X POST "https://console.hola.cloud/api/v1/projects/prod-app-123/team" \
  -H "Authorization: Bearer <su-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "colega@empresa.com",
    "role": "developer"
  }'
```

```json
{
  "id": "team-456",
  "email": "colega@empresa.com",
  "role": "developer",
  "status": "pending"
}
```

## Importar/exportar configuración como JSON

Exporte la configuración completa de su proyecto (configuraciones de servicios, metadatos de claves API, equipo) como un único archivo JSON. Navegue a **Configuración del Proyecto > Importar/Exportar**.

```bash
# Exportar configuración
curl -X GET "https://console.hola.cloud/api/v1/projects/prod-app-123/export" \
  -H "Authorization: Bearer <su-token>"
```

```json
{
  "project": { "name": "Mi App de Producción" },
  "services": {
    "inceptiondb": { "databases": [...] },
    "lambda": { "functions": [...] },
    "files": { "buckets": [...] },
    "config": { "entries": [...] }
  },
  "team": [...]
}
```

Para importar, suba un archivo JSON con el mismo formato a través de la interfaz de la Consola o la API:

```bash
curl -X POST "https://console.hola.cloud/api/v1/projects/prod-app-123/import" \
  -H "Authorization: Bearer <su-token>" \
  -H "Content-Type: application/json" \
  -d @project-config.json
```

```json
{
  "status": "imported",
  "servicesUpdated": ["inceptiondb", "lambda", "files", "config"]
}
```

## Resumen

Los proyectos en HolaCloud proporcionan entornos aislados para sus aplicaciones. A través de la Consola puede crear y cambiar proyectos, gestionar instancias de servicios, claves API, facturación, miembros del equipo e importar/exportar configuración — todo desde una única interfaz.
