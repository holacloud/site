# InstantLogs

InstantLogs provides HTTP loggers for raw byte ingestion and streaming/filtering.

## Authentication

Logger management uses Glue authentication with `X-Glue-Authentication`. Logger-specific endpoints can be accessed by a Glue owner of the logger or by a logger API key using `Api-Key` and `Api-Secret`.

## API Overview

| Method | Path | Description |
|--------|------|-------------|
| GET | `/v1/loggers` | List loggers visible to the Glue user |
| POST | `/v1/loggers` | Create a logger for the Glue user |
| GET | `/v1/loggers/{id}` | Get logger details, including owners and API keys |
| DELETE | `/v1/loggers/{id}` | Delete a logger |
| POST | `/v1/loggers/{id}/ingest` | Ingest raw request bytes, returns `{ "n": bytes }` |
| GET | `/v1/loggers/{id}/filter?regex=...` | Stream entries matching a regex |
| GET | `/v1/loggers/{id}/events` | Stream events; `follow` is enabled by flag presence |
| GET | `/v1/loggers/{id}/stats` | Get runtime stats |
| POST | `/v1/loggers/{id}/apiKeys` | Create a logger API key |
| DELETE | `/v1/loggers/{id}/apiKeys/{apiKey}` | Delete a logger API key |
| POST | `/v1/ingest/events` | Ingest logframe events with `project_id`, returns `events` and `bytes` |
