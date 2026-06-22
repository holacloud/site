<h1>Managing API Keys</h1>

API keys are the primary method for machine-to-machine authentication with HolaCloud services. Each key consists of a UUID-based public identifier (Api-Key) and a UUID-based secret (Api-Secret). The secret is hashed with bcrypt and cannot be retrieved after creation.

## Creating an API Key

### Through the Console

1. Log in to [https://console.hola.cloud](https://console.hola.cloud).
2. Navigate to **Settings > API Keys**.
3. Click **Create API Key**.
4. Optionally set key scopes (projects and host rules).
5. Click **Create**.
6. Copy the Api-Key and Api-Secret immediately — the secret is shown only once.

### Through the Serviceprojects API

API keys are managed by the serviceprojects API. Project creation is outside the current Glue2 API key documentation.

```bash
curl -X POST "https://api.hola.cloud/v0/apikeys" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "CI/CD Key",
    "scopes": [
      {
        "projects": ["project-123"],
        "host_rules": {"my-project.hola.cloud": "{}"}
      }
    ]
  }'
```

Expected response:

```json
{
  "key": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "secret": "f0e1d2c3-b4a5-6789-0fed-cba987654321",
  "name": "CI/CD Key",
  "scopes": [
    {
      "projects": ["project-123"],
      "host_rules": {"my-project.hola.cloud": "{}"}
    }
  ]
}
```

## Key Structure

- **Api-Key**: UUID v4, acts as the public identifier for the key. Sent in the `Api-Key` header.
- **Api-Secret**: UUID v4, acts as the secret. Sent in the `Api-Secret` header. Stored as a bcrypt hash — HolaCloud cannot recover it if lost.

## Scoping

API keys can be restricted to:

| Scope | Field | Example |
|-------|-------|---------|
| Projects | `projects` | `["project-123"]` |
| Host rules | `host_rules` | `{"my-project.hola.cloud": "{}"}` |

Path and HTTP method scopes are not part of the current API key model.

## Using an API Key

Once you have an Api-Key and Api-Secret, include them in all requests:

```bash
curl "https://my-project.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Secret: f0e1d2c3-b4a5-6789-0fed-cba987654321"
```

## Listing API Keys

```bash
curl "https://api.hola.cloud/v0/apikeys" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}"
```

```json
{
  "api_keys": [
    {
      "key": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "name": "CI/CD Key",
      "scopes": [{ "projects": ["project-123"], "host_rules": {"my-project.hola.cloud": "{}"} }]
    }
  ]
}
```

Note: The `api_secret` is never returned in listing responses — it is only shown once at creation.

## Revoking an API Key

```bash
curl -X DELETE "https://api.hola.cloud/v0/apikeys/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}"
```

A successful revocation returns `200 OK`. The key is immediately invalidated. Any subsequent requests using it will receive `401 Unauthorized`.

## Rotating an API Key

To rotate a key, create a new key pair with the same scopes, update your applications to use the new key, then revoke the old key.

```bash
# Step 1: Create a new key
curl -X POST "https://api.hola.cloud/v0/apikeys" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}" \
  -H "Content-Type: application/json" \
  -d '{"name":"Rotated key","scopes":[{"projects":["project-123"],"host_rules":{"my-project.hola.cloud":"{}"}}]}'

# Step 2: Update your application with the new Api-Key and Api-Secret

# Step 3: Delete the old key
curl -X DELETE "https://api.hola.cloud/v0/apikeys/OLD_API_KEY" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}"
```

## How Services Validate Keys

Backend services do not validate API keys directly. Instead, Glue2 performs validation using the `glueauth` package:

1. Glue2 extracts the `Api-Key` and `Api-Secret` headers from the request.
2. It looks up the key record by the Api-Key UUID in InceptionDB.
3. It compares the provided Api-Secret against the stored bcrypt hash.
4. It verifies that the request matches the key's project and host rules.
5. If valid, it injects the JSON `X-Glue-Authentication` header and forwards the request.
6. If invalid, it returns `401 Unauthorized` with an error message.

The backend service trusts the `X-Glue-Authentication` header and uses it for authorization decisions. It never needs to validate the original API key itself.

## Security Recommendations

- **Treat Api-Secret like a password**: Never log it, commit it to version control, or share it in insecure channels.
- **Use short-lived keys**: Consider rotating keys every 90 days.
- **Scope restrictively**: Start with the narrowest scopes and expand only as needed.
- **Monitor usage**: Use the `/v0/stats` endpoint to monitor API key usage patterns.
- **Audit keys regularly**: Delete unused keys to reduce the attack surface.
