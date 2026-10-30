# Insertar documentos

`POST /v1/databases/{databaseId}/collections/{collection}:insert` · API key

Petición: JSON Lines (JSONL), `Content-Type: application/x-ndjson` · Respuesta: JSONL

El cuerpo es JSON Lines: un objeto JSON por línea. Los campos que faltan toman los valores por defecto de la colección.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary '{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}'
```

Respuesta `201 Created`, los documentos guardados como JSON Lines:

```jsonl
{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}
```

Los documentos se guardan en orden. La inserción se detiene en el primer documento rechazado; los anteriores quedan guardados.

- Un array JSON (`[{...},{...}]`) no es JSON Lines: se rechaza con `400`.
- Si la colección no existe, la inserción la crea.
- El `id` no es único salvo que la colección tenga un índice `map` sobre él.

| Código | Cuándo |
|---|---|
| 400 | Una línea no es un objeto JSON (por ejemplo, el cuerpo es un array JSON), o tiene nombres de campo repetidos. |
| 409 | Un índice único ya tiene el valor. |
