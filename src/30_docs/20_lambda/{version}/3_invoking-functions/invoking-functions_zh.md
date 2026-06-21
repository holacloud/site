<h1>调用Lambda函数</h1>

HolaCloud Lambda 支持多种调用方式，以适应不同的使用场景——从直接的同步调用到 Webhook 集成和自定义域名路由。

## 同步调用

所有调用均为同步：HolaCloud 执行您的函数并在执行完成后返回响应。

### 管理端调用（需认证）

使用管理端点进行内部调用（您控制客户端）：

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/您的函数ID" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

### 公共调用（无需认证）

使用公共端点供第三方服务或客户端应用调用：

```bash
curl -X POST "https://api.hola.cloud/run/您的函数ID" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

公共端点无需认证，非常适合 Webhook 和公共 API。

## Webhook 使用

任何 Lambda 函数都可以作为 Webhook 端点。将外部服务（Stripe、GitHub、SendGrid 等）配置为向以下地址发送事件：

```
https://api.hola.cloud/run/您的函数ID
```

函数将 Webhook 负载作为请求体接收，并可返回外部服务能够识别的响应。

示例：记录推送事件的 GitHub Webhook：

```javascript
export default (req) => {
  const event = req.headers["x-github-event"];
  const payload = req.body;

  console.log(`从GitHub收到 ${event} 事件`);

  return {
    status_code: 200,
    body: { received: true }
  };
};
```

## Mux 路由器

Mux 路由器根据所有者将自定义域名或子域名映射到特定的 Lambda 函数。路由模式为：

```
GET /mux/{owner_id}/*
```

当请求到达 Mux 时，HolaCloud 将其路由到该用户拥有的匹配 Lambda 函数。这支持以下场景：

- 多租户 SaaS 平台，每个租户拥有一个子域名
- 用户部署函数的自定义域名映射
- 同一所有者下基于路径路由到不同函数

```bash
curl "https://api.hola.cloud/mux/所有者ID/任意/路径" \
  -H "Content-Type: application/json"
```

`owner_id` 是 HolaCloud 账户标识符。路径的其余部分将转发给函数，函数在 `req` 中接收。

## 监控正在执行的调用

使用公共 `ongoing` 端点查看当前正在运行的函数：

```bash
curl "https://api.hola.cloud/ongoing"
```

预期响应：

```json
[
  {
    "function_id": "f3b2c1a0-1234-5678-9abc-def012345678",
    "function_name": "hello-world",
    "started_at": "2025-06-21T12:00:00Z",
    "duration_ms": 234
  }
]
```

此端点适用于运行监控和长时间运行函数的调试。

## 获取当前用户信息

`/me` 端点返回当前认证用户的信息：

```bash
curl "https://api.hola.cloud/me" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET"
```

## OpenAPI 规范

HolaCloud 在以下地址提供机器可读的 OpenAPI 规范：

```
GET https://api.hola.cloud/openapi
```

可用于 Swagger UI、Postman 或代码生成器等工具，以编程方式探索和交互 Lambda API。
