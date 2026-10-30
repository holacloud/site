# API key 与所有者

仅限控制台：这些端点需要已登录控制台的用户会话。`Api-Key` 和 `Api-Secret` 对它们无效，也没有受支持的方式从你自己的代码调用它们。这里记录它们是为了说明控制台做了什么。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：**API keys & settings**。如何使用和保护它们，请参阅 [API keys](/zh/docs/inceptiondb/v2/api-keys/)。

## 创建 API key

`POST /v1/databases/{databaseId}:createApiKey`

```json
{"name": "backend"}
```

响应 `201 Created`。`secret` 只在此时返回一次，请妥善保存：

```json
{"name":"backend","key":"4fdd0267-fded-4d51-b1a2-e7674e02126e","secret":"02863644-fbf9-4e57-8b64-e11ead5878ec","creation_date":"2026-10-28T06:27:03.555714157+01:00"}
```

## 撤销 API key

`POST /v1/databases/{databaseId}:deleteApiKey`

```json
{"key": "4fdd0267-fded-4d51-b1a2-e7674e02126e"}
```

响应 `204 No Content`。此后使用该密钥的请求都会被拒绝。

## 添加所有者

`POST /v1/databases/{databaseId}:addOwner`

```json
{"owner_id": "user-2"}
```

响应 `200 OK`，包含所有者列表：

```json
["owner-1","user-2"]
```

## 移除所有者

`POST /v1/databases/{databaseId}:deleteOwner`

```json
{"owner_id": "user-2"}
```

响应 `200 OK`，包含剩余的所有者。你不能移除自己。

```json
["owner-1"]
```
