# List Virtualhosts

Returns the current virtual host routing table.

## Description

This endpoint lists the configured virtual host names.

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
["my-project.hola.cloud", "api.my-project.hola.cloud"]
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Routing table returned successfully |
