<h1>Getting Started with Lambda</h1>

This guide walks you through creating your first Lambda function, invoking it, and inspecting the result.

## Prerequisites

- A HolaCloud account with API credentials (Api-Key and Api-Secret).
- [curl](https://curl.se/) installed on your machine.

## Step 1: Obtain Your API Credentials

Log in to the HolaCloud dashboard and navigate to the API Keys section. Generate a new key pair — you will receive an **Api-Key** and an **Api-Secret**. Keep these values safe; they are used to authenticate all admin requests.

## Step 2: Create a Lambda Function

Create a simple "hello-world" function written in JavaScript:

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "runtime": "js",
    "code": "export default (req) => { return { body: \"Hello, World!\" } }"
  }'
```

Expected response:

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "name": "hello-world",
  "runtime": "js",
  "status": "active",
  "created_at": "2025-06-21T12:00:00Z"
}
```

Save the returned `id` — you will need it to invoke and manage your function.

## Step 3: List Your Lambdas

Verify the function was created by listing all lambdas on your account:

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

## Step 4: Run the Function

Invoke the function synchronously using the admin endpoint:

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/YOUR_FUNCTION_ID" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{}'
```

Expected response:

```json
{
  "body": "Hello, World!",
  "status_code": 200
}
```

You can also invoke the function from a public endpoint (no authentication required):

```bash
curl -X POST "https://api.hola.cloud/run/YOUR_FUNCTION_ID" \
  -H "Content-Type: application/json" \
  -d '{}'
```

## Step 5: Check Active Invocations

List currently running invocations with the public ongoing endpoint:

```bash
curl "https://api.hola.cloud/ongoing"
```

## Next Steps

- Learn how to update your function code and set environment variables in [Managing Lambda Functions](../2_managing-functions/managing-functions_en.md).
- Explore invocation patterns including webhooks and the mux router in [Invoking Lambda Functions](../3_invoking-functions/invoking-functions_en.md).
