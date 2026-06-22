# 删除 Lambda

永久删除一个 lambda。

## 认证

需要 `X-Glue-Authentication`。

## Path 参数

| 参数 | 类型 | 描述 |
|------|------|------|
| `lambda_id` | string | Lambda 标识符 |

## HTTP 请求

```http
DELETE /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## 示例

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## 响应

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789"
}
```

## 错误码

| 代码 | 描述 |
|------|------|
| 401 | 认证缺失或无效 |
| 404 | Lambda 未找到 |
