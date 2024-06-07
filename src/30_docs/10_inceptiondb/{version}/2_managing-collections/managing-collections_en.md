
# Managing Collections

---

## Creating a Collection

To create a collection, you can use the following curl command:

```sh
curl -X POST "https://example.com/v1/databases/719e9421-e6e9-42b6-b9b7-b580e532b9d5/collections" \
-H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
-H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
-d '{
    "name": "my-collection"
}'
```

The HTTP request for creating a collection would look like this:

```http
POST /v1/databases/719e9421-e6e9-42b6-b9b7-b580e532b9d5/collections HTTP/1.1
Host: example.com
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
    "name": "my-collection"
}
```

## Deleting a Collection

To delete a collection, use the following curl command:

```sh
curl -X POST "https://example.com/v1/databases/dc992d92-c146-448b-a17d-e94614aff740/collections/my-collection:dropCollection" \
-H "Api-Key: 271ba1e6-87b5-4d0f-84a0-a3d1178d1356" \
-H "Api-Secret: 5346a7cf-6743-49eb-921b-6e977cf11e36"
```

The HTTP request for deleting a collection would look like this:

```http
POST /v1/databases/dc992d92-c146-448b-a17d-e94614aff740/collections/my-collection:dropCollection HTTP/1.1
Host: example.com
Api-Key: 271ba1e6-87b5-4d0f-84a0-a3d1178d1356
Api-Secret: 5346a7cf-6743-49eb-921b-6e977cf11e36
```

## Creating a Collection Implicitly

### When Inserting Documents

Collections are created implicitly when you insert documents into them. For example:

#### Insert One Document

```sh
curl -X POST "https://example.com/v1/databases/436a4499-e9f0-4aaf-9f82-fe2a2e22c3f0/collections/my-collection:insert" \
-H "Api-Key: ea34c657-4125-4527-8cd1-630b33f50b42" \
-H "Api-Secret: a30baf65-c9be-4231-92bc-6f04442142c9" \
-d '{
    "address": "Elm Street 11",
    "id": "my-id",
    "name": "Fulanez"
}'
```

The HTTP request would be:

```http
POST /v1/databases/436a4499-e9f0-4aaf-9f82-fe2a2e22c3f0/collections/my-collection:insert HTTP/1.1
Host: example.com
Api-Key: ea34c657-4125-4527-8cd1-630b33f50b42
Api-Secret: a30baf65-c9be-4231-92bc-6f04442142c9
Content-Type: application/json

{
    "address": "Elm Street 11",
    "id": "my-id",
    "name": "Fulanez"
}
```

#### Insert Multiple Documents

```sh
curl -X POST "https://example.com/v1/databases/7b7d79de-d7c7-47b9-8662-e124be300d3d/collections/my-collection:insert" \
-H "Api-Key: e609f9ce-9f15-4649-b57c-5e9709e5bcca" \
-H "Api-Secret: 19e7a8e9-4cc6-4814-ad3e-c33af518086a" \
-d '{
    "id":"1",
    "name":"Alfonso"
}
{
    "id":"2",
    "name":"Gerardo"
}
{
    "id":"3",
    "name":"Alfonso"
}'
```

The HTTP request would be:

```http
POST /v1/databases/7b7d79de-d7c7-47b9-8662-e124be300d3d/collections/my-collection:insert HTTP/1.1
Host: example.com
Api-Key: e609f9ce-9f15-4649-b57c-5e9709e5bcca
Api-Secret: 19e7a8e9-4cc6-4814-ad3e-c33af518086a
Content-Type: application/json

{
    "id":"1",
    "name":"Alfonso"
}
{
    "id":"2",
    "name":"Gerardo"
}
{
    "id":"3",
    "name":"Alfonso"
}
```

### When Creating an Index

Creating an index can also implicitly create a collection. For example:

```sh
curl -X POST "https://example.com/v1/databases/c4e4ac96-5001-4465-b22e-57e971ecc01f/collections/my-collection:createIndex" \
-H "Api-Key: 4469ae2e-f4d2-4c58-ad03-d34b97957f3a" \
-H "Api-Secret: aa7d2784-50b7-4846-be77-d7609034ab22" \
-d '{
    "field": "id",
    "name": "my-index-id",
    "sparse": true,
    "type": "map"
}'
```

The HTTP request would be:

```http
POST /v1/databases/c4e4ac96-5001-4465-b22e-57e971ecc01f/collections/my-collection:createIndex HTTP/1.1
Host: example.com
Api-Key: 4469ae2e-f4d2-4c58-ad03-d34b97957f3a
Api-Secret: aa7d2784-50b7-4846-be77-d7609034ab22
Content-Type: application/json

{
    "field": "id",
    "name": "my-index-id",
    "sparse": true,
    "type": "map"
}
```

---

This document explains how to manage collections, including creating, deleting, and the various implicit methods of creation. You can download this document in Markdown format using the link below:

[Download Managing Collections Document](sandbox:/mnt/data/Managing_Collections.md)
