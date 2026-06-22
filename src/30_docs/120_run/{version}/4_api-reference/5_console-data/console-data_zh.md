# Console 数据

返回某个仓库的 Console 数据。

## Request

```http
GET /api/console?repository=my-project/my-app
```

`repository` query parameter 是必需的。

## 示例

```bash
curl "https://api.hola.cloud/api/console?repository=my-project/my-app"
```

## Response

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
