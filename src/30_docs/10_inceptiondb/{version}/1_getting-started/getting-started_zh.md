# 快速开始

五分钟内，你将创建一个集合，保存一些客户，查询它们，修改一个并删除另一个。所有命令都可以直接复制粘贴。

## 1. 获取凭据

1. 打开 [InceptionDB 控制台](https://inceptiondb.hola.cloud/) 并登录。
2. 创建一个数据库（例如 `my-app`），或选择已有的数据库。
3. 打开 **API keys & settings** 并创建一个 API key。secret **只显示一次**：请立即复制。

控制台会给出可以直接粘贴到终端的三个值。把它们导出为环境变量，这样它们就不会出现在你的代码中：

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

本文档中的所有示例都会读取这些变量，Go、Python、Node.js、PHP 和 Java 版本也一样。

## 2. 创建集合

集合用于组织文档。创建一个名为 `customers` 的集合：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "customers"}'
```

响应（`201 Created`）：

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

`defaults` 表示没有 `id` 的文档会获得一个随机 id。

## 3. 插入文档

每行发送一个 JSON 文档（这种格式称为 JSON Lines）：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  --data-binary '{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}'
```

响应（`201 Created`）返回已保存的文档，每行一个：

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## 4. 查询

查找来自英国的客户。`limit` 是你想返回的文档数量：**不设置时只返回一个**。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"country": "GB"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}
```

过滤器还可以比较数字、查看数组和嵌套对象内部，以及组合条件。每月付费超过 100 的客户：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"mrr": {"$gt": 100}}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

所有运算符见 [查询](/zh/docs/inceptiondb/v2/querying/)。

## 5. 读取一个文档

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/ada" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

文档位于 `document` 中：

```json
{"id":"ada","document":{"address":{"city":"London","zip":"N1"},"country":"GB","email":"ada@example.com","id":"ada","mrr":120,"name":"Ada Lovelace","plan":"pro","tags":["vip","beta"]},"source":{"type":"fullscan"}}
```

## 6. 更新文档

Alan 升级到 `pro` 套餐。patch 只修改你发送的字段：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
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
  -d '{"filter": {"id": "grace"}}'
```

响应返回被删除的文档：

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## 下一步

- [查询](/zh/docs/inceptiondb/v2/querying/)：所有过滤运算符、分页和导出。
- [索引](/zh/docs/inceptiondb/v2/indexes/)：按邮箱即时查找、按金额范围查询、按词搜索。
- [处理数据](/zh/docs/inceptiondb/v2/working-with-data/)：批量插入、更新多个文档以及错误处理。
- [控制台](https://inceptiondb.hola.cloud/) 会显示你在其中执行的每个操作的代码：点击屏幕右侧的小球即可查看。
