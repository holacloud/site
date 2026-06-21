
# Managing Projects and Services

Projects are the top-level organizational unit in HolaCloud. They group together service instances, team members, API keys, and billing. This guide covers how to create and manage projects through the Console.

## Creating and Switching Between Projects

To create a new project, open the project selector in the sidebar and click **New Project**. Provide a name and optional description. Once created, the Console switches to the new project automatically.

```bash
# Create a project via the Console API
curl -X POST "https://console.hola.cloud/api/v1/projects" \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Production App",
    "description": "Production environment for my application"
  }'
```

```json
{
  "id": "prod-app-123",
  "name": "My Production App",
  "description": "Production environment for my application",
  "createdAt": "2026-06-20T12:00:00Z"
}
```

To switch projects, click the current project name in the sidebar and select a different one from the dropdown.

```bash
# List all accessible projects
curl -X GET "https://console.hola.cloud/api/v1/projects" \
  -H "Authorization: Bearer <your-token>"
```

```json
{
  "projects": [
    { "id": "default", "name": "Default", "role": "admin" },
    { "id": "prod-app-123", "name": "My Production App", "role": "admin" }
  ]
}
```

## Managing Service Instances per Project

Each project contains its own instances of HolaCloud services. Navigate to the service section (e.g., Databases, Functions, Files, Config, Logs, Queues, Scheduler) to manage instances within the current project.

For example, to list InceptionDB databases in the current project:

```bash
curl -X GET "https://console.hola.cloud/api/v1/projects/prod-app-123/databases" \
  -H "Authorization: Bearer <your-token>"
```

```json
{
  "databases": [
    { "id": "db-001", "name": "users-db", "engine": "inceptiondb", "status": "active" },
    { "id": "db-002", "name": "analytics-db", "engine": "inceptiondb", "status": "active" }
  ]
}
```

The same pattern applies to Lambda functions, Files buckets, Config stores, InstantLogs streams, Queues, and Scheduler jobs.

## API Key Management per Project

API keys authenticate programmatic access to HolaCloud services. Manage them under **Project Settings > API Keys**.

```bash
# Create an API key
curl -X POST "https://console.hola.cloud/api/v1/projects/prod-app-123/api-keys" \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "CI/CD Key",
    "scopes": ["lambda:invoke", "files:read", "config:read"]
  }'
```

```json
{
  "id": "key-abc123",
  "name": "CI/CD Key",
  "key": "hc_api_abc123...",
  "scopes": ["lambda:invoke", "files:read", "config:read"],
  "createdAt": "2026-06-20T14:00:00Z"
}
```

You can revoke keys at any time from the same settings page.

## Viewing Project Billing and Usage

The **Billing** section shows current usage metrics, invoice history, and cost breakdowns per service.

```bash
curl -X GET "https://console.hola.cloud/api/v1/projects/prod-app-123/billing" \
  -H "Authorization: Bearer <your-token>"
```

```json
{
  "currentMonth": {
    "period": "2026-06",
    "total": 125.40,
    "services": {
      "inceptiondb": 45.00,
      "lambda": 60.00,
      "files": 15.40,
      "config": 5.00
    }
  },
  "invoices": [
    { "id": "inv-001", "period": "2026-05", "total": 110.20, "status": "paid" }
  ]
}
```

## Team Member Management

Invite team members from **Project Settings > Team**. Assign roles: Admin, Developer, or Viewer.

```bash
curl -X POST "https://console.hola.cloud/api/v1/projects/prod-app-123/team" \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "colleague@company.com",
    "role": "developer"
  }'
```

```json
{
  "id": "team-456",
  "email": "colleague@company.com",
  "role": "developer",
  "status": "pending"
}
```

## Import/Export Configuration as JSON

Export your entire project configuration (service settings, API keys metadata, team roster) as a single JSON file. Navigate to **Project Settings > Import/Export**.

```bash
# Export configuration
curl -X GET "https://console.hola.cloud/api/v1/projects/prod-app-123/export" \
  -H "Authorization: Bearer <your-token>"
```

```json
{
  "project": { "name": "My Production App" },
  "services": {
    "inceptiondb": { "databases": [...] },
    "lambda": { "functions": [...] },
    "files": { "buckets": [...] },
    "config": { "entries": [...] }
  },
  "team": [...]
}
```

To import, upload a JSON file of the same format via the Console UI or API:

```bash
curl -X POST "https://console.hola.cloud/api/v1/projects/prod-app-123/import" \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d @project-config.json
```

```json
{
  "status": "imported",
  "servicesUpdated": ["inceptiondb", "lambda", "files", "config"]
}
```

## Summary

Projects in HolaCloud provide isolated environments for your applications. Through the Console you can create and switch projects, manage service instances, API keys, billing, team members, and import/export configuration — all from a single interface.
