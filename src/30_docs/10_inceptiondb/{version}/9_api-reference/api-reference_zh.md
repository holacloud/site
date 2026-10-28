# API 参考

基础 URL：`https://inceptiondb.hola.cloud`

## 认证

| 端点 | 凭据 |
|---|---|
| 数据：`/v1/databases/{databaseId}/collections` 下的所有端点 | 该数据库某个 [API key](/zh/docs/inceptiondb/v2/api-keys/) 的 `Api-Key` 和 `Api-Secret` 请求头。 |
| 账户：数据库、API key 和所有者 | 已登录 hola.cloud 用户的会话。请在 [控制台](https://inceptiondb.hola.cloud/) 中使用。 |

示例从环境变量读取凭据：

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

## 格式

- 请求发送 JSON。插入发送 **JSON Lines**：每行一个 JSON 对象。
- `:find`、`:insert`、`:patch` 和 `:remove` 返回 **JSON Lines**，每行一个文档。其他端点返回 JSON。
- 没有匹配的查询返回 `204 No Content`，响应体为空。
- 错误返回 JSON 对象：`{"error": {"message": "...", "description": "..."}}`。

## 端点

### 数据

| 方法 | 路径 | 作用 |
|---|---|---|
| GET | `/v1/databases/{databaseId}/collections` | [列出集合](/zh/docs/inceptiondb/v2/api-reference/list-collections/) |
| POST | `/v1/databases/{databaseId}/collections` | [创建集合](/zh/docs/inceptiondb/v2/api-reference/create-collection/) |
| GET | `/v1/databases/{databaseId}/collections/{collection}` | [获取集合](/zh/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:setDefaults` | [设置默认值](/zh/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:size` | [内存与磁盘大小](/zh/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:dropCollection` | [删除集合](/zh/docs/inceptiondb/v2/api-reference/collection-management/) |
| POST | `…/collections/{collection}:insert` | [插入文档](/zh/docs/inceptiondb/v2/api-reference/insert-documents/) |
| POST | `…/collections/{collection}:find` | [查找文档](/zh/docs/inceptiondb/v2/api-reference/find-documents/) |
| GET | `…/collections/{collection}/documents/{id}` | [按 id 获取文档](/zh/docs/inceptiondb/v2/api-reference/get-document/) |
| POST | `…/collections/{collection}:patch` | [更新文档](/zh/docs/inceptiondb/v2/api-reference/patch-documents/) |
| POST | `…/collections/{collection}:remove` | [删除文档](/zh/docs/inceptiondb/v2/api-reference/remove-documents/) |
| POST | `…/collections/{collection}:createIndex` | [创建索引](/zh/docs/inceptiondb/v2/api-reference/index-management/) |
| POST | `…/collections/{collection}:listIndexes` | [列出索引](/zh/docs/inceptiondb/v2/api-reference/index-management/) |
| POST | `…/collections/{collection}:getIndex` | [获取索引](/zh/docs/inceptiondb/v2/api-reference/index-management/) |
| POST | `…/collections/{collection}:dropIndex` | [删除索引](/zh/docs/inceptiondb/v2/api-reference/index-management/) |

`…` 表示 `/v1/databases/{databaseId}`。

### 账户

| 方法 | 路径 | 作用 |
|---|---|---|
| GET | `/v1/databases` | [列出你的数据库](/zh/docs/inceptiondb/v2/api-reference/list-databases/) |
| POST | `/v1/databases` | [创建数据库](/zh/docs/inceptiondb/v2/api-reference/create-database/) |
| GET | `/v1/databases/{databaseId}` | [获取数据库](/zh/docs/inceptiondb/v2/api-reference/get-database/) |
| PATCH | `/v1/databases/{databaseId}` | [重命名数据库](/zh/docs/inceptiondb/v2/api-reference/update-database/) |
| DELETE | `/v1/databases/{databaseId}` | [删除数据库](/zh/docs/inceptiondb/v2/api-reference/drop-database/) |
| POST | `/v1/databases/{databaseId}:createApiKey` | [创建 API key](/zh/docs/inceptiondb/v2/api-reference/access-management/) |
| POST | `/v1/databases/{databaseId}:deleteApiKey` | [撤销 API key](/zh/docs/inceptiondb/v2/api-reference/access-management/) |
| POST | `/v1/databases/{databaseId}:addOwner` | [添加所有者](/zh/docs/inceptiondb/v2/api-reference/access-management/) |
| POST | `/v1/databases/{databaseId}:deleteOwner` | [移除所有者](/zh/docs/inceptiondb/v2/api-reference/access-management/) |

整个 API 的 OpenAPI 描述位于 [`/openapi.json`](https://inceptiondb.hola.cloud/openapi.json)。

## 状态码

| 状态码 | 何时 |
|---|---|
| 200 OK | 请求成功并有响应体。 |
| 201 Created | 创建了集合、文档或索引。 |
| 204 No Content | 请求成功但没有内容返回：没有匹配的文档，或是删除操作。 |
| 400 Bad Request | JSON 无效、文档不是 JSON 对象，或字段名重复。 |
| 401 Unauthorized | 未登录就调用了账户端点。 |
| 403 Forbidden | API key 或 secret 错误，或无权访问该数据库。 |
| 404 Not Found | 文档不存在。 |
| 409 Conflict | 唯一索引中已存在该值。 |
| 500 Internal Server Error | 集合不存在，或创建时已存在，也会返回此状态。 |
