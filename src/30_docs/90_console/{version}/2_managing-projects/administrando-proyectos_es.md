# Administrando Proyectos y Servicios

Los proyectos agrupan instancias de servicios y configuración de enrutamiento. Las APIs de administración de proyectos de la Consola están actualmente fuera del alcance de esta documentación.

## Cambiar entre proyectos

Use el selector de proyectos en la Consola. En modo demo/desarrollo, la lista proviene de la fake API real:

```bash
curl -X GET "https://console.hola.cloud/fakeapi/projectsapi/v0/projects"
```

## Servicios simulados disponibles

```bash
curl -X GET "https://console.hola.cloud/fakeapi/inceptionapi/v1/databases"
curl -X GET "https://console.hola.cloud/fakeapi/lambdasapi/api/v0/lambdas"
curl -X GET "https://console.hola.cloud/fakeapi/filesapi/v1/buckets"
```

## Cobertura de API

Las APIs de proyectos de producción están actualmente fuera del alcance de estos documentos. La gestión real de claves API pertenece a serviceprojects en `/v0/apikeys`.
