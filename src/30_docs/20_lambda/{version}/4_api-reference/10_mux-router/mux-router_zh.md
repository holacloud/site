# Mux Router

通过带所有者范围的 path 路由公开请求。不需要认证。

mux 路由为 `/mux/{owner_id}/*`。剩余 path 会转发给所有者的 lambda 路由逻辑。

## Path 参数

| 参数 | 类型 | 描述 |
|------|------|------|
| `owner_id` | string | 所有者标识符 |
| `*` | path | 转发到所有者范围的剩余 path |

## HTTP 请求

```http
GET /mux/user_123/hello-world HTTP/1.1
Host: api.hola.cloud
```

## 示例

```bash
curl -X GET "https://api.hola.cloud/mux/user_123/hello-world"
```

## 响应

响应由所有者路由选中的 lambda 产生。

```json
{
  "body": {
    "message": "Hello from mux",
    "path": "/hello-world"
  }
}
```

## 错误码

| 代码 | 描述 |
|------|------|
| 404 | 所有者或 lambda 路由未找到 |
| 500 | Lambda 执行错误 |
