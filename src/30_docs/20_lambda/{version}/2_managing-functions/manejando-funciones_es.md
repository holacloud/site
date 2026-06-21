<h1>Gestionando Funciones Lambda</h1>

Una vez creada una función Lambda, puedes actualizar su código, configurar variables de entorno, consultar su estado o eliminarla a través de la API.

## Estructura de la Función

Cada función Lambda debe exportar un manejador por defecto que recibe un objeto de solicitud y devuelve una respuesta. La firma del manejador varía ligeramente según el runtime:

### JavaScript (Node.js)

```javascript
export default (req) => {
  // req.body      - payload de la solicitud
  // req.headers   - objeto con los encabezados
  // req.query     - parámetros de consulta
  // req.method    - método HTTP

  return {
    body: "cuerpo de la respuesta",
    status_code: 200,
    headers: { "X-Custom": "valor" }
  };
};
```

### Go

```go
package main

type Request struct {
    Body    interface{} `json:"body"`
    Headers map[string]string `json:"headers"`
    Method  string `json:"method"`
}

type Response struct {
    Body    interface{} `json:"body"`
    StatusCode int `json:"status_code"`
}

func Handler(req Request) (Response, error) {
    return Response{
        Body:       "Hola desde Go",
        StatusCode: 200,
    }, nil
}
```

### Python

```python
def handler(req):
    # req["body"]      - payload
    # req["headers"]   - dict de encabezados
    # req["method"]    - método HTTP
    return {
        "body": "Hola desde Python",
        "status_code": 200
    }
```

## Actualizar el Código de la Función

Usa `PATCH /api/v0/lambdas/{id}` para modificar una función existente. Puedes actualizar los campos `name`, `runtime`, `code` y `environment`.

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/ID_DE_TU_FUNCION" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "export default (req) => { return { body: \"¡Actualizado!\" } }",
    "environment": {
      "DB_URL": "https://db.example.com",
      "LOG_LEVEL": "debug"
    }
  }'
```

## Configurar Variables de Entorno

Las variables de entorno se definen como un objeto clave-valor en el campo `environment` al crear o actualizar una función. Se inyectan en el proceso de la función en tiempo de ejecución y son accesibles mediante los mecanismos estándar de cada lenguaje (`process.env` en Node.js, `os.Getenv` en Go, `os.environ` en Python).

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/ID_DE_TU_FUNCION" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "environment": {
      "MY_VAR": "mi-valor"
    }
  }'
```

## Consultar Detalles y Estado

Obtén los detalles completos de una función:

```bash
curl "https://api.hola.cloud/api/v0/lambdas/ID_DE_TU_FUNCION" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

Respuesta esperada:

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "name": "hello-world",
  "runtime": "js",
  "status": "active",
  "environment": {
    "LOG_LEVEL": "debug"
  },
  "created_at": "2025-06-21T12:00:00Z",
  "updated_at": "2025-06-21T14:30:00Z"
}
```

El campo `status` puede ser `active`, `inactive` o `error`.

## Listar Todas las Funciones

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

## Eliminar una Función

Elimina permanentemente una función y todos sus datos asociados:

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/ID_DE_TU_FUNCION" \
  -H "Api-Key: TU_API_KEY" \
  -H "Api-Secret: TU_API_SECRET"
```

Una eliminación exitosa devuelve un `200 OK` sin cuerpo.
