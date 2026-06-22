# Salud Check

Devuelve el estado de salud del nodo KVNode.

## Autenticación

Este endpoint es público. No requiere autenticación.

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/healthz"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "status": "ok"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 500 | internal_error | El nodo no está saludable |
