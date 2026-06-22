# InstantLogs API Reference

Base URL: `https://api.hola.cloud`

## Authentication

Management endpoints use `X-Glue-Authentication`. Logger endpoints accept either a Glue owner of the logger or a logger API key with `Api-Key` and `Api-Secret`.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/v1/loggers` | List owned loggers |
| POST | `/v1/loggers` | Create logger |
| GET | `/v1/loggers/{id}` | Get logger with owners/API keys |
| DELETE | `/v1/loggers/{id}` | Delete logger |
| POST | `/v1/loggers/{id}/ingest` | Ingest raw bytes, returns `{ "n": number }` |
| GET | `/v1/loggers/{id}/filter?regex=...` | Stream filtered logs |
| GET | `/v1/loggers/{id}/events` | Stream events; `follow` is a presence flag |
| GET | `/v1/loggers/{id}/stats` | Runtime stats |
| POST | `/v1/loggers/{id}/apiKeys` | Create logger API key |
| DELETE | `/v1/loggers/{id}/apiKeys/{apiKey}` | Delete logger API key |
| POST | `/v1/loggers/{id}/owners` | Add logger owner |
| DELETE | `/v1/loggers/{id}/owners/{ownerId}` | Delete logger owner |
| POST | `/v1/ingest/events` | Ingest logframe events with `project_id`, returns `events` and `bytes` |
