# 启动仓库

使用已推送的 reference 或 digest 启动仓库。

```json
{
  "repository": "my-project/my-app",
  "reference": "latest"
}
```

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "reference": "latest"}'
```

Digest 形式：

```json
{
  "repository": "my-project/my-app",
  "digest": "sha256:a1b2c3d4..."
}
```
