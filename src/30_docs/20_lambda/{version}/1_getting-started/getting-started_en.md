<h1>Getting Started with Lambda</h1>

This guide walks you through creating a JavaScript lambda, listing it, and invoking it through the admin and public routes.

## Prerequisites

- A HolaCloud account with an `X-Glue-Authentication` token.
- [curl](https://curl.se/) installed on your machine.

## Step 1: Create a Lambda

Create a simple `hello-world` lambda:

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "language": "javascript",
    "method": "GET",
    "path": "/hello-world",
    "code": "export default (req) => ({ body: { message: \"Hello, World!\", method: req.method, path: req.path } })"
  }'
```

Expected response:

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "created_timestamp": 1750507200,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-world",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Hello, World!\", method: req.method, path: req.path } })",
  "method": "GET",
  "path": "/hello-world"
}
```

Save the returned `id`; it is the `lambda_id` used by the run, get, update, and delete endpoints.

## Step 2: List Lambdas

Verify the lambda was created:

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## Step 3: Run the Lambda

Invoke it through the authenticated admin route:

```bash
curl -X GET "https://api.hola.cloud/api/v0/run/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

Expected response:

```json
{
  "body": {
    "message": "Hello, World!",
    "method": "GET",
    "path": "/hello-world"
  }
}
```

You can also invoke the lambda through its public run route:

```bash
curl -X GET "https://api.hola.cloud/run/YOUR_LAMBDA_ID"
```

## Step 4: Check Ongoing Invocations

List currently running invocations:

```bash
curl "https://api.hola.cloud/ongoing"
```

## Next Steps

- Learn how to update code and route fields in [Managing Lambda Functions](../2_managing-functions/managing-functions_en.md).
- Explore public run routes, mux routing, `/me`, and `/openapi` in [Invoking Lambda Functions](../3_invoking-functions/invoking-functions_en.md).
