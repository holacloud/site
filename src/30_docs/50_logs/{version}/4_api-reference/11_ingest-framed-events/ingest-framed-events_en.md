# Ingest Framed Events

`POST /v1/ingest/events` accepts logframe-encoded events. Each frame metadata must include `project_id`.

```json
{ "project_id": "project-1" }
```

Response:

```json
{
  "events": 2,
  "bytes": 128
}
```
