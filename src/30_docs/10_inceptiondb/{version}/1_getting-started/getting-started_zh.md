# 快速开始

五分钟内，你将创建一个集合，保存一些客户，查询它们，修改一个并删除另一个。每条命令都可以直接复制粘贴。

**开始之前：** 你需要一个 hola.cloud 账户。本指南会创建 `customers` 集合，并在其中留下两个客户：Ada 和 Grace。[处理数据](/zh/docs/inceptiondb/v2/working-with-data/)、[查询](/zh/docs/inceptiondb/v2/querying/) 和 [索引](/zh/docs/inceptiondb/v2/indexes/) 都从这里开始。

## 1. 获取凭据

1. 打开 [InceptionDB 控制台](https://inceptiondb.hola.cloud/) 并登录。
2. 创建一个数据库（例如 `my-app`），或选择一个已有的数据库。
3. 打开 **API keys & settings** 并创建一个 API key。secret **只显示一次**：现在就复制它。

控制台会给出可以直接粘贴到终端的三个值。把它们导出为环境变量，这样它们永远不会出现在你的代码中：

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

列出数据库中的集合来检查它们：

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

- `200 OK` 表示凭据有效。新数据库中的列表为空：`[]`。
- `403 Forbidden` 表示密钥、secret 或数据库 ID 不正确。请检查这三个变量。

## 用其他语言运行示例

每个示例都有 curl、HTTP、Go、PHP、Python、Node.js 和 Java 标签页。它们都读取相同的三个环境变量，并打印响应的状态码和响应体。把代码保存到新文件中并运行：

| 语言 | 需要 | 运行 |
|---|---|---|
| Go | Go 1.21 或更高版本。只使用标准库。 | `go run main.go` |
| PHP | PHP 8 及 `curl` 扩展。 | `php example.php` |
| Python | Python 3.8 或更高版本，以及 `pip install requests`。 | `python example.py` |
| Node.js | Node.js 18 或更高版本。将文件保存为 `.mjs`。 | `node example.mjs` |
| Java | Java 17 或更高版本。只使用标准库。 | `java Main.java` |

没有浏览器版本：secret 可以完全访问你的数据，所以它必须留在服务器上。见 [API key](/zh/docs/inceptiondb/v2/api-keys/)。

## 2. 创建集合

集合用于组织文档。创建一个名为 `customers` 的集合：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "customers"}'
```

响应（`201 Created`）：

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

`defaults` 表示没有 `id` 的文档会获得一个随机 id。

如果集合已存在（你之前运行过本指南），响应为 `500`，消息为 ``collection `customers` already exists``。删除它，然后从这一步重新开始。这会删除其中的文档且无法撤销，所以请先确认名称：

<!-- docs-check: skip -->
```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

## 3. 插入文档

插入使用 **JSON Lines**（JSONL）：每行一个 JSON 对象。它不是 JSON 数组，所以不要用 `[...]` 包裹文档，也不要用逗号分隔。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary '{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}'
```

响应（`201 Created`）包含已保存的文档，每行一个：

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## 4. 查询

查找来自英国的客户。`limit` 是你想要返回的文档数量：**不设置时只返回一个**。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"country": "GB"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

过滤器还可以比较数字、查看数组和嵌套对象内部，以及组合条件。每月付费 120 或以上的客户：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"mrr": {"$ge": 120}}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

> **运算符名称：** InceptionDB 使用 `$ge`（大于或等于）和 `$le`（小于或等于）。`$gte` 和 `$lte` 无效：使用它们的过滤器不会匹配任何文档，响应为 `204 No Content`，不会报错。

[查询](/zh/docs/inceptiondb/v2/querying/) 中列出了所有运算符。

## 5. 读取一个文档

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/ada" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

文档在 `document` 中：

```json
{"id":"ada","document":{"address":{"city":"London","zip":"N1"},"country":"GB","email":"ada@example.com","id":"ada","mrr":120,"name":"Ada Lovelace","plan":"pro","tags":["vip","beta"]},"source":{"type":"fullscan"}}
```

## 6. 更新文档

Alan 换成了 `pro` 套餐。patch 只修改你发送的字段：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}, "patch": {"plan": "pro", "mrr": 120}}'
```

```jsonl
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"pro","country":"GB","mrr":120,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

## 7. 删除文档

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}}'
```

响应包含被删除的文档：

```jsonl
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"pro","country":"GB","mrr":120,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

现在 `customers` 中有两个文档：Ada 和 Grace。请保留它：后面的指南会用到。

## 下一步

- [处理数据](/zh/docs/inceptiondb/v2/working-with-data/)：批量插入，更新和删除多个文档。
- [查询](/zh/docs/inceptiondb/v2/querying/)：所有过滤运算符、嵌套字段、数组和分页。
- [索引](/zh/docs/inceptiondb/v2/indexes/)：按邮箱即时查找、按金额范围查询、按词搜索。
- [集合](/zh/docs/inceptiondb/v2/managing-collections/)：默认值、大小和删除集合。
- [控制台](https://inceptiondb.hola.cloud/) 会显示你在其中执行的每个操作的代码。[API key](/zh/docs/inceptiondb/v2/api-keys/) 中说明了在哪里查看。
