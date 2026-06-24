# Startup Check

Devuelve el estado de inicio del nodo KVNode.

## Autenticación

Este endpoint es público. No requiere autenticación.

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/startupz"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "ok": true,
  "node": "node-abc123",
  "role": "leader"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 500 | internal_error | El nodo no está saludable |
