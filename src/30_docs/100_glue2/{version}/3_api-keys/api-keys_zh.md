<h1>管理API密钥</h1>

API 密钥是与 HolaCloud 服务进行机器对机器认证的主要方法。每个密钥由一个基于 UUID 的公共标识符（Api-Key）和一个基于 UUID 的密钥（Api-Secret）组成。密钥以 bcrypt 哈希形式存储，创建后无法检索。

## 创建 API 密钥

### 通过控制台

1. 登录 [https://console.hola.cloud](https://console.hola.cloud)。
2. 导航到 **设置 > API 密钥**。
3. 点击 **创建 API 密钥**。
4. 可选地设置密钥范围（项目、主机、路径、方法）。
5. 点击 **创建**。
6. 立即复制 Api-Key 和 Api-Secret —— 密钥仅显示一次。

### 通过 Serviceprojects API

API 密钥属于项目。首先，创建一个项目（如果您还没有）：

```bash
curl -X POST "https://api.hola.cloud/api/v0/projects" \
  -H "Api-Key: 您的管理员_API_KEY" \
  -H "Api-Secret: 您的管理员_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "我的项目",
    "slug": "my-project"
  }'
```

预期响应：

```json
{
  "id": "p3b2c1a0-1234-5678-9abc-def012345678",
  "name": "我的项目",
  "slug": "my-project",
  "created_at": "2025-06-21T12:00:00Z"
}
```

然后为项目生成 API 密钥：

```bash
curl -X POST "https://api.hola.cloud/api/v0/projects/p3b2c1a0-1234-5678-9abc-def012345678/apikeys" \
  -H "Api-Key: 您的管理员_API_KEY" \
  -H "Api-Secret: 您的管理员_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "scopes": {
      "hosts": ["my-project.hola.cloud"],
      "paths": ["/api/v0/lambdas/*"],
      "methods": ["GET", "POST"]
    }
  }'
```

预期响应：

```json
{
  "api_key": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "api_secret": "f0e1d2c3-b4a5-6789-0fed-cba987654321",
  "project_id": "p3b2c1a0-1234-5678-9abc-def012345678",
  "scopes": {
    "hosts": ["my-project.hola.cloud"],
    "paths": ["/api/v0/lambdas/*"],
    "methods": ["GET", "POST"]
  },
  "created_at": "2025-06-21T12:00:00Z"
}
```

## 密钥结构

- **Api-Key**：UUID v4，作为密钥的公共标识符。通过 `Api-Key` 标头发送。
- **Api-Secret**：UUID v4，作为密钥。通过 `Api-Secret` 标头发送。以 bcrypt 哈希形式存储 —— HolaCloud 无法在丢失后恢复。

## 范围

API 密钥可以限制到：

| 范围 | 字段 | 示例 |
|------|------|------|
| 项目 | `project_id` | `p3b2c1a0-...` |
| 主机 | `hosts` | `["my-project.hola.cloud"]` |
| 路径 | `paths` | `["/api/v0/lambdas/*"]` |
| 方法 | `methods` | `["GET", "POST"]` |

当设置了多个范围时，**所有**条件必须匹配才能接受请求。路径模式支持 glob 风格的通配符（`*` 匹配任意序列）。

## 使用 API 密钥

获得 Api-Key 和 Api-Secret 后，在每次请求中包含它们：

```bash
curl "https://my-project.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Secret: f0e1d2c3-b4a5-6789-0fed-cba987654321"
```

## 列出 API 密钥

```bash
curl "https://api.hola.cloud/api/v0/projects/p3b2c1a0-1234-5678-9abc-def012345678/apikeys" \
  -H "Api-Key: 您的管理员_API_KEY" \
  -H "Api-Secret: 您的管理员_API_SECRET"
```

```json
{
  "api_keys": [
    {
      "api_key": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "project_id": "p3b2c1a0-1234-5678-9abc-def012345678",
      "scopes": { "hosts": ["my-project.hola.cloud"] },
      "created_at": "2025-06-21T12:00:00Z",
      "last_used_at": "2025-06-21T14:30:00Z"
    }
  ]
}
```

注意：`api_secret` 永远不会在列表响应中返回 —— 仅创建时显示一次。

## 撤销 API 密钥

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/projects/p3b2c1a0-1234-5678-9abc-def012345678/apikeys/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Key: 您的管理员_API_KEY" \
  -H "Api-Secret: 您的管理员_API_SECRET"
```

成功撤销返回 `200 OK`。密钥立即失效。任何后续使用该密钥的请求将收到 `401 Unauthorized`。

## 轮换 API 密钥

要轮换密钥，创建具有相同范围的新密钥对，更新应用程序使用新密钥，然后撤销旧密钥。

```bash
# 步骤 1：创建新密钥
curl -X POST "https://api.hola.cloud/api/v0/projects/您的项目ID/apikeys" \
  -H "Api-Key: 您的管理员_API_KEY" \
  -H "Api-Secret: 您的管理员_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"scopes": {"hosts": ["my-project.hola.cloud"]}}'

# 步骤 2：使用新的 Api-Key 和 Api-Secret 更新应用程序

# 步骤 3：删除旧密钥
curl -X DELETE "https://api.hola.cloud/api/v0/projects/您的项目ID/apikeys/旧_API_KEY" \
  -H "Api-Key: 您的管理员_API_KEY" \
  -H "Api-Secret: 您的管理员_API_SECRET"
```

## 服务如何验证密钥

后端服务不直接验证 API 密钥。相反，Glue2 使用 `glueauth` 包进行验证：

1. Glue2 从请求中提取 `Api-Key` 和 `Api-Secret` 标头。
2. 在 InceptionDB 中通过 Api-Key UUID 查找密钥记录。
3. 将提供的 Api-Secret 与存储的 bcrypt 哈希进行比较。
4. 验证请求是否匹配密钥的范围（主机、路径、方法）。
5. 如果有效，注入 `X-Glue-Authentication` JWT 标头并转发请求。
6. 如果无效，返回 `401 Unauthorized` 和错误消息。

后端服务信任 `X-Glue-Authentication` 标头，并用于授权决策。它本身不需要验证原始 API 密钥。

## 安全建议

- **将 Api-Secret 视为密码**：切勿记录它、提交到版本控制或在非安全渠道中共享。
- **使用短期密钥**：考虑每 90 天轮换一次密钥。
- **限制范围**：从最窄的范围开始，仅在需要时扩展。
- **监控使用情况**：使用 `/v0/stats` 端点监控 API 密钥使用模式。
- **定期审计密钥**：删除未使用的密钥以减小攻击面。
