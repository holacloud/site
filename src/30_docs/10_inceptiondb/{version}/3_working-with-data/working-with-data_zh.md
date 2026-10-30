# 处理数据

插入、读取、更新和删除文档。

**开始之前：** 本指南使用 [快速开始](/zh/docs/inceptiondb/v2/getting-started/) 留下的 `customers` 集合（包含 Ada 和 Grace）及其环境变量。它会添加两个标记为 `"trial": true` 的临时客户，并在最后删除它们，因此 `customers` 会恢复原样。

## 插入文档

插入使用 **JSON Lines**（JSONL）：每行一个 JSON 对象。一个请求可以包含一个文档，也可以包含数百万个。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary '{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0,"trial":true}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true}'
```

响应（`201 Created`）返回已保存的文档，其中已应用集合的默认值：

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0,"trial":true}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true}
```

JSON 数组**不是** JSONL。这个请求体会以 `400 Bad Request` 被拒绝，消息为 `document must be a JSON object`：

```text
[{"id":"linus"},{"id":"barbara"}]
```

同样的两个文档，使用 JSONL：

```text
{"id":"linus"}
{"id":"barbara"}
```

注意：

- 没有 `id` 的文档会从集合的 [默认值](/zh/docs/inceptiondb/v2/managing-collections/) 获得一个（除非你修改，否则是随机 UUID）。
- InceptionDB 不检查 id 是否唯一：用同一个 `id` 插入两次会保存两个文档。要避免这种情况，请在 `id` 上创建 [`map` 索引](/zh/docs/inceptiondb/v2/indexes/)。之后重复的 id 会返回 `409 Conflict`，按 id 读取也会立即完成。
- 文档按顺序保存，插入会**在第一个被拒绝的文档处停止**。之前的文档仍会保存。
- 要加载文件，使用 `--data-binary @customers.jsonl`。`-d` 会去掉换行符。
- 每个文档必须是没有重复字段名的 JSON 对象，否则返回 `400 Bad Request`。
- 如果 [唯一索引](/zh/docs/inceptiondb/v2/indexes/) 中已存在该值，返回 `409 Conflict`。
- 向不存在的集合插入会创建该集合。请检查 URL 中的集合名称。

## 按 id 读取文档

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/linus" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"id":"linus","document":{"country":"FI","id":"linus","mrr":0,"name":"Linus Torvalds","plan":"free","trial":true},"source":{"type":"fullscan"}}
```

`source` 说明文档是如何找到的：`index` 表示由 `id` 上的索引响应，`fullscan` 表示遍历了集合。文档不存在时返回 `404 Not Found`。

要一次读取多个文档，请使用 [查询](/zh/docs/inceptiondb/v2/querying/)。

## 更新文档

patch 只修改你发送的字段（[JSON merge patch](https://www.rfc-editor.org/rfc/rfc7386)）：新字段会被添加，已有字段会被替换，`null` 会删除字段。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "linus"}, "patch": {"plan": "pro", "mrr": 49, "country": null}}'
```

响应返回修改后的文档：

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49,"trial":true}
```

`filter` 用于选择文档，语法与 [查询](/zh/docs/inceptiondb/v2/querying/) 相同。与查询一样，patch **默认只修改一个文档，除非设置 `limit`**。给所有试用客户加上 `priority` 标记：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"trial": true}, "patch": {"priority": true}, "limit": -1}'
```

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49,"trial":true,"priority":true}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true,"priority":true}
```

`"limit": -1` 表示所有匹配的文档。没有任何匹配时，返回 `204 No Content`。

patch 也可以修改文档的 `id`。如果集合在 `id` 上有 `map` 索引，改成已存在的 id 的 patch 不会被应用。

## 删除文档

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "linus"}}'
```

响应返回被删除的文档：

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49,"trial":true,"priority":true}
```

默认只删除一个文档，除非设置 `limit`。删除所有试用客户，这同时也清理了本指南创建的数据：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"trial": true}, "limit": -1}'
```

```jsonl
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true,"priority":true}
```

被删除的文档无法恢复。请先用 `:find` 运行相同的过滤器，确认会匹配哪些文档。

## 导出集合

`:find` 加上 `"limit": -1` 会以 JSONL 返回所有文档，可以直接保存到文件或导入到别处：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"limit": -1}' \
  -o customers.jsonl
```

`customers.jsonl` 包含 Ada 和 Grace，每行一个：集合已恢复到本指南之前的状态。
