<h1>调用 Lambda 函数</h1>

HolaCloud Lambda 支持直接管理调用、公开调用、按所有者划分的 mux 路由，以及少量服务检查端点。

## 管理调用

当客户端可以发送 `X-Glue-Authentication` 时，使用 `/api/v0/run/{lambda_id}`。该端点接受任何 HTTP 方法。

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/YOUR_LAMBDA_ID" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

## 公开调用

对 webhook、浏览器调用和不应发送管理凭据的其他客户端，使用 `/run/{lambda_id}`。该端点接受任何 HTTP 方法。

```bash
curl -X POST "https://api.hola.cloud/run/YOUR_LAMBDA_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "value"
  }'
```

lambda 通过 `req` 接收请求数据，包括方法、path、headers、query 值和 body。

## Webhook 用法

将外部服务配置为向以下地址发送事件：

```text
https://api.hola.cloud/run/YOUR_LAMBDA_ID
```

Webhook 处理器示例：

```javascript
export default (req) => {
  return {
    body: {
      received: true,
      event: req.headers["x-github-event"],
      payload: req.body
    }
  };
};
```

## Mux Router

Mux router 通过 `/mux/{owner_id}/*` 转发按所有者划分的路由。它接受任何 HTTP 方法。

```bash
curl -X GET "https://api.hola.cloud/mux/OWNER_ID/any/path/here"
```

`owner_id` 标识 HolaCloud 所有者。剩余 path 会转发给 lambda 路由逻辑。

## 进行中的调用

查看当前正在运行的调用：

```bash
curl "https://api.hola.cloud/ongoing"
```

## 当前用户

`/me` 端点返回已认证用户的信息：

```bash
curl "https://api.hola.cloud/me" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## OpenAPI 规范

HolaCloud 在以下地址公开 Lambda OpenAPI 文档：

```text
GET https://api.hola.cloud/openapi
```
