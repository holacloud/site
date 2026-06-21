
# POST /v1/databases/{id}/collections/{col}

对集合执行文档操作。具体操作由集合名称后的冒号后缀决定。

## 身份验证

需要 `Api-Key`、`Api-Secret` 和 `X-Project` 请求头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|------|
| id | uuid | 数据库的唯一标识符 |
| col | string | 带操作后缀的集合名称 |

## 操作

### 插入 — `{collection}:insert`

向集合中插入一个或多个文档。

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "Alice",
  "email": "alice@example.com",
  "role": "admin"
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice",
    "email": "alice@example.com",
    "role": "admin"
  }'
```

响应：
```json
{
  "id": "d7e8f9a0-b1c2-3456-7890-123456789abc",
  "message": "文档插入成功"
}
```

### 查找 — `{collection}:find`

使用过滤条件查询文档。

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "filter": {
    "role": "admin"
  }
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "role": "admin"
    }
  }'
```

响应：
```json
{
  "documents": [
    {
      "id": "d7e8f9a0-b1c2-3456-7890-123456789abc",
      "name": "Alice",
      "email": "alice@example.com",
      "role": "admin"
    }
  ]
}
```

### 删除 — `{collection}:remove`

删除匹配过滤条件的文档。

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "filter": {
    "email": "alice@example.com"
  }
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "email": "alice@example.com"
    }
  }'
```

响应：
```json
{
  "deleted": 1,
  "message": "文档删除成功"
}
```

### 更新 — `{collection}:patch`

更新匹配过滤条件的文档。

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "filter": {
    "role": "admin"
  },
  "patch": {
    "role": "superadmin"
  }
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "role": "admin"
    },
    "patch": {
      "role": "superadmin"
    }
  }'
```

响应：
```json
{
  "updated": 1,
  "message": "文档更新成功"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 400 | 无效操作名称或请求体格式错误 |
| 401 | 缺少或无效的身份验证标头 |
| 403 | 项目访问被拒绝 |
| 404 | 未找到数据库或集合 |
