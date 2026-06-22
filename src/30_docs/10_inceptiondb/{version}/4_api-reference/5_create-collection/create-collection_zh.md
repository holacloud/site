
# Create Collection

在指定数据库中创建新集合。集合是文档的容器。

## 身份验证

需要 `Api-Key`、`Api-Secret` 和 `X-Project` 请求头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|------|
| id | uuid | 数据库的唯一标识符 |

## 请求体

| 字段 | 类型 | 描述 |
|-------|------|------|
| name | string | 集合的名称 |

## HTTP 请求

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "my-collection"
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-collection"
  }'
```

## 响应

```json
{
  "defaults": {
    "id": "uuid()"
  },
  "indexes": 0,
  "name": "my-collection",
  "total": 0
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 400 | 缺少或无效的 `name` 字段 |
| 401 | 缺少或无效的身份验证标头 |
| 403 | 项目访问被拒绝 |
| 404 | 未找到数据库 |
| 409 | 同名集合已存在 |
