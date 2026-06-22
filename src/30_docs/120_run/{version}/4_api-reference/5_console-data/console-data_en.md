# Console Data

Returns Console data for a repository.

## Request

```http
GET /api/console?repository=my-project/my-app
```

The `repository` query parameter is required.

## Example

```bash
curl "https://api.hola.cloud/api/console?repository=my-project/my-app"
```

## Response

The response contains the repository data known to Run, including pushed references and the saved runtime configuration. Exact fields can vary with service state.

```json
{
  "repository": "my-project/my-app",
  "references": ["latest", "v1"],
  "env": [
    {"key": "LOG_LEVEL", "desired_value": "info"}
  ],
  "volumes": [
    {"name": "my-data", "target": "/data", "mode": "rw"}
  ]
}
```
