# 访问管理

数据库访问管理 action 用于创建 API key 并管理 owner 用户 ID。

## 认证

需要数据库 owner 用户的 `X-Glue-Authentication`。

## 创建 API key

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:createApiKey HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"name":"production"}
```

响应状态：`201 Created`

```json
{
  "name": "production",
  "key": "1abbe476-6ad6-4b97-9cca-6deb6ab2901d",
  "secret": "4bda6d52-762b-4e5d-bed7-85614c13b8bf",
  "creation_date": "2026-06-24T12:00:00Z"
}
```

secret 只会在创建 key 时返回。

## 删除 API key

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:deleteApiKey HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"key":"1abbe476-6ad6-4b97-9cca-6deb6ab2901d"}
```

成功响应为空 body。

## 添加 owner

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:addOwner HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"owner_id":"user-123"}
```

响应：

```json
["current-owner","user-123"]
```

## 删除 owner

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:deleteOwner HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"owner_id":"user-123"}
```

响应：

```json
["current-owner"]
```

认证用户不能把自己从 owner 列表中删除。
