# Manejando colecciones

## Crear una Colección

Para crear una colección en la base de datos, se utiliza un comando `POST` a la URL correspondiente con las credenciales API necesarias. A continuación, se presenta un ejemplo de cómo se hace esto usando `curl`.

```sh
curl -X POST "https://example.com/v1/databases/{database_id}/collections" \
-H "Api-Key: {api_key}" \
-H "Api-Secret: {api_secret}" \
-d '{
    "name": "my-collection"
}'
```

Ejemplo de solicitud y respuesta HTTP:

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

## Listar Colecciones

Para listar todas las colecciones en una base de datos, se utiliza un comando `GET` a la URL correspondiente con las credenciales API necesarias. A continuación, se presenta un ejemplo de cómo se hace esto usando `curl`.

```sh
curl "https://example.com/v1/databases/{database_id}/collections" \
-H "Api-Key: {api_key}" \
-H "Api-Secret: {api_secret}"
```

Ejemplo de solicitud y respuesta HTTP:

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

## Eliminar una Colección

Para eliminar una colección, se utiliza un comando `POST` a la URL específica con las credenciales API necesarias, indicando la colección que se desea eliminar. A continuación, se presenta un ejemplo de cómo se hace esto usando `curl`.

```sh
curl -X POST "https://example.com/v1/databases/{database_id}/collections/my-collection:dropCollection" \
-H "Api-Key: {api_key}" \
-H "Api-Secret: {api_secret}"
```

Ejemplo de solicitud y respuesta HTTP:

```http
POST /v1/databases/dc992d92-c146-448b-a17d-e94614aff740/collections/my-collection:dropCollection HTTP/1.1
Host: example.com
Api-Key: 271ba1e6-87b5-4d0f-84a0-a3d1178d1356
Api-Secret: 5346a7cf-6743-49eb-921b-6e977cf11e36

HTTP/1.1 200 OK
Content-Length: 0
Date: Mon, 15 Aug 2022 02:08:13 GMT
```

## Formas Implícitas de Crear una Colección

### Al Insertar

Una colección puede ser creada implícitamente cuando se inserta un documento en una colección que no existe previamente.

### Al Crear un Índice

Si se intenta crear un índice en una colección que no existe, la colección se crea automáticamente.

### Al Usar `setDefaults`

Utilizando el método `setDefaults` en una colección inexistente también se crea la colección de forma implícita.

Este documento proporciona una guía detallada sobre cómo manejar colecciones en la base de datos, incluyendo la creación, listado y eliminación de colecciones, así como las formas implícitas de crearlas.
