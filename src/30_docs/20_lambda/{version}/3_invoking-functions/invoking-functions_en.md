<h1>Invoking Lambda Functions</h1>

HolaCloud Lambda supports multiple invocation methods to suit different use cases — from direct synchronous calls to webhook integrations and custom domain routing.

## Synchronous Invocation

All invocations are synchronous: HolaCloud executes your function and returns the response once execution completes.

### Admin Invocation (Authenticated)

Use the admin endpoint for internal calls where you control the client:

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/YOUR_FUNCTION_ID" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

### Public Invocation (Unauthenticated)

Use the public endpoint for third-party services or client-side apps:

```bash
curl -X POST "https://api.hola.cloud/run/YOUR_FUNCTION_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

The public endpoint requires no authentication, making it ideal for webhooks and public APIs.

## Webhook Usage

Any Lambda function can serve as a webhook endpoint. Configure your external service (Stripe, GitHub, SendGrid, etc.) to send events to:

```
https://api.hola.cloud/run/YOUR_FUNCTION_ID
```

The function receives the webhook payload as the request body and can return a response that the external service will interpret.

Example: GitHub webhook that logs push events:

```javascript
export default (req) => {
  const event = req.headers["x-github-event"];
  const payload = req.body;

  console.log(`Received ${event} from GitHub`);

  return {
    status_code: 200,
    body: { received: true }
  };
};
```

## The Mux Router

The mux router maps custom domains or subdomains to specific Lambda functions by owner. The route pattern is:

```
GET /mux/{owner_id}/*
```

When a request hits the mux, HolaCloud routes it to the matching Lambda function owned by that user. This enables scenarios such as:

- Multi-tenant SaaS platforms where each tenant gets a subdomain
- Custom domain mapping for user-deployed functions
- Path-based routing to different functions under the same owner

```bash
curl "https://api.hola.cloud/mux/OWNER_ID/any/path/here" \
  -H "Content-Type: application/json"
```

The `owner_id` is the HolaCloud account identifier. The remainder of the path is forwarded to the function, which receives it as part of `req`.

## Monitoring Ongoing Executions

Check which functions are currently running using the public `ongoing` endpoint:

```bash
curl "https://api.hola.cloud/ongoing"
```

Expected response:

```json
[
  {
    "function_id": "f3b2c1a0-1234-5678-9abc-def012345678",
    "function_name": "hello-world",
    "started_at": "2025-06-21T12:00:00Z",
    "duration_ms": 234
  }
]
```

This endpoint is useful for operational monitoring and debugging long-running functions.

## Current User Endpoint

The `/me` endpoint returns information about the currently authenticated user:

```bash
curl "https://api.hola.cloud/me" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

## OpenAPI Specification

HolaCloud exposes a machine-readable OpenAPI specification at:

```
GET https://api.hola.cloud/openapi
```

This can be used with tools like Swagger UI, Postman, or code generators to explore and interact with the Lambda API programmatically.
