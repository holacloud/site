# Insertar documentos

`POST /v1/databases/{databaseId}/collections/{collection}:insert` · API key

El cuerpo es JSON Lines: un objeto JSON por línea. Los campos que faltan toman los valores por defecto de la colección.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  --data-binary '{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}'
```

Respuesta `201 Created`, los documentos guardados como JSON Lines:

```jsonl
{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}
```

Los documentos se guardan en orden. La inserción se detiene en el primer documento rechazado; los anteriores quedan guardados.

| Código | Cuándo |
|---|---|
| 400 | Una línea no es un objeto JSON, o tiene nombres de campo repetidos. |
| 409 | Un índice único ya tiene el valor. |
