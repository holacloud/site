<h1>Authentication</h1>

Glue2 supports three authentication mechanisms: **API Keys**, **Sessions**, and **OAuth 2.0**. Each endpoint can be configured to require authentication (`Require`) or make it optional (`Optional`). Unauthenticated requests to `Require` endpoints receive a `401 Unauthorized` response.

## API Keys

API keys provide machine-to-machine authentication. Each key is a two-part credential consisting of an **Api-Key** (public identifier) and an **Api-Secret** (private secret). The secret is stored as a bcrypt hash and cannot be retrieved after creation.

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "my-function", "runtime": "js", "code": "..."}'
```

API keys can be scoped to specific projects and host rules:

- **Projects**: Key is valid only for the configured project scope.
- **Host rules**: Restrict to configured virtual hosts and optional host metadata.

Path and HTTP method scopes are not part of the current API key model.

### Optional vs Required Auth

- **Require**: The endpoint rejects unauthenticated requests with `401`.
- **Optional**: The endpoint accepts both authenticated and anonymous requests. Authenticated requests receive injected headers; anonymous requests pass through without them.

When a valid API key is provided, Glue2 injects the `X-Glue-Authentication` header into the proxied request as JSON, allowing the backend to identify the caller.

## Sessions

Session-based authentication is used by browser users of the HolaCloud Console. When a user logs in through `auth.hola.cloud`, the authentication service creates a session stored in InceptionDB. The session ID is returned as an HTTP-only, Secure, SameSite cookie.

```bash
# Sessions are managed automatically by the browser.
# For API access, use API keys or OAuth tokens instead.
```

Session cookies are automatically attached to requests to HolaCloud subdomains. Glue2 reads the cookie, looks up the session in InceptionDB, and validates it before forwarding the request.

## OAuth 2.0 / Google Auth

Console users can authenticate via Google OAuth 2.0. The flow works as follows:

1. User clicks "Sign in with Google" on the Console login page.
2. Browser redirects to Google's authorization endpoint.
3. User grants permission and Google redirects back with an authorization code.
4. The Console exchanges the code for an access token and ID token.
5. The ID token is sent to Glue2, which validates it and creates a session.

OAuth bearer tokens can be used directly for API calls:

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "Authorization: Bearer YOUR_OAUTH_TOKEN"
```

## Auth Flow Summary

```
Method       Credentials              Typical Use        Scope
────────────────────────────────────────────────────────────────
API Key     Api-Key + Api-Secret     Machine-to-machine  Project/Host rules
Session     Cookie (http-only)       Browser users       User identity
OAuth 2.0   Bearer token             Console / SSO       Google identity
```

## Managing API Keys Through the Console

API keys are managed in the HolaCloud Console under **Settings > API Keys**. From there you can:

- Create new key pairs (Api-Key + Api-Secret)
- View existing keys (secret is shown once at creation)
- Revoke keys
- Set key scopes

## Security Best Practices

- **Rotate keys regularly**: Generate new keys and update your applications periodically.
- **Least privilege**: Scope each key to the minimum set of projects and host rules required.
- **HTTPS only**: Always use `https://` — never send credentials over plain HTTP.
- **Store secrets securely**: Use environment variables or a secrets manager. Never hardcode secrets in source code.
- **Revoke compromised keys immediately**: Use the Console or API to revoke keys if they are exposed.

## Next Steps

- Learn how to create and manage API keys in [Managing API Keys](../3_api-keys/api-keys_en.md).
- Review architecture and routing in [Architecture & Routing](../1_architecture-routing/architecture-routing_en.md).
