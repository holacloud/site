<h1>Invoking Lambda Functions</h1>

HolaCloud Lambda supports direct admin calls, public calls, owner-scoped mux routes, and a small set of service inspection endpoints.

## Admin Invocation

Use `/api/v0/run/{lambda_id}` when the client can send `X-Glue-Authentication`. The endpoint accepts any HTTP method.

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

## Public Invocation

Use `/run/{lambda_id}` for webhooks, browser calls, and other clients that should not send admin credentials. The endpoint accepts any HTTP method.

```bash
curl -X POST "https://api.hola.cloud/run/YOUR_LAMBDA_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

The lambda receives request data through `req`, including the request method, path, headers, query values, and body.

## Webhook Usage

Configure an external service to send events to:

```text
https://api.hola.cloud/run/YOUR_LAMBDA_ID
```

Example webhook handler:

```javascript
export default (req) => {
  return {
    body: {
      received: true,
      event: req.headers["x-github-event"],
      payload: req.body
    }
  };
};
```

## Mux Router

The mux router forwards owner-scoped routes through `/mux/{owner_id}/*`. It accepts any HTTP method.

```bash
curl -X GET "https://api.hola.cloud/mux/OWNER_ID/any/path/here"
```

The `owner_id` identifies the HolaCloud owner. The remaining path is forwarded to the lambda routing logic.

## Ongoing Invocations

Check currently running invocations:

```bash
curl "https://api.hola.cloud/ongoing"
```

## Current User

The `/me` endpoint returns information for the authenticated user:

```bash
curl "https://api.hola.cloud/me" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## OpenAPI Specification

HolaCloud exposes the Lambda OpenAPI document at:

```text
GET https://api.hola.cloud/openapi
```
