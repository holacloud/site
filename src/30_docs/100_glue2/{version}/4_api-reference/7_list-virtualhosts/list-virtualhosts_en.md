# List Virtualhosts

Returns the current virtual host routing table.

## Description

This endpoint lists all configured virtual hosts and their corresponding backend service targets. It is used by the HolaCloud Console to display the routing configuration.

## Authentication

None. This endpoint is public.

## Request

No request body required.

## Example

```bash
curl -X GET "https://api.hola.cloud/v0/virtualhosts"
```

## Response

```json
{
  "virtualhosts": [
    {
      "host": "my-project.hola.cloud",
      "target": "inceptiondb",
      "port": 8080,
      "tls": true
    },
    {
      "host": "api.my-project.hola.cloud",
      "target": "lambda",
      "port": 8081,
      "tls": true
    }
  ]
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Routing table returned successfully |
