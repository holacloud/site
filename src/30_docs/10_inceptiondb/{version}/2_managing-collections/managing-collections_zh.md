
# 管理集合

## 创建集合

要在数据库中创建集合，使用带有必要API凭据的`POST`命令到相应的URL。以下是使用`curl`进行操作的示例。

```sh
curl -X POST "https://example.com/v1/databases/{database_id}/collections" \
-H "Api-Key: {api_key}" \
-H "Api-Secret: {api_secret}" \
-d '{
    "name": "my-collection"
}'
```

HTTP请求和响应示例：

```http
POST /v1/databases/719e9421-e6e9-42b6-b9b7-b580e532b9d5/collections HTTP/1.1
Host: example.com
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf

{
    "name": "my-collection"
}

HTTP/1.1 201 Created
Content-Length: 74
Content-Type: text/plain; charset=utf-8
Date: Mon, 15 Aug 2022 02:08:13 GMT

{
    "defaults": {
        "id": "uuid()"
    },
    "indexes": 0,
    "name": "my-collection",
    "total": 0
}
```

## 列出集合

要列出数据库中的所有集合，使用带有必要API凭据的`GET`命令到相应的URL。以下是使用`curl`进行操作的示例。

```sh
curl "https://example.com/v1/databases/{database_id}/collections" \
-H "Api-Key: {api_key}" \
-H "Api-Secret: {api_secret}"
```

HTTP请求和响应示例：

```http
GET /v1/databases/13f8a9fa-7e46-4732-9545-658ea2b87112/collections HTTP/1.1
Host: example.com
Api-Secret: 07510cc8-40cc-41f1-825b-449ae72895b2
Api-Key: 8f8617b1-320b-47bf-b534-6d701e03b2b0

HTTP/1.1 200 OK
Content-Length: 76
Content-Type: text/plain; charset=utf-8
Date: Mon, 15 Aug 2022 02:08:13 GMT

[
    {
        "defaults": {
            "id": "uuid()"
        },
        "indexes": 0,
        "name": "my-collection",
        "total": 0
    }
]
```

## 删除集合

要删除集合，使用带有必要API凭据的`POST`命令到特定URL，指定要删除的集合。以下是使用`curl`进行操作的示例。

```sh
curl -X POST "https://example.com/v1/databases/{database_id}/collections/my-collection:dropCollection" \
-H "Api-Key: {api_key}" \
-H "Api-Secret: {api_secret}"
```

HTTP请求和响应示例：

```http
POST /v1/databases/dc992d92-c146-448b-a17d-e94614aff740/collections/my-collection:dropCollection HTTP/1.1
Host: example.com
Api-Key: 271ba1e6-87b5-4d0f-84a0-a3d1178d1356
Api-Secret: 5346a7cf-6743-49eb-921b-6e977cf11e36

HTTP/1.1 200 OK
Content-Length: 0
Date: Mon, 15 Aug 2022 02:08:13 GMT
```

## 隐式创建集合的方法

### 插入时

当在不存在的集合中插入文档时，可以隐式创建集合。

### 创建索引时

如果尝试在不存在的集合中创建索引，集合会自动创建。

### 使用`setDefaults`时

在不存在的集合中使用`setDefaults`方法时，也会隐式创建集合。

本文档提供了有关如何在数据库中管理集合的详细指南，包括集合的创建、列出和删除以及隐式创建集合的方法。
