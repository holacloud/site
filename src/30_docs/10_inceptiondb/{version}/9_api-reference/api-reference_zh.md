# API 参考

基础 URL：`https://inceptiondb.hola.cloud`

## 认证

| 端点 | 凭据 |
|---|---|
| 数据：`/v1/databases/{databaseId}/collections` 下的所有端点 | 该数据库某个 [API key](/zh/docs/inceptiondb/v2/api-keys/) 的 `Api-Key` 和 `Api-Secret` 请求头。 |
| 账户：数据库、API key 和所有者 | **仅限控制台。** 已登录 [控制台](https://inceptiondb.hola.cloud/) 的用户会话。`Api-Key` 和 `Api-Secret` 对它们无效。 |

示例从环境变量读取凭据：

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

## 格式

| 端点 | 请求体 | 响应体 |
|---|---|---|
| `:insert` | JSON Lines（JSONL）：每行一个 JSON 对象，不是 JSON 数组。发送 `Content-Type: application/x-ndjson`。 | JSONL：已保存的文档。 |
| `:find`、`:patch`、`:remove` | JSON。发送 `Content-Type: application/json`。 | JSONL：匹配的文档。 |
| 其他所有端点 | JSON（有请求体时）。 | JSON。 |

- InceptionDB 根据端点而不是 `Content-Type` 解析请求体，但示例始终会发送它，以便任何客户端或代理都知道格式。
- `204 No Content` 的响应体为空：没有文档匹配，或删除了某些内容。解码响应体之前请先检查状态码。
- 错误返回 JSON 对象：`{"error": {"description": "...", "message": "..."}}`。`description` 通常是通用的 `Unexpected error`；详细信息在 `message` 中。

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

### 账户（仅限控制台）

这些端点需要已登录控制台的用户会话。列出它们是为了说明控制台做了什么，而不是用于自动化：API key 无法调用它们。见 [API key](/zh/docs/inceptiondb/v2/api-keys/)。

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
| 200 OK | 请求成功，并且有响应体。 |
| 201 Created | 创建了集合、文档或索引。 |
| 204 No Content | 请求成功，但没有内容可返回：没有匹配的文档，或是删除操作。响应体为空。 |
| 400 Bad Request | 请求体无效：见下面的错误。 |
| 401 Unauthorized | 在未登录会话的情况下调用账户端点。 |
| 403 Forbidden | API key 或 secret 错误，或无权访问该数据库。 |
| 404 Not Found | 集合、文档或索引不存在。 |
| 409 Conflict | 唯一索引中已存在该值。 |
| 500 Internal Server Error | 目前一些已知错误会返回 500：见下面的错误。 |

## 错误及处理方法

错误的 `message` 说明了原因。以下是你可能遇到的错误及其真实消息：

| 状态 | `message` | 原因 | 处理方法 |
|---|---|---|---|
| 403 | ``can not access database '…'`` | 密钥、secret 或数据库 ID 不正确，或密钥已被撤销。 | 检查 `INCEPTIONDB_DATABASE`、`INCEPTIONDB_API_KEY` 和 `INCEPTIONDB_API_SECRET`。 |
| 401 | `unauthorized` | 没有控制台会话就调用账户端点。 | 在 [控制台](https://inceptiondb.hola.cloud/) 中操作。 |
| 400 | `document must be a JSON object` | `:insert` 发送了 JSON 数组或不是对象的值。 | 每行发送一个对象（JSONL）。 |
| 400 | ``jsontext: unexpected EOF …``、``jsontext: duplicate object member name …`` | `:insert` 中某个文档不是有效 JSON，或重复了字段名。 | 修正这一行。之前的文档已经保存。 |
| 400 | `unexpected end of JSON input`（description 为 `Malformed JSON`） | `:find`、`:patch` 或 `:remove` 的请求体不是有效 JSON。 | 检查引号、逗号和括号。 |
| 404 | `collection not found` | `GET` 一个不存在的集合。 | 检查名称：区分大小写。 |
| 404 | ``document '…' not found`` | 按不存在的 id `GET` 文档。 | 检查集合和 id；两者都区分大小写。要按其他字段查找，请使用 `:find`。 |
| 404 | ``index '…' not found in collection '…'`` | `:getIndex` 一个不存在的索引。 | 用 `:listIndexes` 列出索引。 |
| 409 | ``index add '…': index conflict: field '…' with value '…'`` | 该值已存在于 `map` 索引或设置了 `unique` 的 `btree` 索引中。 | 使用其他值，或更新已有文档。 |
| 500 | ``collection `…` already exists`` | 用已存在集合的名称 `POST /collections`。 | 使用该集合，或换一个名称。 |
| 500 | `collection not found` | 对不存在的集合执行 `:find`、`:patch` 或 `:remove`。 | 检查名称。注意 `:insert` 会创建集合。 |
| 500 | ``collection `…` not found in database `…` `` | `:dropCollection` 一个不存在的集合。 | 没有需要删除的内容。 |
| 500 | ``index '…' already exists`` | 用已存在索引的名称 `:createIndex`。 | 换一个名称，或先删除旧索引。 |
| 500 | ``error indexing existing data: …`` | 在部分文档没有的字段上创建 `map` 或 `btree` 索引，或在重复的值上创建唯一索引。 | 加上 `"sparse": true`，或修正文档。 |
| 500 | ``index '…' not found, available indexes […]`` | `:find` 使用了不存在的 `index`。 | 使用消息中列出的索引之一。 |
| 204 | 无 | 过滤器使用了不存在的运算符，例如 `$gte`。 | 使用 `$ge` 和 `$le`。见 [查询](/zh/docs/inceptiondb/v2/querying/)。 |

上面的 `500` 响应是已知的请求错误，而不是服务故障。任何其他 `500` 都是我们这边的问题。
