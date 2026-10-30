# 索引

没有索引时，查询会读取集合中的所有文档。几千个文档时这很快。如果有数百万个文档、需要有序结果，或要确保某个值永不重复，请在你查询的字段上创建索引。

**开始之前：** 本指南使用 [快速开始](/zh/docs/inceptiondb/v2/getting-started/) 留下的 `customers` 集合（包含 Ada 和 Grace）及其环境变量。它会创建四个索引，并在最后删除它们。文档始终不会被修改。

## 我需要哪种索引？

| 类型 | 用途 | 值唯一 | 查询参数 |
|---|---|---|---|
| `map` | 按精确值查找文档：邮箱、用户名、外部 id。 | 是，始终唯一 | `value` |
| `btree` | 范围和排序：日期、金额、分数。值可以重复。 | 可选（`unique`） | `from`、`to`、`reverse` |
| `fts` | 按文本中的词查找文档：名称、标题、描述。 | 否 | `match` |

索引保存在内存中，每次写入时更新。`map` 索引以及设置了 `unique` 的 `btree` 索引会在每次写入前检查，因此重复的值会被拒绝。

## 按邮箱查找（map）

创建索引：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email", "type": "map", "field": "email"}'
```

```json
{"field":"email","name":"by_email","sparse":false,"type":"map"}
```

按邮箱查找客户：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"index": "by_email", "value": "grace@example.com"}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

从现在起，插入一个邮箱已存在的客户会失败，返回 `409 Conflict`：

```json
{"error":{"description":"Unexpected error","message":"index add 'by_email': index conflict: field 'email' with value 'ada@example.com'"}}
```

`description` 是通用文本；有用的细节始终在 `message` 中：索引、字段和重复的值。该文档不会被保存。

要让文档的 `id` 也保持唯一，请在 `id` 上创建 `map` 索引。之后按 id 读取会使用它：其 `source` 为 `index`，而不是 `fullscan`。

## 范围和排序（btree）

在月收入上创建索引：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_mrr", "type": "btree", "fields": ["mrr"]}'
```

```json
{"fields":["mrr"],"name":"by_mrr","sparse":false,"type":"btree","unique":false}
```

付费在 100 到 500 之间的客户，从低到高：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"index": "by_mrr", "from": {"mrr": 100}, "to": {"mrr": 500}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

收入最高的客户，从高到低：`"reverse": true`。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"index": "by_mrr", "reverse": true, "limit": 3}'
```

注意：

- `from` 和 `to` 是可选的；省略其中一个即为开放范围。
- `fields` 可以包含多个字段，例如 `["country", "-mrr"]`：先按国家，每个国家内按收入从高到低（`-` 表示降序）。
- 也可以同时加上 `filter`：索引决定顺序和范围，过滤器再排除文档。
- `"unique": true` 会拒绝重复值，与 `map` 索引一样。

## 按词搜索（fts）

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "search_name", "type": "fts", "field": "name"}'
```

查找名字中包含 `grace` 的客户，不区分大小写：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"index": "search_name", "match": "grace", "limit": 10}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## 没有该字段的文档

只有当集合中已有的每个文档都包含该字段时，才能创建 `map` 或 `btree` 索引。否则响应为 `500`，消息类似 ``error indexing existing data: field `company` is indexed and mandatory``。

对于只有部分文档才有的字段（例如 `company`），创建索引时设置 `"sparse": true`。没有该字段的文档会正常保存，但不属于该索引：使用该索引的查询不会返回它们，不带 `index` 的查询仍能找到它们。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_company", "type": "btree", "fields": ["company"], "sparse": true}'
```

## 列出索引

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:listIndexes" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"fields":["company"],"name":"by_company","sparse":true,"type":"btree","unique":false},{"field":"email","name":"by_email","sparse":false,"type":"map"},{"fields":["mrr"],"name":"by_mrr","sparse":false,"type":"btree","unique":false},{"field":"name","name":"search_name","type":"fts"}]
```

## 清理

删除索引不会影响文档。响应为 `204 No Content`。删除本指南创建的四个索引：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_mrr"}'
```

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email"}'
```

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "search_name"}'
```

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_company"}'
```

使用已存在的名称创建索引会返回 `500`，消息为 `index 'by_email' already exists`。如果你之前运行过本指南但没有清理，请运行这四条命令，然后重新开始。

在 [控制台](https://inceptiondb.hola.cloud/) 中，**Schema** 标签页可以一键在任意字段上创建索引。
