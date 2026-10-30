# 查询

`:find` 以 JSON Lines（JSONL）返回匹配过滤器的文档：每行一个文档。

**开始之前：** 本指南读取 [快速开始](/zh/docs/inceptiondb/v2/getting-started/) 留下的 `customers` 集合（包含 Ada 和 Grace），并使用其中的环境变量。它不会修改任何数据。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"plan": "pro"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
```

请求体：

| 字段 | 默认值 | 含义 |
|---|---|---|
| `filter` | `{}`（全部） | 文档必须满足的条件。 |
| `limit` | `1` | 返回的文档数量。`-1` 返回全部。 |
| `skip` | `0` | 先跳过的匹配文档数，用于分页。 |
| `index` | 无 | 使用 [索引](/zh/docs/inceptiondb/v2/indexes/) 响应查询。 |

没有任何匹配时，返回 `204 No Content`，响应体为空。

## 过滤器

过滤器是一个 JSON 对象。每个字段是一个条件，所有条件都必须满足。

**等于某个值**：`{"plan": "pro", "country": "GB"}`。

使用运算符**比较**：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"mrr": {"$ge": 100, "$lt": 500}}, "limit": 10}'
```

| 运算符 | 字段满足条件时匹配 | 示例 |
|---|---|---|
| `$eq` | 等于该值 | `{"plan": {"$eq": "pro"}}` |
| `$ne` | 不等于 | `{"plan": {"$ne": "free"}}` |
| `$gt` | 大于 | `{"mrr": {"$gt": 100}}` |
| `$ge` | 大于或等于 | `{"mrr": {"$ge": 100}}` |
| `$lt` | 小于 | `{"mrr": {"$lt": 100}}` |
| `$le` | 小于或等于 | `{"mrr": {"$le": 100}}` |
| `$in` | 是其中一个值 | `{"plan": {"$in": ["pro", "team"]}}` |
| `$nin` | 不是其中任何一个值 | `{"country": {"$nin": ["US", "GB"]}}` |
| `$contains` | 是包含该值的文本 | `{"email": {"$contains": "@example.com"}}` |

> **请从这张表中复制运算符名称。** 其他数据库中常见的 `$gte` 和 `$lte` 在 InceptionDB 中不存在：请使用 `$ge` 和 `$le`。目前未知运算符不会报错：过滤器不匹配任何文档，响应为 `204 No Content`。因此 `{"mrr": {"$gte": 100}}` 不返回任何客户，而 `{"mrr": {"$ge": 100}}` 返回 Ada 和 Grace。

以 ISO 格式文本存储的日期（`"2026-10-28T09:30:00Z"`）按时间顺序比较：`{"created_at": {"$gt": "2026-01-01"}}`。

## 组合条件

多个字段必须全部满足。要匹配多个条件中的**任意一个**，使用 `$or`：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"$or": [{"plan": "team"}, {"country": "US"}]}, "limit": 10}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

`$and` 同样接收一个列表，可用于对同一字段应用两个条件。

## 嵌套字段

用点号访问对象内部的字段：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"address.city": "London"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
```

## 数组

对数组字段的条件，只要**任意一个元素**满足即匹配。带有 `vip` 标签的客户：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"tags": "vip"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## 分页

组合使用 `skip` 和 `limit`。每页 25 个文档的第二页：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {}, "skip": 25, "limit": 25}'
```

没有索引时，文档不按任何特定顺序返回。要得到有序结果，或让大集合上的查询更快，请使用 [索引](/zh/docs/inceptiondb/v2/indexes/)。

## 查询返回 204，但你预期有文档

`204 No Content` 表示没有文档匹配。它的响应体为空：解码之前请先检查状态码。常见原因：

- **运算符名称不存在**，例如用了 `$gte` 或 `$lte`，而不是 `$ge` 和 `$le`。
- **字段名拼写不同。** 名称区分大小写：`Country` 不是 `country`。
- **值的类型不同。** `{"mrr": "120"}` 查找的是文本 `"120"`，而不是数字 `120`。
- **`skip` 超过了最后一个匹配**，例如集合中的文档少于一页时请求第二页。

## 在控制台中试用

[控制台](https://inceptiondb.hola.cloud/) 会在你输入时运行过滤器，提示你数据中的字段名，并根据数据给出示例过滤器。**Schema** 标签页列出每个字段的类型和最常见的值。
