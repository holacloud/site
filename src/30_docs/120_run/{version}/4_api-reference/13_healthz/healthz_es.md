# Health Check

Devuelve el estado de salud del servicio.

## Solicitud

```http
GET /healthz
```

## Ejemplo

```bash
curl "https://api.hola.cloud/healthz"
```

## Respuesta

Devuelve `ok` con un código de estado 200 cuando el servicio está saludable.

```text
ok
```
