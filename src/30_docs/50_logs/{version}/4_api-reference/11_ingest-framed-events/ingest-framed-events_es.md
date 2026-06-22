# Ingerir Eventos Enmarcados

`POST /v1/ingest/events` acepta eventos codificados como logframe. La metadata de cada frame debe incluir `project_id`.

```json
{ "project_id": "project-1" }
```

Respuesta:

```json
{
  "events": 2,
  "bytes": 128
}
```
