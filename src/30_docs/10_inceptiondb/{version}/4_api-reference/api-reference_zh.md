# InceptionDB API 参考

Base URL: `https://api.hola.cloud`

## 认证

`GET /v1/databases` 和 `POST /v1/databases` management 端点使用 `X-Glue-Authentication`。

数据库访问端点使用 `Api-Key` 和 `Api-Secret`。当数据库 owner 被允许时，也可以使用 Glue owner token。

## 端点

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/v1/databases` | 列出数据库 |
| POST | `/v1/databases` | 创建数据库 |
| GET | `/v1/databases/{databaseId}` | 根据 ID 获取数据库 |
| DELETE | `/v1/databases/{databaseId}` | 删除数据库 |
| PATCH | `/v1/databases/{databaseId}` | 更新数据库 |
| POST | `/v1/databases/{databaseId}:createApiKey` | 创建数据库 API key |
| POST | `/v1/databases/{databaseId}:deleteApiKey` | 删除数据库 API key |
| POST | `/v1/databases/{databaseId}:addOwner` | 添加数据库 owner |
| POST | `/v1/databases/{databaseId}:deleteOwner` | 删除数据库 owner |
| GET | `/v1/databases/{databaseId}/collections` | 列出数据库中的集合 |
| POST | `/v1/databases/{databaseId}/collections` | 创建集合 |
| GET | `/v1/databases/{databaseId}/collections/{collection}` | 获取集合 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:dropCollection` | 删除集合 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:setDefaults` | 设置集合默认值 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:size` | 获取集合大小指标 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insert` | 插入文档 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insertStream` | 从流插入文档 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:insertFullduplex` | 通过 full-duplex 流插入文档 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:find` | 查询文档 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:patch` | 修改文档 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:remove` | 删除文档 |
| GET | `/v1/databases/{databaseId}/collections/{collection}/documents/{documentId}` | 根据 ID 获取文档 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:listIndexes` | 列出集合索引 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:createIndex` | 创建集合索引 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:getIndex` | 获取集合索引 |
| POST | `/v1/databases/{databaseId}/collections/{collection}:dropIndex` | 删除集合索引 |
| GET | `/doc` | 获取服务路由树 |
| GET | `/v1` | 获取 v1 欢迎响应 |
| GET | `/release` | 获取服务 release 元数据 |
| GET | `/openapi.json` | 获取 OpenAPI 文档 |
| GET | `/v1/debug/dbs` | Debug 数据库列表 |

## 常见错误码

| 代码 | 描述 |
|------|------|
| 400 | 请求错误：语法或参数无效 |
| 401 | 未授权：缺少或无效凭据 |
| 403 | 禁止访问：凭据无权访问资源 |
| 404 | 未找到：请求的资源不存在 |
| 409 | 冲突：资源已存在 |
| 500 | 服务器内部错误 |
