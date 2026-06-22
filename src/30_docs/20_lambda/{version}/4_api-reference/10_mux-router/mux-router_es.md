# Mux Router

Enruta solicitudes públicas mediante un path con alcance de propietario. No requiere autenticación.

La ruta mux es `/mux/{owner_id}/*`. El resto del path se reenvía a la lógica de rutas de las lambdas del propietario.

## Parámetros de Path

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `owner_id` | string | Identificador del propietario |
| `*` | path | Resto del path reenviado al alcance del propietario |

## Solicitud HTTP

```http
GET /mux/user_123/hello-world HTTP/1.1
Host: api.hola.cloud
```

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/mux/user_123/hello-world"
```

## Respuesta

La respuesta la produce la lambda seleccionada por la ruta del propietario.

```json
{
  "body": {
    "message": "Hello from mux",
    "path": "/hello-world"
  }
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 404 | Propietario o ruta de lambda no encontrada |
| 500 | Error de ejecución de Lambda |
