
# Modify Lambda

更新现有 Lambda 函数的代码、运行环境或激活状态。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 请求头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|------|
| id | uuid | Lambda 函数的唯一标识符 |

## 请求体

| 字段 | 类型 | 描述 |
|-------|------|------|
| name | string | 新名称（可选） |
| runtime | string | 新运行环境（可选） |
| code | string | 新源代码（可选） |
| active | boolean | Lambda 是否激活（可选） |

## HTTP 请求

```http
PATCH /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "code": "export default async (req) => { return { status: 200, body: { message: 'Updated lambda' } }; }",
  "active": false
}
```

## 示例

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "export default async (req) => { return { status: 200, body: { message: \"Updated lambda\" } }; }",
    "active": false
  }'
```

## 响应

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "name": "hello-world",
  "runtime": "javascript",
  "active": false,
  "updated_at": "2025-07-02T09:15:00Z"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 400 | 无效的请求体 |
| 401 | 缺少或无效的身份验证标头 |
| 404 | 未找到 Lambda 函数 |
