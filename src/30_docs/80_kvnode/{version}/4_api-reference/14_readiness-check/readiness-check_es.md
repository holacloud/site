# Disponibilidad Check

Devuelve el estado de disponibilidad del nodo KVNode, incluyendo la salud de su conexión padre.

## Autenticación

Este endpoint es público. No requiere autenticación.

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/readyz"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "status": "ok",
  "parent_connected": true
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 503 | unavailable | El nodo no está disponible |
