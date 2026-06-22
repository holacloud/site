<h1>Managing Lambda Functions</h1>

After creating a lambda, use the management API to inspect it, update its supported fields, list all lambdas in the account, or delete it.

## Function Structure

JavaScript lambdas export a default handler. The handler receives a request object and returns the response body you want HolaCloud to send back.

```javascript
export default (req) => {
  return {
    body: {
      method: req.method,
      path: req.path,
      headers: req.headers,
      data: req.body
    }
  };
};
```

Static lambdas use one of the static language modes: `static-html`, `static-css`, or `static-js`. In those modes, `code` is the content served for the matching lambda.

## Updating a Lambda

Use `PATCH /api/v0/lambdas/{lambda_id}` to update `name`, `language`, `code`, `method`, or `path`.

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-updated",
    "method": "POST",
    "path": "/hello-updated",
    "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })"
  }'
```

Expected response:

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "created_timestamp": 1750507200,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-updated",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })",
  "method": "POST",
  "path": "/hello-updated"
}
```

## Viewing Lambda Details

Retrieve one lambda by ID:

```bash
curl "https://api.hola.cloud/api/v0/lambdas/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## Listing All Lambdas

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## Deleting a Lambda

Permanently remove a lambda:

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

A successful deletion returns the API response for the deleted lambda or a confirmation payload, depending on the deployed API version.
