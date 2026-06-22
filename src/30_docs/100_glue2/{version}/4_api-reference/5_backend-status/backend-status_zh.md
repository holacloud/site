# 后端状态

返回项目 host 状态数组。

```bash
curl -X GET "https://api.hola.cloud/v0/status"
```

```json
[
  {
    "id": "project-123",
    "name": "My Project",
    "host": "my-project.hola.cloud",
    "status": 200,
    "statusText": "200 OK"
  }
]
```
