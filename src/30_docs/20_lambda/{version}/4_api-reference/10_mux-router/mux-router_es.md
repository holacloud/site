
# Mux Router

Enruta las solicitudes entrantes a las funciones lambda del propietario especificado según la sub-ruta. Este es un endpoint público, no requiere autenticación.

El Mux Router permite mapear dominios personalizados o rutas a propietarios de lambda específicos. La ruta restante después de `/mux/{owner}/` se reenvía a la lógica de enrutamiento del propietario.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| owner | string | Identificador del propietario (nombre de usuario o proyecto) |
| `*` | ruta | El resto de la ruta reenviada a las lambdas del propietario |

## Solicitud HTTP

```http
GET /mux/acme-webapp/hello-world HTTP/1.1
Host: api.hola.cloud
```

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/mux/acme-webapp/hello-world"
```

## Respuesta

La respuesta depende completamente de la función lambda que maneja la solicitud enrutada.

```json
{
  "status": 200,
  "body": {
    "message": "¡Hola desde la lambda de acme-webapp!"
  }
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 404 | Propietario o ruta lambda no encontrada |
| 500 | Error de ejecución de la lambda |
