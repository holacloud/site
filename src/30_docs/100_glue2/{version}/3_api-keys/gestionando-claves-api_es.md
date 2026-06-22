<h1>Gestión de Claves API</h1>

Las claves API autentican acceso máquina a máquina. Cada clave usa los encabezados `Api-Key` y `Api-Secret`; el secreto se muestra solo al crear o rotar secretos y se almacena hasheado.

## API real

Las claves API se gestionan en la API de serviceprojects bajo `/v0/apikeys`. No se gestionan bajo `/projects/{id}/apikeys`. La creación y gestión de proyectos está fuera del alcance actual de esta página.

```bash
curl -X POST "https://api.hola.cloud/v0/apikeys" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "CI/CD Key",
    "scopes": [
      {
        "projects": ["project-123"],
        "host_rules": {"mi-proyecto.hola.cloud": "{}"}
      }
    ]
  }'
```

## Ámbitos

El modelo actual usa `projects` y `host_rules`.

| Ámbito | Campo | Ejemplo |
|--------|-------|---------|
| Proyectos | `projects` | `["project-123"]` |
| Reglas de host | `host_rules` | `{"mi-proyecto.hola.cloud": "{}"}` |

Los ámbitos de ruta y método HTTP no forman parte del modelo actual.

## Listar claves

```bash
curl "https://api.hola.cloud/v0/apikeys" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}"
```

## Revocar una clave

```bash
curl -X DELETE "https://api.hola.cloud/v0/apikeys/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}"
```

## Cómo valida Glue2

1. Glue2 extrae `Api-Key` y `Api-Secret`.
2. Busca la clave y compara el secreto con el hash almacenado.
3. Verifica `projects` y `host_rules`.
4. Si es válido, inyecta `X-Glue-Authentication` como JSON y reenvía la solicitud.
