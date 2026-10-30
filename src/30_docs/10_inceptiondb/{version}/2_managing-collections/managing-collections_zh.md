# 集合

集合是数据库中一组有名称的 JSON 文档，例如 `customers` 或 `orders`。同一集合中的文档可以有不同的字段：无需声明模式。

**开始之前：** 你只需要 [快速开始](/zh/docs/inceptiondb/v2/getting-started/) 中的环境变量。本指南可以单独使用：它会创建一个名为 `collections-demo` 的集合，并在最后删除它。它不会改动 `customers`。

## 创建集合

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "collections-demo"}'
```

响应（`201 Created`）：

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"collections-demo","total":0}
```

| 字段 | 含义 |
|---|---|
| `name` | 集合名称，用于它的所有 URL。 |
| `total` | 文档数量。 |
| `indexes` | 索引数量。 |
| `defaults` | 添加到缺少这些字段的文档中的值（见下文）。 |

名称区分大小写。集合已存在时返回 `500`，消息为 ``collection `collections-demo` already exists``：如果你之前运行过本指南，请执行最后一步 *删除集合*，然后重新开始。

向不存在的集合插入会创建该集合，并使用默认值。如果 `:insert` 中的集合名称有拼写错误，文档会保存到一个新集合中，而不会报错。

## 列出集合

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"collections-demo","total":0},{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":2}]
```

如果你完成了 [快速开始](/zh/docs/inceptiondb/v2/getting-started/)，列表中还会出现 `customers`。

## 获取单个集合

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/collections-demo" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"collections-demo","total":0}
```

集合不存在时返回 `404 Not Found`。

## 默认值

默认值会填充插入文档中缺少的字段。值为 `null` 的字段也会使用默认值。除了固定值，还有三种生成值：

| 值 | 生成 |
|---|---|
| `"uuid()"` | 一个随机唯一 ID，例如 `"35ec7d8f-96f4-4cdf-a09e-463c720e7c8c"`。 |
| `"auto()"` | 一个递增数字：1、2、3……永不重复，但可能跳号。 |
| `"unixnano()"` | 插入时间，单位为自 1970 年起的纳秒。 |

为每个新文档设置 ID、`free` 套餐和创建时间：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/collections-demo:setDefaults" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"id": "uuid()", "plan": "free", "created": "unixnano()"}'
```

```json
{"created":"unixnano()","id":"uuid()","plan":"free"}
```

现在插入一个只有姓名和邮箱的文档：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/collections-demo:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/x-ndjson" \
  -d '{"name": "Margaret Hamilton", "email": "margaret@example.com"}'
```

```jsonl
{"created":1793164963789350717,"email":"margaret@example.com","id":"35ec7d8f-96f4-4cdf-a09e-463c720e7c8c","name":"Margaret Hamilton","plan":"free"}
```

默认值只作用于之后插入的文档，已有文档不会改变。`setDefaults` 会把你发送的字段添加到当前默认值中，并替换已存在的同名字段。要删除某个默认值，把它设为 `null`：`{"plan": null}`。

## 大小

文档占用的内存、磁盘大小，以及每个索引以 `index.<名称>` 表示的内存，单位为字节：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/collections-demo:size" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"disk":265,"memory":203}
```

## 删除集合

这会删除集合及其所有文档和索引，**无法撤销**：运行前请检查 URL 中的名称。删除本指南的集合 `collections-demo`：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/collections-demo:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

响应为 `204 No Content`。删除不存在的集合会返回 `500`，消息为 ``collection `collections-demo` not found``。
