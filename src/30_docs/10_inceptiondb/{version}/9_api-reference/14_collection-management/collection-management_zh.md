# 集合管理

获取、配置、测量和删除集合。都需要 API key。

## 获取集合

`GET /v1/databases/{databaseId}/collections/{collection}`

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":3}
```

## 设置默认值

`POST …/collections/{collection}:setDefaults`

你发送的字段会添加到当前默认值中，并替换同名字段；值为 `null` 的字段会删除该默认值。除了固定值，`"uuid()"` 生成随机 id，`"auto()"` 生成递增数字，`"unixnano()"` 生成以纳秒为单位的插入时间。当文档没有该字段或该字段为 `null` 时，默认值生效。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:setDefaults" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"id": "uuid()", "plan": "free", "created": "unixnano()"}'
```

响应 `200 OK`，包含当前生效的全部默认值：

```json
{"created":"unixnano()","id":"uuid()","plan":"free"}
```

## 大小

`POST …/collections/{collection}:size`

文档占用的内存字节数（`memory`）、每个索引占用的内存（`index.<名称>`）以及磁盘大小（`disk`）。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:size" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"disk":630,"index.by_email":260,"memory":646}
```

## 删除集合

`POST …/collections/{collection}:dropCollection`

删除集合及其文档和索引，无法撤销。响应 `204 No Content`。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```
