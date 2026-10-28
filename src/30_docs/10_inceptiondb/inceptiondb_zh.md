# InceptionDB

InceptionDB 是一个通过 HTTP 使用的 JSON 文档数据库。无需安装驱动，也无需定义模式：发送 JSON，返回 JSON，任何能发起 HTTP 请求的语言都能使用。

## 基本概念

| 概念 | 说明 |
|---|---|
| **数据库** | 你的空间。包含集合，并拥有自己的 API key。在[控制台](https://inceptiondb.hola.cloud/)中创建。 |
| **集合** | 一组有名称的文档，例如 `customers` 或 `orders`。同一集合中的文档可以有不同的字段。 |
| **文档** | 一个 JSON 对象。每个文档都有 `id`；如果你不提供，InceptionDB 会自动生成。 |
| **索引** | 可选。让按字段查找和范围查询更快，并可保证值唯一。 |
| **API key** | 一个密钥和一个 secret，应用通过 `Api-Key` 和 `Api-Secret` 请求头发送，以访问某个数据库的数据。 |

所有请求都发往 `https://inceptiondb.hola.cloud`，示例从三个环境变量读取你的凭据，因此可以直接复制粘贴：

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

## 从哪里开始

- **[快速开始](/zh/docs/inceptiondb/v2/getting-started/)**：五分钟内完成第一次查询。
- **[集合](/zh/docs/inceptiondb/v2/managing-collections/)**：创建集合、设置默认值、查看大小。
- **[处理数据](/zh/docs/inceptiondb/v2/working-with-data/)**：插入、读取、更新和删除文档。
- **[查询](/zh/docs/inceptiondb/v2/querying/)**：过滤器、运算符、嵌套字段、数组和分页。
- **[索引](/zh/docs/inceptiondb/v2/indexes/)**：让查询更快并保证值唯一。
- **[API key](/zh/docs/inceptiondb/v2/api-keys/)**：创建数据库和密钥，并保证其安全。
- **[API 参考](/zh/docs/inceptiondb/v2/api-reference/)**：每个端点的请求和响应。

## 适用场景

- **应用数据**：用户、设置、购物车和订单，任何结构会随时间变化的 JSON。
- **内容**：文章、页面和字段因条目而异的产品目录。
- **事件与遥测**：高速写入的设备读数和产品分析数据。
- **原型**：几分钟内开始存储数据，等查询清晰后再添加索引。
