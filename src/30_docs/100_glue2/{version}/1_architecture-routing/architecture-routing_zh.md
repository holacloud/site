<h1>架构与路由</h1>

Glue2 是位于所有 HolaCloud 服务前面的中央反向代理。所有传入流量——无论是来自浏览器用户、API 客户端还是内部服务——在到达目的地之前都经过 Glue2。

![Glue2 架构图](img/glue2-architecture.png)

## 反向代理

Glue2 作为第 7 层反向代理运行。它终止 TLS 连接，检查每个请求，验证调用者身份，然后将请求转发到相应的后端服务。没有外部流量直接到达 HolaCloud 服务；Glue2 是唯一的入口点。

```
客户端 ──▶ Glue2（身份验证检查）──▶ 后端服务
              │
              ├──▶ InceptionDB
              ├──▶ Lambda
              ├──▶ Files
              ├──▶ Config
              ├──▶ KVNode
              ├──▶ Scheduler
              └──▶ Logs
```

## 基于虚拟主机的路由

Glue2 根据 `Host` 标头路由请求。每个 HolaCloud 项目被分配一个唯一的子域名（`<项目>.hola.cloud`），Glue2 将该子域名映射到项目的后端服务。

内置平台服务使用保留子域名：

| 主机 | 服务 |
|------|------|
| `inceptiondb.hola.cloud` | InceptionDB |
| `api.hola.cloud` | 公共 API（Lambda、Files、Config 等） |
| `console.hola.cloud` | HolaCloud 控制台 |
| `auth.hola.cloud` | 身份验证 / 会话服务 |
| `logs.hola.cloud` | InstantLogs |

当请求到达 `my-project.hola.cloud` 时，Glue2 查找具有该别名的项目，解析其主后端服务，并转发请求。

## 代理流程

1. 客户端向 `<项目>.hola.cloud` 发送 HTTP 请求。
2. Glue2 终止 TLS 连接并读取 `Host` 标头。
3. Glue2 检查身份验证（会话 cookie、API 密钥或 OAuth 令牌）。如果端点需要身份验证而未提供，则请求被拒绝并返回 `401 Unauthorized`。
4. Glue2 将虚拟主机解析为项目 ID 和后端服务目标。
5. Glue2 将身份验证和项目标头注入请求。
6. 请求被转发到后端服务。
7. 后端响应被流式传回客户端，Glue2 记录事务。

## 注入的标头

转发到后端时，Glue2 注入：

| 标头 | 描述 |
|------|------|
| `X-Glue-Authentication` | 纯 JSON 身份验证上下文 |
| `X-Holacloud-Project-Id` | 从虚拟主机派生的项目 UUID |
| `X-Holacloud-Tenant-Project-Id` | 用于多租户隔离的租户范围项目 ID |
| `X-Forwarded-For` | 原始客户端 IP 地址 |
| `X-Forwarded-Proto` | 原始协议（`http` 或 `https`） |

后端服务使用这些标头识别调用者并执行授权。`X-Glue-Authentication` 是纯 JSON。

## V0 管理端点

Glue2 在 `/v0/` 下公开一组管理端点：

### 列出虚拟主机

```bash
curl "https://api.hola.cloud/v0/virtualhosts" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET"
```

```json
["my-project.hola.cloud", "api.my-project.hola.cloud"]
```

### 流量统计

```bash
curl "https://api.hola.cloud/v0/stats" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET"
```

### 后端健康状态

```bash
curl "https://api.hola.cloud/v0/status" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET"
```

```json
{
  "backends": [
    {
      "name": "inceptiondb",
      "status": "healthy",
      "latency_ms": 12
    }
  ]
}
```

## 下一步

- 了解身份验证机制，请参阅[认证](../2_authentication/authentication_zh.md)。
- 了解如何管理 API 密钥，请参阅[管理API密钥](../3_api-keys/api-keys_zh.md)。
