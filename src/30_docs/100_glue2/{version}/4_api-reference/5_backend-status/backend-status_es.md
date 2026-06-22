# Estado del Backend

Devuelve un arreglo con el estado de hosts por proyecto.

```bash
curl -X GET "https://api.hola.cloud/v0/status"
```

```json
[
  {
    "id": "project-123",
    "name": "Mi Proyecto",
    "host": "mi-proyecto.hola.cloud",
    "status": 200,
    "statusText": "200 OK"
  }
]
```
