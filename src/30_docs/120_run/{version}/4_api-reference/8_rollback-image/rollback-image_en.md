# Rollback Image

Rolls a repository back to a pushed reference or digest.

## Request Body

Use `repository` plus either `reference` or `digest`.

```json
{
  "repository": "my-project/my-app",
  "reference": "v1.0.0"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/console/rollback" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "reference": "v1.0.0"
  }'
```

Digest form:

```json
{
  "repository": "my-project/my-app",
  "digest": "sha256:a1b2c3d4..."
}
```
