<h1>认证</h1>

Glue2 支持三种认证机制：**API 密钥**、**会话**和 **OAuth 2.0**。每个端点可以配置为要求认证（`Require`）或可选认证（`Optional`）。对 `Require` 端点的未认证请求将收到 `401 Unauthorized` 响应。

## API 密钥

API 密钥提供机器对机器的认证。每个密钥由两部分组成：**Api-Key**（公共标识符）和 **Api-Secret**（私有密钥）。密钥以 bcrypt 哈希形式存储，创建后无法检索。

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "my-function", "runtime": "js", "code": "..."}'
```

API 密钥可以限定到项目和 host 规则：

- **项目**：密钥仅对配置的项目有效。
- **Host 规则**：限制虚拟主机和可选 host 元数据。

路径和 HTTP 方法作用域不属于当前模型。

### 可选与必需认证

- **Require**：端点拒绝未认证的请求并返回 `401`。
- **Optional**：端点同时接受认证和匿名请求。认证请求会收到注入的标头；匿名请求则直接通过。

当提供有效的 API 密钥时，Glue2 会将 JSON `X-Glue-Authentication` 标头注入代理请求中。

## 会话

基于会话的认证由 HolaCloud 控制台的浏览器用户使用。用户通过 `auth.hola.cloud` 登录时，认证服务会创建一个存储在 InceptionDB 中的会话。会话 ID 作为 HTTP-only、Secure、SameSite cookie 返回。

```bash
# 会话由浏览器自动管理。
# 对于 API 访问，请改用 API 密钥或 OAuth 令牌。
```

会话 cookie 会自动附加到对 HolaCloud 子域名的请求中。Glue2 读取 cookie，在 InceptionDB 中查找会话，并在转发请求之前进行验证。

## OAuth 2.0 / Google 登录

控制台用户可以通过 Google OAuth 2.0 进行认证。流程如下：

1. 用户在控制台登录页面点击"使用 Google 登录"。
2. 浏览器重定向到 Google 的授权端点。
3. 用户授予权限，Google 重定向回带有授权码的页面。
4. 控制台用授权码交换访问令牌和 ID 令牌。
5. ID 令牌发送到 Glue2，Glue2 验证令牌并创建会话。

OAuth 令牌也可以直接用于 API 调用：

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "Authorization: Bearer 您的_OAuth_令牌"
```

## 认证流程总结

```
方法          凭证                    典型用途              范围
──────────────────────────────────────────────────────────────────────
API 密钥     Api-Key + Api-Secret    机器对机器            项目/Host 规则
会话         Cookie (http-only)       浏览器用户            用户身份
OAuth 2.0   Bearer 令牌              控制台 / SSO          Google 身份
```

## 通过控制台管理 API 密钥

API 密钥在 HolaCloud 控制台的 **设置 > API 密钥** 中管理。您可以：

- 创建新的密钥对（Api-Key + Api-Secret）
- 查看现有密钥（密钥仅在创建时显示一次）
- 撤销密钥
- 设置密钥范围

## 安全最佳实践

- **定期轮换密钥**：定期生成新密钥并更新您的应用程序。
- **最小权限原则**：将每个密钥的范围限制到所需的最小项目和 host 规则。
- **仅使用 HTTPS**：始终使用 `https://` —— 永远不要通过明文 HTTP 发送凭证。
- **安全存储密钥**：使用环境变量或密钥管理器。切勿在源代码中硬编码密钥。
- **立即撤销泄露的密钥**：如果密钥暴露，请使用控制台或 API 立即撤销。

## 下一步

- 了解如何创建和管理 API 密钥，请参阅[管理API密钥](../3_api-keys/api-keys_zh.md)。
- 查看架构和路由，请参阅[架构与路由](../1_architecture-routing/architecture-routing_zh.md)。
