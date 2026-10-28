# 集合

集合是数据库中一组有名称的 JSON 文档，例如 `customers` 或 `orders`。同一集合中的文档可以有不同的字段：无需声明模式。

示例使用 [快速开始](/zh/docs/inceptiondb/v2/getting-started/) 中的环境变量。

## 创建集合

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

| 字段 | 含义 |
|---|---|
| `name` | 集合名称，用于它的所有 URL。 |
| `total` | 文档数量。 |
| `indexes` | 索引数量。 |
| `defaults` | 添加到缺少这些字段的文档中的值（见下文）。 |

## 列出集合

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}]
```

## 获取单个集合

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

## 默认值

默认值会填充插入文档中缺少的字段。值为 `null` 的字段也会使用默认值。除了固定值，还有三种生成值：

| 值 | 生成 |
|---|---|
| `"uuid()"` | 一个随机唯一 ID，例如 `"35ec7d8f-96f4-4cdf-a09e-463c720e7c8c"`。 |
| `"auto()"` | 一个递增数字：1、2、3……永不重复，但可能跳号。 |
| `"unixnano()"` | 插入时间，单位为自 1970 年起的纳秒。 |

为每个新客户设置 ID、`free` 套餐和创建时间：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:setDefaults" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"id": "uuid()", "plan": "free", "created": "unixnano()"}'
```

```json
{"created":"unixnano()","id":"uuid()","plan":"free"}
```

现在插入一个只有姓名和邮箱的客户：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "Margaret Hamilton", "email": "margaret@example.com"}'
```

```jsonl
{"created":1793164963789350717,"email":"margaret@example.com","id":"35ec7d8f-96f4-4cdf-a09e-463c720e7c8c","name":"Margaret Hamilton","plan":"free"}
```

默认值只作用于之后插入的文档，已有文档不会改变。`setDefaults` 会把你发送的字段添加到当前默认值中，并替换已存在的同名字段。要删除某个默认值，把它设为 `null`：`{"plan": null}`。

## 大小

文档占用的内存、磁盘大小，以及每个索引以 `index.<名称>` 表示的内存，单位为字节：

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:size" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"disk":265,"memory":203}
```

## 删除集合

这会删除集合及其所有文档和索引，无法撤销。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

响应为 `204 No Content`。
