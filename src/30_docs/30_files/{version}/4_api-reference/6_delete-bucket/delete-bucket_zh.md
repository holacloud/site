# DELETE /v1/buckets/{id}

删除空的存储桶。存储桶不能包含任何文件。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `id` | string | 存储桶 ID（例如 `bkt_abc123`） |

## 请求

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

## 响应

```http
HTTP/1.1 204 No Content
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 404 | Not Found | 指定的存储桶不存在 |
| 409 | Conflict | 存储桶不为空；请先删除所有文件 |
| 500 | Internal Server Error | 发生意外错误 |
