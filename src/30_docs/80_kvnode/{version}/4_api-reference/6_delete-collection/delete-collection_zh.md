# 删除集合

删除集合及其所有键。

## 身份验证

需要内部身份验证。通过 `X-Glue-Authentication` 头部或 `apikey` 和 `secret` 头部传递凭据。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| col | string | 集合名称 |

## 请求示例

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/users" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## 响应示例

```http
HTTP/1.1 204 No Content
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 401 | unauthorized | 缺少或无效的身份验证 |
| 404 | not_found | 未找到集合 |
| 500 | internal_error | 服务器内部错误 |
