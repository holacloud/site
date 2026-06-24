# 我的信息

返回当前已认证用户的信息。需要 `X-Glue-Authentication` 头部。

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
