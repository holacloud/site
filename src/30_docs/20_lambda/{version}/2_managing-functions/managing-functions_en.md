<h1>Managing Lambda Functions</h1>

Once you have created a Lambda function, you can update its code, configure environment variables, inspect its status, or delete it through the API.

## Function Structure

Every Lambda function must export a default handler that receives a request object and returns a response. The handler signature varies slightly by runtime:

### JavaScript (Node.js)

```javascript
export default (req) => {
  // req.body      - parsed request payload
  // req.headers   - object with request headers
  // req.query     - parsed query string parameters
  // req.method    - HTTP method string

  return {
    body: "response body",
    status_code: 200,
    headers: { "X-Custom": "value" }
  };
};
```

### Go

```go
package main

import "encoding/json"

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
        Body:       "Hello from Go",
        StatusCode: 200,
    }, nil
}
```

### Python

```python
def handler(req):
    # req["body"]      - request payload
    # req["headers"]   - dict of headers
    # req["method"]    - HTTP method
    return {
        "body": "Hello from Python",
        "status_code": 200
    }
```

## Updating Function Code

Use `PATCH /api/v0/lambdas/{id}` to modify an existing function. You can update the `name`, `runtime`, `code`, and `environment` fields.

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/YOUR_FUNCTION_ID" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "export default (req) => { return { body: \"Updated!\" } }",
    "environment": {
      "DB_URL": "https://db.example.com",
      "LOG_LEVEL": "debug"
    }
  }'
```

## Setting Environment Variables

Environment variables are passed as a key-value object in the `environment` field when creating or updating a function. They are injected into the function's process at runtime and accessible via the standard language mechanisms (`process.env` in Node.js, `os.Getenv` in Go, `os.environ` in Python).

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/YOUR_FUNCTION_ID" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "environment": {
      "MY_VAR": "my-value"
    }
  }'
```

## Viewing Function Details and Status

Retrieve a single function's full details:

```bash
curl "https://api.hola.cloud/api/v0/lambdas/YOUR_FUNCTION_ID" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

Expected response:

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

The `status` field can be `active`, `inactive`, or `error`.

## Listing All Functions

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

## Deleting a Function

Permanently remove a function and all its associated data:

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/YOUR_FUNCTION_ID" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

A successful deletion returns a `200 OK` with no body.
