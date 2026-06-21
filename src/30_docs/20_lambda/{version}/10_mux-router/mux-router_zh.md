
# GET /mux/{owner}/*

根据子路径将传入请求路由到指定所有者的 Lambda 函数。此端点公开可用，无需身份验证。

Mux 路由器允许将自定义域名或路径映射到特定的 Lambda 所有者。`/mux/{owner}/` 之后的剩余路径将转发给所有者的 Lambda 路由逻辑。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|------|
| owner | string | 所有者标识符（用户名或项目名） |
| `*` | 路径 | 转发给所有者 Lambda 的剩余路径部分 |

## HTTP 请求

```http
GET /mux/acme-webapp/hello-world HTTP/1.1
Host: api.hola.cloud
```

## 示例

```bash
curl -X GET "https://api.hola.cloud/mux/acme-webapp/hello-world"
```

## 响应

响应完全取决于处理路由请求的 Lambda 函数。

```json
{
  "status": 200,
  "body": {
    "message": "来自 acme-webapp 的 Lambda 的问候！"
  }
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 404 | 未找到所有者或 Lambda 路由 |
| 500 | Lambda 执行错误 |
