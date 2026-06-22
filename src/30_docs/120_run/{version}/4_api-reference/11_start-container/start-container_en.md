# Start Repository

Starts a repository at a pushed image reference or digest.

## Request Body

Use `repository` plus either `reference` or `digest`.

```json
{
  "repository": "my-project/my-app",
  "reference": "latest"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "reference": "latest"
  }'
```

Digest form:

```json
{
  "repository": "my-project/my-app",
  "digest": "sha256:a1b2c3d4..."
}
```
