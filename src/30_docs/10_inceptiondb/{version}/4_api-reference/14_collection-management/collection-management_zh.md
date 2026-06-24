# 集合管理

这些 action 端点用于管理数据库中的现有集合。

## 认证

需要 `Api-Key` 和 `Api-Secret`，或者在数据库 owner 被允许时使用 Glue owner token。

## 获取集合

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

响应：

```json
{"name":"users","total":5400,"indexes":2,"defaults":{"id":"uuid()"}}
```

## 删除集合

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:dropCollection HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

成功响应为空 body。

## 设置默认值

默认值会在插入新文档时合并到文档中。发送 `null` 可删除某个字段的默认值。

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:setDefaults HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{"id":"uuid()","status":"active","legacy":null}
```

响应：

```json
{"id":"uuid()","status":"active"}
```

## 大小

`size` 是实验性端点，返回集合存储指标。

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:size HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

响应：

```json
{"memory":123456,"disk":234567,"index.id":12345}
```
