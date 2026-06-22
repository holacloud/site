
# Delete Lambda

永久删除 Lambda 函数。此操作不可撤销。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 请求头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|------|
| id | uuid | 要删除的 Lambda 函数的唯一标识符 |

## HTTP 请求

```http
DELETE /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## 示例

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf"
```

## 响应

```json
{
  "message": "Lambda 函数已成功删除",
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 401 | 缺少或无效的身份验证标头 |
| 404 | 未找到 Lambda 函数 |
