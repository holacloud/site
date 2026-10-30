# 索引管理

创建、列出、获取和删除集合的索引。都需要 API key。各类型的适用场景见 [索引](/zh/docs/inceptiondb/v2/indexes/)。

## 创建索引

`POST …/collections/{collection}:createIndex`

| 请求体字段 | 类型 | 含义 |
|---|---|---|
| `name` | 全部 | 索引名称。 |
| `type` | 全部 | `map`、`btree` 或 `fts`。 |
| `field` | `map`、`fts` | 被索引的字段。 |
| `fields` | `btree` | 被索引的字段，按顺序。`-字段` 表示降序。 |
| `unique` | `btree` | 拒绝重复值。`map` 索引始终唯一。 |
| `sparse` | `map`、`btree` | 没有该字段的文档不会进入索引，但始终会被保存。 |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email", "type": "map", "field": "email"}'
```

响应 `201 Created`：

```json
{"field":"email","name":"by_email","sparse":false,"type":"map"}
```

## 列出索引

`POST …/collections/{collection}:listIndexes`

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:listIndexes" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"field":"email","name":"by_email","sparse":false,"type":"map"}]
```

## 获取索引

`POST …/collections/{collection}:getIndex`

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:getIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email"}'
```

```json
{"field":"email","name":"by_email","sparse":false,"type":"map"}
```

## 删除索引

`POST …/collections/{collection}:dropIndex`

文档不受影响。响应 `204 No Content`。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email"}'
```
