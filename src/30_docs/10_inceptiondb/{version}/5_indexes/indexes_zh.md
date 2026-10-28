# 索引

没有索引时，查询会读取集合中的所有文档。几千个文档时这很快。如果有数百万个文档、需要有序结果，或要确保某个值永不重复，请在你查询的字段上创建索引。

示例使用 `customers` 集合以及 [快速开始](/zh/docs/inceptiondb/v2/getting-started/) 中的环境变量。

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
  -d '{"index": "by_email", "value": "grace@example.com"}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

从现在起，插入一个邮箱已存在的客户会失败，返回 `409 Conflict`：

```json
{"error":{"description":"Unexpected error","message":"index add 'by_email': index conflict: field 'email' with value 'ada@example.com'"}}
```

## 范围和排序（btree）

在月收入上创建索引：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
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
  -d '{"name": "search_name", "type": "fts", "field": "name"}'
```

查找名字中包含 `grace` 的客户，不区分大小写：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"index": "search_name", "match": "grace", "limit": 10}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## 没有该字段的文档

没有被索引字段的文档会正常保存，但不属于该索引：使用该索引的查询不会返回它们。不带 `index` 的查询仍能找到它们。创建索引时设置 `"sparse": true` 可以明确表达这一点：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "by_company", "type": "btree", "fields": ["company"], "sparse": true}'
```

## 列出和删除索引

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:listIndexes" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"fields":["company"],"name":"by_company","sparse":true,"type":"btree","unique":false},{"field":"email","name":"by_email","sparse":false,"type":"map"},{"fields":["mrr"],"name":"by_mrr","sparse":false,"type":"btree","unique":false},{"field":"name","name":"search_name","type":"fts"}]
```

删除索引不会影响文档（`204 No Content`）：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "by_mrr"}'
```

在 [控制台](https://inceptiondb.hola.cloud/) 中，**Schema** 标签页可以一键在任意字段上创建索引。
