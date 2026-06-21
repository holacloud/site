# Config

Config is a centralized configuration management service within the Hola.Cloud ecosystem. It enables you to store, retrieve, and manage application configurations in a hierarchical, versioned, and secure manner.

## Key Benefits

### Hierarchical Configuration
Config supports a structured hierarchy of project, environment, and service levels. This allows you to define base configurations and override specific values per environment (development, staging, production) or per service, reducing duplication and simplifying management.

### JSON-Based
All configurations are stored as JSON documents, making them easy to read, write, and integrate with any programming language or toolchain. Schemas are flexible, so you can store any valid JSON structure.

### Versioned
Every configuration update creates a new version. You can track changes over time, compare revisions, and roll back to a previous version if needed.

### Easy Integration
Config integrates seamlessly with other Hola.Cloud services such as InceptionDB, Lambda, and InstantLogs. Centralizing configuration across services reduces boilerplate and keeps your infrastructure consistent.

### Secure Access
The service exposes two API surfaces. The admin API (`/v0/configs`) is publicly accessible for management operations. The user API (`/v1/config`) requires an API key, ensuring that only authorized clients read or modify runtime configuration.

## API Overview

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | `/v0/configs` | List all configs (admin) | Public |
| POST | `/v0/configs` | Create a new config (admin) | Public |
| GET | `/v0/configs/{id}` | Get a config by ID (admin) | Public |
| DELETE | `/v0/configs/{id}` | Delete a config (admin) | Public |
| PATCH | `/v0/configs/{id}` | Partial update of a config (admin) | Public |
| GET | `/v1/config` | Retrieve the current user config | API Key |
| PATCH | `/v1/config` | Update the current user config | API Key |

Base URL: `https://api.hola.cloud`

## Best Use Cases

### Microservice Configuration
Centralize and version-control configuration for all your microservices. Each service fetches its own config at startup and receives updates without redeployment.

### Multi-Environment Deployments
Manage separate configurations for development, staging, and production from a single control plane. Override only the values that differ between environments.

### Feature Flags
Store feature-toggle flags as configuration entries and update them at runtime without code changes. The PATCH endpoint supports partial updates, making toggling individual flags efficient.
