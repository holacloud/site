# Readiness Check

Devuelve el estado de disponibilidad del servicio.

## Solicitud

```http
GET /readyz
```

## Ejemplo

```bash
curl "https://api.hola.cloud/readyz"
```

## Respuesta

Devuelve `ok` con un código de estado 200 cuando el servicio está listo para recibir tráfico.

```text
ok
```
