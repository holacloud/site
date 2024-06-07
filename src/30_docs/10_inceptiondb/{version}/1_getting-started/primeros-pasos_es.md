# Guía de Inicio Rápido con InceptionDB

Esta guía te ayudará a empezar con InceptionDB mediante la creación de una colección, la inserción de tres elementos y
la consulta de elementos filtrados.

## Paso 1: Crear una Colección

Para crear una colección en InceptionDB, necesitas definir el nombre de la colección. Aquí hay un ejemplo de cómo hacerlo usando curl en bash:

```bash
curl -X POST "https://example.com/v1/databases/tu-id-de-base-de-datos/collections" \
-H "Api-Key: tu-api-key" \
-H "Api-Secret: tu-api-secret" \
-d '{
"name": "mi-colección"
}'
```

Respuesta Esperada

```json
{
	"defaults": {
		"id": "uuid()"
	},
	"indexes": 0,
	"name": "mi-colección",
	"total": 0
}
```

## Paso 2: Insertar Elementos

Una vez que hayas creado la colección, puedes insertar elementos en ella. A continuación se muestra cómo insertar tres elementos en la colección recién creada:

```bash
curl -X POST "https://example.com/v1/databases/tu-id-de-base-de-datos/collections/mi-colección/documents" \
-H "Api-Key: tu-api-key" \
-H "Api-Secret: tu-api-secret" \
-d '{"nombre": "Elemento 1", "valor": 100}
{"nombre": "Elemento 2", "valor": 200}
{"nombre": "Elemento 3", "valor": 300}
'
```

## Paso 3: Listar Elementos con Filtro

Para listar elementos en la colección con un filtro específico, puedes usar el siguiente ejemplo de solicitud con curl:

```bash
curl -X POST "https://example.com/v1/databases/tu-id-de-base-de-datos/collections/mi-colección/find" \
-H "Api-Key: tu-api-key" \
-H "Api-Secret: tu-api-secret" \
-d '{
      "filter": {
        "valor": { "$gte": 200 }
      }
    }'
```

Respuesta Esperada


```json
{
		"id": "id-del-documento-2",
		"nombre": "Elemento 2",
		"valor": 200
}
{
		"id": "id-del-documento-3",
		"nombre": "Elemento 3",
		"valor": 300
}
```

## Resumen

Siguiendo estos pasos, has creado una colección en InceptionDB, insertado tres elementos y consultado elementos con un filtro. Esto te proporciona una base sólida para empezar a trabajar con InceptionDB y aprovechar sus capacidades para gestionar tus datos de manera eficiente.
