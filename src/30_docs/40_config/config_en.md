# Config

Config exposes a small JSON configuration API in the Hola.Cloud ecosystem. It stores configuration entries and lets authenticated users read or update their runtime configuration.

## Key Benefits

### Entry-Based Configuration
The v0 API stores a config as a Thing with an `id` and `entries`. The v1 user API returns the current user's configuration as an `entries` map.

### JSON-Based
All configurations are stored as JSON documents, making them easy to read, write, and integrate with any programming language or toolchain. Schemas are flexible, so you can store any valid JSON structure.

### Simple Updates
Updates replace or patch the entry data accepted by the API.

### Easy Integration
Config integrates seamlessly with other Hola.Cloud services such as InceptionDB, Lambda, and InstantLogs. Centralizing configuration across services reduces boilerplate and keeps your infrastructure consistent.

### Secure Access
The service exposes two API surfaces. The v0 API (`/v0/configs`) manages config Things. The v1 user API (`/v1/config`) requires `X-Glue-Authentication` and reads or updates the authenticated user's runtime configuration.

## API Overview

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | `/v0/configs` | List all configs (admin) | Public |
| POST | `/v0/configs` | Create a new config (admin) | Public |
| GET | `/v0/configs/{id}` | Get a config by ID (admin) | Public |
| DELETE | `/v0/configs/{id}` | Delete a config (admin) | Public |
| PATCH | `/v0/configs/{id}` | Partial update of a config (admin) | Public |
| GET | `/v1/config` | Retrieve the current user config entries map | Glue auth |
| PATCH | `/v1/config` | Update the current user config entries map | Glue auth |

Base URL: `https://api.hola.cloud`

## Best Use Cases

### Runtime Configuration
Fetch the authenticated user's configuration from `/v1/config` at startup or during runtime.

### Feature Flags
Store feature-toggle flags as configuration entries and update them at runtime without code changes.
