# Me

Returns the current authenticated user information. Requires the `X-Glue-Authentication` header.

```bash
curl "https://api.hola.cloud/me" \
  -H "X-Glue-Authentication: <token>"
```

```json
{
  "session": {
    "id": "session_123"
  },
  "user": {
    "id": "user_456",
    "nick": "johndoe",
    "picture": "https://example.com/avatar.png",
    "email": "john@example.com"
  }
}
```
