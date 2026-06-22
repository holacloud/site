# 回滚镜像

将仓库回滚到已推送的 reference 或 digest。

## Body

```json
{
  "repository": "my-project/my-app",
  "reference": "v1.0.0"
}
```

也可以使用 `digest`：

```json
{
  "repository": "my-project/my-app",
  "digest": "sha256:a1b2c3d4..."
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/api/console/rollback" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "reference": "v1.0.0"}'
```
