
# 处理数据

本文档解释了如何在集合中插入、列出、修改和删除文档。还包括如何在列出时过滤数据。

## 插入文档

### 插入单个文档

要在集合中插入单个文档，请使用以下HTTP请求：

```http
POST /databases/{database}/collections/{collection}/documents
Content-Type: application/json
X-API-Key: your-api-key

{
  "document": {
    "key1": "value1",
    "key2": "value2"
  }
}
```

### 插入多个文档

要在集合中插入多个文档，请使用以下HTTP请求：

```http
POST /databases/{database}/collections/{collection}/documents/bulk
Content-Type: application/json
X-API-Key: your-api-key

{
  "documents": [
    {
      "key1": "value1",
      "key2": "value2"
    },
    {
      "key1": "value3",
      "key2": "value4"
    }
  ]
}
```

## 列出文档

### 列出所有文档（全扫描）

要列出集合中的所有文档，请使用以下HTTP请求：

```http
GET /databases/{database}/collections/{collection}/documents
X-API-Key: your-api-key
```

### 通过唯一索引列出文档

要使用唯一索引列出文档，请使用以下HTTP请求：

```http
GET /databases/{database}/collections/{collection}/documents?filter={"key":"value"}
X-API-Key: your-api-key
```

### 列出具有特定限制的文档

要列出具有特定限制的文档，请使用以下HTTP请求：

```http
GET /databases/{database}/collections/{collection}/documents?limit=10
X-API-Key: your-api-key
```

## 修改文档

### 通过全扫描修改文档

要通过全扫描修改文档，请使用以下HTTP请求：

```http
PATCH /databases/{database}/collections/{collection}/documents
Content-Type: application/json
X-API-Key: your-api-key

{
  "filter": {"key": "value"},
  "update": {"$set": {"key": "new_value"}}
}
```

### 通过索引修改文档

要使用索引修改文档，请使用以下HTTP请求：

```http
PATCH /databases/{database}/collections/{collection}/documents?filter={"key":"value"}
Content-Type: application/json
X-API-Key: your-api-key

{
  "update": {"$set": {"key": "new_value"}}
}
```

## 删除文档

### 通过索引删除文档

要使用索引删除文档，请使用以下HTTP请求：

```http
DELETE /databases/{database}/collections/{collection}/documents?filter={"key":"value"}
X-API-Key: your-api-key
```

### 通过全扫描删除文档

要通过全扫描删除文档，请使用以下HTTP请求：

```http
DELETE /databases/{database}/collections/{collection}/documents
Content-Type: application/json
X-API-Key: your-api-key

{
  "filter": {"key": "value"}
}
```
