
# InceptionDB快速入门指南

本指南将帮助您通过创建集合、插入三个元素和查询过滤的元素来开始使用InceptionDB。

## 第一步：创建集合

要在InceptionDB中创建集合，您需要定义集合的名称。以下是使用curl在bash中执行此操作的示例：

```bash
curl -X POST "https://example.com/v1/databases/你的数据库ID/collections" \
-H "Api-Key: 你的API密钥" \
-H "Api-Secret: 你的API秘密" \
-d '{
      "name": "我的集合"
    }'
```

预期响应

```json
{
    "defaults": {
        "id": "uuid()"
    },
    "indexes": 0,
    "name": "我的集合",
    "total": 0
}
```

## 第二步：插入元素

创建集合后，您可以向其中插入元素。以下是在新创建的集合中插入三个元素的方法：

```bash
curl -X POST "https://example.com/v1/databases/你的数据库ID/collections/我的集合/documents" \
-H "Api-Key: 你的API密钥" \
-H "Api-Secret: 你的API秘密" \
-d '{
      "name": "元素1",
      "value": 100
    }
    {
      "name": "元素2",
      "value": 200
    }
    {
      "name": "元素3",
      "value": 300
    }'
```

## 第三步：列出过滤的元素

要在集合中按特定过滤条件列出元素，可以使用以下curl请求：

```bash
curl -X POST "https://example.com/v1/databases/你的数据库ID/collections/我的集合/find" \
-H "Api-Key: 你的API密钥" \
-H "Api-Secret: 你的API秘密" \
-d '{
      "filter": {
        "value": { "$gte": 200 }
      }
    }'
```

预期响应

```json
{
    "documents": [
        {
            "id": "文档ID-2",
            "name": "元素2",
            "value": 200
        },
        {
            "id": "文档ID-3",
            "name": "元素3",
            "value": 300
        }
    ]
}
```

## 总结

通过按照这些步骤操作，您已经在InceptionDB中创建了一个集合，插入了三个元素，并通过过滤条件查询了元素。这为您开始使用InceptionDB并利用其高效管理数据的能力奠定了坚实的基础。
