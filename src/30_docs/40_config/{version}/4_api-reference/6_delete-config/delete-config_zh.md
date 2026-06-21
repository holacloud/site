# DELETE /v0/configs/{id}

按 ID 删除配置。此端点为公开访问（管理 API）。

## 身份验证

无需身份验证。此为公开端点。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `id` | string | 配置 ID（例如 `cfg_abc123`） |

## 请求

```bash
curl -X DELETE "https://api.hola.cloud/v0/configs/cfg_abc123"
```

## 响应

```http
HTTP/1.1 204 No Content
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 404 | Not Found | 指定的配置不存在 |
| 500 | Internal Server Error | 发生意外错误 |
