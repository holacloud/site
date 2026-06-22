<h1>管理 API 密钥</h1>

API 密钥用于机器到机器认证。每个请求使用 `Api-Key` 和 `Api-Secret` 标头；secret 只在创建或轮换时显示，并以哈希形式存储。

## 真实 API

API 密钥由 serviceprojects API 的 `/v0/apikeys` 管理，不在 `/projects/{id}/apikeys` 下管理。项目创建和项目管理 API 当前不在本页范围内。

```bash
curl -X POST "https://api.hola.cloud/v0/apikeys" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "CI/CD Key",
    "scopes": [
      {
        "projects": ["project-123"],
        "host_rules": {"my-project.hola.cloud": "{}"}
      }
    ]
  }'
```

## 作用域

当前模型使用 `projects` 和 `host_rules`。

| 作用域 | 字段 | 示例 |
|--------|------|------|
| 项目 | `projects` | `["project-123"]` |
| Host 规则 | `host_rules` | `{"my-project.hola.cloud": "{}"}` |

路径和 HTTP 方法作用域不属于当前模型。

## 列出密钥

```bash
curl "https://api.hola.cloud/v0/apikeys" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}"
```

## 撤销密钥


```bash
curl -X DELETE "https://api.hola.cloud/v0/apikeys/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-Glue-Authentication: {\"user\":{\"id\":\"user-1234\"}}"
```

## Glue2 如何验证

1. Glue2 提取 `Api-Key` 和 `Api-Secret`。
2. 查找密钥并将 secret 与存储的哈希比较。
3. 验证 `projects` 和 `host_rules`。
4. 如果有效，将 `X-Glue-Authentication` 作为 JSON 注入并转发请求。
